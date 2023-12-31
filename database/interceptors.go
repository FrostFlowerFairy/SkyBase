package database

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/graphikDB/graphik/gen/grpc/go"
	"github.com/graphikDB/graphik/helpers"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/pkg/errors"
	"go.etcd.io/bbolt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func (g *Graph) UnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if strings.Contains(info.FullMethod, "JoinCluster") {
			return handler(ctx, req)
		}
		token, err := grpc_auth.AuthFromMD(ctx, "Bearer")
		if err != nil {
			return nil, err
		}
		ctx = g.tokenToContext(ctx, token)
		tokenHash := helpers.Hash([]byte(token))
		if val, ok := g.jwtCache.Get(tokenHash); ok {
			payload := val.(map[string]interface{})
			ctx, err := g.checkRequest(ctx, info.FullMethod, req, payload)
			if err != nil {
				return nil, err
			}
			hresp, err := handler(ctx, req)
			if err != nil {
				return hresp, err
			}
			if err := g.checkResponse(ctx, info.FullMethod, hresp, payload); err != nil {
				return nil, err
			}
			return hresp, nil
		}
		ctx = g.methodToContext(ctx, info.FullMethod)
		userinfoReq, err := http.NewRequest(http.MethodGet, g.openID.UserinfoEndpoint, nil)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "failed to create user info request: %s", err.Error())
		}
		userinfoReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		resp, err := http.DefaultClient.Do(userinfoReq)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "failed to get userinfo: %s", err.Error())
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			return nil, status.Errorf(codes.Unauthenticated, "failed to get userinfo: %v", resp.StatusCode)
		}
		bits, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "failed to get userinfo: %s", err.Error())
		}
		payload := map[string]interface{}{}
		if err := json.Unmarshal(bits, &payload); err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "failed to get userinfo: %s", err.Error())
		}
		g.jwtCache.Set(tokenHash, payload, 1*time.Hour)
		ctx, err = g.checkRequest(ctx, info.FullMethod, req, payload)
		if err != nil {
			return nil, err
		}
		hresp, err := handler(ctx, req)
		if err != nil {
			return hresp, err
		}
		if err := g.checkResponse(ctx, info.FullMethod, hresp, payload); err != nil {
			return nil, err
		}
		return hresp, nil
	}
}

func (g *Graph) StreamInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx := g.methodToContext(ss.Context(), info.FullMethod)
		token, err := grpc_auth.AuthFromMD(ctx, "Bearer")
		if err != nil {
			return err
		}
		ctx = g.tokenToContext(ctx, token)
		tokenHash := helpers.Hash([]byte(token))
		if val, ok := g.jwtCache.Get(tokenHash); ok {
			payload := val.(map[string]interface{})
			if info.IsClientStream {
				ctx, err := g.checkRequest(ctx, info.FullMethod, srv, payload)
				if err != nil {
					return err
				}
				wrapped := grpc_middleware.WrapServerStream(ss)
				wrapped.WrappedContext = ctx
				return handler(srv, wrapped)
			} else {
				if err := g.checkResponse(ctx, info.FullMethod, srv, payload); err != nil {
					return err
				}
			}
		}
		userinfoReq, err := http.NewRequest(http.MethodGet, g.openID.UserinfoEndpoint, nil)
		if err != nil {
			return status.Errorf(codes.Unauthenticated, "failed to get userinfo: %s", err.Error())
		}
		userinfoReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		resp, err := http.DefaultClient.Do(userinfoReq)
		if err != nil {
			return status.Errorf(codes.Unauthenticated, "failed to get userinfo: %s", err.Error())
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			return status.Errorf(codes.Unauthenticated, "failed to get userinfo: %v", resp.StatusCode)
		}
		bits, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return errors.Wrap(err, "failed to get userinfo")
		}
		payload := map[string]interface{}{}
		if err := json.Unmarshal(bits, &payload); err != nil {
			return status.Errorf(codes.Unauthenticated, "failed to get userinfo: %s", err.Error())
		}
		g.jwtCache.Set(tokenHash, payload, 1*time.Hour)
		ctx, err = g.checkRequest(ctx, info.FullMethod, srv, payload)
		if err != nil {
			return err
		}
		wrapped := grpc_middleware.WrapServerStream(ss)
		wrapped.WrappedContext = ctx
		return handler(srv, wrapped)
	}
}

func (a *Graph) userToContext(ctx context.Context, payload map[string]interface{}) (context.Context, *apipb.Doc, error) {
	email, ok := payload["email"].(string)
	if !ok {
		return nil, nil, errors.New("email not present in userinfo")
	}
	var (
		doc *apipb.Doc
		err error
	)
	a.db.View(func(tx *bbolt.Tx) error {
		doc, err = a.getDoc(ctx, tx, &apipb.Ref{
			Gtype: string(userType),
			Gid:   email,
		})
		if err != nil {
			return err
		}
		return nil
	})
	if doc == nil {
		strct, err := structpb.NewStruct(payload)
		if err != nil {
			return nil, nil, err
		}
		doc, err = a.createIdentity(ctx, &apipb.DocConstructor{
			Ref: &apipb.RefConstructor{
				Gtype: string(userType),
				Gid:   email,
			},
			Attributes: strct,
		})
		if err != nil {
			return nil, nil, err
		}
	}
	return context.WithValue(ctx, authCtxKey, doc), doc, nil
}

func (s *Graph) getIdentity(ctx context.Context) *apipb.Doc {
	val, ok := ctx.Value(authCtxKey).(*apipb.Doc)
	if ok {
		return val
	}
	val2, ok := ctx.Value(authCtxKey).(apipb.Doc)
	if ok {
		return &val2
	}
	return nil
}

func (r *Graph) getMethod(ctx context.Context) string {
	val, ok := ctx.Value(methodCtxKey).(string)
	if ok {
		return val
	}
	return ""
}

func (r *Graph) methodToContext(ctx context.Context, path string) context.Context {
	return context.WithValue(ctx, methodCtxKey, path)
}

func (r *Graph) tokenToContext(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, tokenCtxKey, token)
}

func (r *Graph) getToken(ctx context.Context) string {
	val, ok := ctx.Value(tokenCtxKey).(string)
	if ok {
		return val
	}
	return ""
}

func (g *Graph) checkRequest(ctx context.Context, method string, req interface{}, payload map[string]interface{}) (context.Context, error) {
	ctx = g.methodToContext(ctx, method)
	ctx, user, err := g.userToContext(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}
	if g.isGraphikAdmin(user) {
		return context.WithValue(ctx, bypassAuthorizersCtxKey, true), nil
	}
	var authorizers []*authorizer
	g.rangeAuthorizers(func(a *authorizer) bool {
		if a.authorizer.GetTargetRequests() && a.authorizer.GetMethod() == method {
			authorizers = append(authorizers, a)
		}
		return true
	})
	if g.options.requireResponseAuthorizers && len(authorizers) == 0 {
		return nil, status.Errorf(
			codes.PermissionDenied,
			"zero registered request authorizers found for invoked gRPC method %s", method)
	}
	if len(authorizers) == 0 {
		return ctx, nil
	}
	request := &apipb.AuthTarget{
		User:    user,
		Headers: map[string]string{},
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		for k, val := range md {
			if len(val) > 0 {
				request.Headers[k] = val[0]
			}
		}
	}
	if val, ok := req.(apipb.Mapper); ok {
		request.Target = apipb.NewStruct(val.AsMap())
	} else if val, ok := req.(proto.Message); ok {
		bits, _ := helpers.MarshalJSON(val)
		reqMap := map[string]interface{}{}
		if err := json.Unmarshal(bits, &reqMap); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		request.Target = apipb.NewStruct(reqMap)
	}
	for _, a := range authorizers {
		if err := a.decision.Eval(request.AsMap()); err == nil {
			return ctx, nil
		}
	}
	return nil, status.Errorf(codes.PermissionDenied, "request from %s.%s  authorization = denied", user.GetRef().GetGtype(), user.GetRef().GetGid())
}

func (g *Graph) checkResponse(ctx context.Context, method string, response interface{}, payload map[string]interface{}) error {
	if _, ok := response.(*empty.Empty); ok {
		return nil
	}
	ctx = g.methodToContext(ctx, method)
	var (
		err  error
		user = g.getIdentity(ctx)
	)
	if user == nil {
		ctx, user, err = g.userToContext(ctx, payload)
		if err != nil {
			return status.Errorf(codes.Unauthenticated, err.Error())
		}
	}
	if g.isGraphikAdmin(user) {
		return nil
	}
	var authorizers []*authorizer
	g.rangeAuthorizers(func(a *authorizer) bool {
		if a.authorizer.GetTargetResponses() && a.authorizer.GetMethod() == method {
			authorizers = append(authorizers, a)
		}
		return true
	})
	if g.options.requireResponseAuthorizers && len(authorizers) == 0 {
		return status.Errorf(
			codes.PermissionDenied,
			"zero registered response authorizers found for invoked gRPC method %s", method)
	}
	if len(authorizers) == 0 {
		return nil
	}
	request := &apipb.AuthTarget{
		User:    user,
		Target:  nil,
		Headers: map[string]string{},
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		for k, val := range md {
			if len(val) > 0 {
				request.Headers[k] = val[0]
			}
		}
	}
	md, ok = metadata.FromOutgoingContext(ctx)
	if ok {
		for k, val := range md {
			if len(val) > 0 {
				request.Headers[k] = val[0]
			}
		}
	}
	if val, ok := response.(apipb.Mapper); ok {
		request.Target = apipb.NewStruct(val.AsMap())
	} else if val, ok := response.(proto.Message); ok {
		bits, _ := helpers.MarshalJSON(val)
		reqMap := map[string]interface{}{}
		if err := json.Unmarshal(bits, &reqMap); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		request.Target = apipb.NewStruct(reqMap)
	}
	for _, a := range authorizers {
		if err := a.decision.Eval(request.AsMap()); err == nil {
			return nil
		}
	}
	return status.Errorf(codes.PermissionDenied, "response to %s.%s  authorization = denied", user.GetRef().GetGtype(), user.GetRef().GetGid())
}

func (g *Graph) isGraphikAdmin(user *apipb.Doc) bool {
	if user.GetAttributes().GetFields() == nil {
		return false
	}
	return helpers.ContainsString(user.GetAttributes().GetFields()["email"].GetStringValue(), g.options.rootUsers)
}
