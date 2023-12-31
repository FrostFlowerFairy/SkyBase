// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: graphik.proto

package apipb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/struct"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
	regexp "regexp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_Ref_Gtype = regexp.MustCompile(`^.{1,225}$`)
var _regex_Ref_Gid = regexp.MustCompile(`^.{1,225}$`)

func (this *Ref) Validate() error {
	if !_regex_Ref_Gtype.MatchString(this.Gtype) {
		return github_com_mwitkow_go_proto_validators.FieldError("Gtype", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Gtype))
	}
	if !_regex_Ref_Gid.MatchString(this.Gid) {
		return github_com_mwitkow_go_proto_validators.FieldError("Gid", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Gid))
	}
	return nil
}

var _regex_RefConstructor_Gtype = regexp.MustCompile(`^.{1,225}$`)

func (this *RefConstructor) Validate() error {
	if !_regex_RefConstructor_Gtype.MatchString(this.Gtype) {
		return github_com_mwitkow_go_proto_validators.FieldError("Gtype", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Gtype))
	}
	return nil
}
func (this *Refs) Validate() error {
	for _, item := range this.Refs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Refs", err)
			}
		}
	}
	return nil
}
func (this *Doc) Validate() error {
	if nil == this.Ref {
		return github_com_mwitkow_go_proto_validators.FieldError("Ref", fmt.Errorf("message must exist"))
	}
	if this.Ref != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Ref); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Ref", err)
		}
	}
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *DocConstructor) Validate() error {
	if nil == this.Ref {
		return github_com_mwitkow_go_proto_validators.FieldError("Ref", fmt.Errorf("message must exist"))
	}
	if this.Ref != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Ref); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Ref", err)
		}
	}
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *DocConstructors) Validate() error {
	for _, item := range this.Docs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Docs", err)
			}
		}
	}
	return nil
}
func (this *Traversal) Validate() error {
	if this.Doc != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Doc); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Doc", err)
		}
	}
	for _, item := range this.TraversalPath {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("TraversalPath", err)
			}
		}
	}
	return nil
}
func (this *Traversals) Validate() error {
	for _, item := range this.Traversals {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Traversals", err)
			}
		}
	}
	return nil
}
func (this *Docs) Validate() error {
	for _, item := range this.Docs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Docs", err)
			}
		}
	}
	return nil
}
func (this *Connection) Validate() error {
	if nil == this.Ref {
		return github_com_mwitkow_go_proto_validators.FieldError("Ref", fmt.Errorf("message must exist"))
	}
	if this.Ref != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Ref); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Ref", err)
		}
	}
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	if nil == this.From {
		return github_com_mwitkow_go_proto_validators.FieldError("From", fmt.Errorf("message must exist"))
	}
	if this.From != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.From); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("From", err)
		}
	}
	if nil == this.To {
		return github_com_mwitkow_go_proto_validators.FieldError("To", fmt.Errorf("message must exist"))
	}
	if this.To != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.To); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("To", err)
		}
	}
	return nil
}
func (this *ConnectionConstructor) Validate() error {
	if nil == this.Ref {
		return github_com_mwitkow_go_proto_validators.FieldError("Ref", fmt.Errorf("message must exist"))
	}
	if this.Ref != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Ref); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Ref", err)
		}
	}
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	if nil == this.From {
		return github_com_mwitkow_go_proto_validators.FieldError("From", fmt.Errorf("message must exist"))
	}
	if this.From != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.From); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("From", err)
		}
	}
	if nil == this.To {
		return github_com_mwitkow_go_proto_validators.FieldError("To", fmt.Errorf("message must exist"))
	}
	if this.To != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.To); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("To", err)
		}
	}
	return nil
}
func (this *SearchConnectFilter) Validate() error {
	if this.Filter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Filter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Filter", err)
		}
	}
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	if nil == this.From {
		return github_com_mwitkow_go_proto_validators.FieldError("From", fmt.Errorf("message must exist"))
	}
	if this.From != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.From); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("From", err)
		}
	}
	return nil
}
func (this *SearchConnectMeFilter) Validate() error {
	if this.Filter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Filter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Filter", err)
		}
	}
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *ConnectionConstructors) Validate() error {
	for _, item := range this.Connections {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Connections", err)
			}
		}
	}
	return nil
}
func (this *Connections) Validate() error {
	for _, item := range this.Connections {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Connections", err)
			}
		}
	}
	return nil
}

var _regex_ConnectFilter_Gtype = regexp.MustCompile(`^.{1,225}$`)
var _regex_ConnectFilter_Sort = regexp.MustCompile(`((^|, )(|ref.gid|ref.gtype|^attributes.(.*)))+$`)

func (this *ConnectFilter) Validate() error {
	if nil == this.DocRef {
		return github_com_mwitkow_go_proto_validators.FieldError("DocRef", fmt.Errorf("message must exist"))
	}
	if this.DocRef != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DocRef); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DocRef", err)
		}
	}
	if !_regex_ConnectFilter_Gtype.MatchString(this.Gtype) {
		return github_com_mwitkow_go_proto_validators.FieldError("Gtype", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Gtype))
	}
	if !(this.Limit > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Limit", fmt.Errorf(`value '%v' must be greater than '0'`, this.Limit))
	}
	if !_regex_ConnectFilter_Sort.MatchString(this.Sort) {
		return github_com_mwitkow_go_proto_validators.FieldError("Sort", fmt.Errorf(`value '%v' must be a string conforming to regex "((^|, )(|ref.gid|ref.gtype|^attributes.(.*)))+$"`, this.Sort))
	}
	return nil
}

var _regex_Filter_Gtype = regexp.MustCompile(`^.{1,225}$`)
var _regex_Filter_Sort = regexp.MustCompile(`((^|, )(|ref.gid|ref.gtype|^attributes.(.*)))+$`)

func (this *Filter) Validate() error {
	if !_regex_Filter_Gtype.MatchString(this.Gtype) {
		return github_com_mwitkow_go_proto_validators.FieldError("Gtype", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Gtype))
	}
	if !(this.Limit > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Limit", fmt.Errorf(`value '%v' must be greater than '0'`, this.Limit))
	}
	if !_regex_Filter_Sort.MatchString(this.Sort) {
		return github_com_mwitkow_go_proto_validators.FieldError("Sort", fmt.Errorf(`value '%v' must be a string conforming to regex "((^|, )(|ref.gid|ref.gtype|^attributes.(.*)))+$"`, this.Sort))
	}
	return nil
}

var _regex_AggFilter_Field = regexp.MustCompile(`((^|, )(|^attributes.(.*)))+$`)

func (this *AggFilter) Validate() error {
	if nil == this.Filter {
		return github_com_mwitkow_go_proto_validators.FieldError("Filter", fmt.Errorf("message must exist"))
	}
	if this.Filter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Filter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Filter", err)
		}
	}
	if !_regex_AggFilter_Field.MatchString(this.Field) {
		return github_com_mwitkow_go_proto_validators.FieldError("Field", fmt.Errorf(`value '%v' must be a string conforming to regex "((^|, )(|^attributes.(.*)))+$"`, this.Field))
	}
	return nil
}

var _regex_TraverseFilter_Sort = regexp.MustCompile(`((^|, )(|ref.gid|ref.gtype|^attributes.(.*)))+$`)

func (this *TraverseFilter) Validate() error {
	if nil == this.Root {
		return github_com_mwitkow_go_proto_validators.FieldError("Root", fmt.Errorf("message must exist"))
	}
	if this.Root != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Root); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Root", err)
		}
	}
	if !(this.Limit > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Limit", fmt.Errorf(`value '%v' must be greater than '0'`, this.Limit))
	}
	if !_regex_TraverseFilter_Sort.MatchString(this.Sort) {
		return github_com_mwitkow_go_proto_validators.FieldError("Sort", fmt.Errorf(`value '%v' must be a string conforming to regex "((^|, )(|ref.gid|ref.gtype|^attributes.(.*)))+$"`, this.Sort))
	}
	if !(this.MaxDepth > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("MaxDepth", fmt.Errorf(`value '%v' must be greater than '0'`, this.MaxDepth))
	}
	if !(this.MaxHops > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("MaxHops", fmt.Errorf(`value '%v' must be greater than '0'`, this.MaxHops))
	}
	return nil
}

var _regex_TraverseMeFilter_Sort = regexp.MustCompile(`((^|, )(|ref.gid|ref.gtype|^attributes.(.*)))+$`)

func (this *TraverseMeFilter) Validate() error {
	if !(this.Limit > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Limit", fmt.Errorf(`value '%v' must be greater than '0'`, this.Limit))
	}
	if !_regex_TraverseMeFilter_Sort.MatchString(this.Sort) {
		return github_com_mwitkow_go_proto_validators.FieldError("Sort", fmt.Errorf(`value '%v' must be a string conforming to regex "((^|, )(|ref.gid|ref.gtype|^attributes.(.*)))+$"`, this.Sort))
	}
	if !(this.MaxDepth > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("MaxDepth", fmt.Errorf(`value '%v' must be greater than '0'`, this.MaxDepth))
	}
	if !(this.MaxHops > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("MaxHops", fmt.Errorf(`value '%v' must be greater than '0'`, this.MaxHops))
	}
	return nil
}

var _regex_IndexConstructor_Name = regexp.MustCompile(`^.{1,225}$`)
var _regex_IndexConstructor_Gtype = regexp.MustCompile(`^.{1,225}$`)
var _regex_IndexConstructor_Expression = regexp.MustCompile(`^.{1,225}$`)

func (this *IndexConstructor) Validate() error {
	if !_regex_IndexConstructor_Name.MatchString(this.Name) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Name))
	}
	if !_regex_IndexConstructor_Gtype.MatchString(this.Gtype) {
		return github_com_mwitkow_go_proto_validators.FieldError("Gtype", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Gtype))
	}
	if !_regex_IndexConstructor_Expression.MatchString(this.Expression) {
		return github_com_mwitkow_go_proto_validators.FieldError("Expression", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Expression))
	}
	return nil
}
func (this *AuthTarget) Validate() error {
	if nil == this.User {
		return github_com_mwitkow_go_proto_validators.FieldError("User", fmt.Errorf("message must exist"))
	}
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	if this.Target != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Target); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Target", err)
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}

var _regex_Authorizer_Name = regexp.MustCompile(`^.{1,225}$`)
var _regex_Authorizer_Method = regexp.MustCompile(`^.{1,225}$`)
var _regex_Authorizer_Expression = regexp.MustCompile(`^.{1,225}$`)

func (this *Authorizer) Validate() error {
	if !_regex_Authorizer_Name.MatchString(this.Name) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Name))
	}
	if !_regex_Authorizer_Method.MatchString(this.Method) {
		return github_com_mwitkow_go_proto_validators.FieldError("Method", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Method))
	}
	if !_regex_Authorizer_Expression.MatchString(this.Expression) {
		return github_com_mwitkow_go_proto_validators.FieldError("Expression", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Expression))
	}
	return nil
}
func (this *Authorizers) Validate() error {
	for _, item := range this.Authorizers {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Authorizers", err)
			}
		}
	}
	return nil
}

var _regex_Constraint_Name = regexp.MustCompile(`^.{1,225}$`)
var _regex_Constraint_Gtype = regexp.MustCompile(`^.{1,225}$`)
var _regex_Constraint_Expression = regexp.MustCompile(`^.{1,225}$`)

func (this *Constraint) Validate() error {
	if !_regex_Constraint_Name.MatchString(this.Name) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Name))
	}
	if !_regex_Constraint_Gtype.MatchString(this.Gtype) {
		return github_com_mwitkow_go_proto_validators.FieldError("Gtype", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Gtype))
	}
	if !_regex_Constraint_Expression.MatchString(this.Expression) {
		return github_com_mwitkow_go_proto_validators.FieldError("Expression", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Expression))
	}
	return nil
}
func (this *Constraints) Validate() error {
	for _, item := range this.Constraints {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Constraints", err)
			}
		}
	}
	return nil
}

var _regex_Index_Name = regexp.MustCompile(`^.{1,225}$`)
var _regex_Index_Gtype = regexp.MustCompile(`^.{1,225}$`)
var _regex_Index_Expression = regexp.MustCompile(`^.{1,225}$`)

func (this *Index) Validate() error {
	if !_regex_Index_Name.MatchString(this.Name) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Name))
	}
	if !_regex_Index_Gtype.MatchString(this.Gtype) {
		return github_com_mwitkow_go_proto_validators.FieldError("Gtype", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Gtype))
	}
	if !_regex_Index_Expression.MatchString(this.Expression) {
		return github_com_mwitkow_go_proto_validators.FieldError("Expression", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Expression))
	}
	return nil
}
func (this *Indexes) Validate() error {
	for _, item := range this.Indexes {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Indexes", err)
			}
		}
	}
	return nil
}

var _regex_Trigger_Name = regexp.MustCompile(`^.{1,225}$`)
var _regex_Trigger_Gtype = regexp.MustCompile(`^.{1,225}$`)
var _regex_Trigger_Trigger = regexp.MustCompile(`^.{1,225}$`)

func (this *Trigger) Validate() error {
	if !_regex_Trigger_Name.MatchString(this.Name) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Name))
	}
	if !_regex_Trigger_Gtype.MatchString(this.Gtype) {
		return github_com_mwitkow_go_proto_validators.FieldError("Gtype", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Gtype))
	}
	if !_regex_Trigger_Trigger.MatchString(this.Trigger) {
		return github_com_mwitkow_go_proto_validators.FieldError("Trigger", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Trigger))
	}
	return nil
}
func (this *Triggers) Validate() error {
	for _, item := range this.Triggers {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Triggers", err)
			}
		}
	}
	return nil
}

var _regex_StreamFilter_Channel = regexp.MustCompile(`^.{1,225}$`)

func (this *StreamFilter) Validate() error {
	if !_regex_StreamFilter_Channel.MatchString(this.Channel) {
		return github_com_mwitkow_go_proto_validators.FieldError("Channel", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Channel))
	}
	if this.Min != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Min); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Min", err)
		}
	}
	if this.Max != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Max); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Max", err)
		}
	}
	return nil
}
func (this *Graph) Validate() error {
	if this.Docs != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Docs); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Docs", err)
		}
	}
	if this.Connections != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Connections); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Connections", err)
		}
	}
	return nil
}
func (this *Flags) Validate() error {
	return nil
}
func (this *UIFlags) Validate() error {
	return nil
}
func (this *Boolean) Validate() error {
	return nil
}
func (this *Number) Validate() error {
	return nil
}

var _regex_ExistsFilter_Gtype = regexp.MustCompile(`^.{1,225}$`)
var _regex_ExistsFilter_Expression = regexp.MustCompile(`^.{1,225}$`)

func (this *ExistsFilter) Validate() error {
	if !_regex_ExistsFilter_Gtype.MatchString(this.Gtype) {
		return github_com_mwitkow_go_proto_validators.FieldError("Gtype", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Gtype))
	}
	if !_regex_ExistsFilter_Expression.MatchString(this.Expression) {
		return github_com_mwitkow_go_proto_validators.FieldError("Expression", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Expression))
	}
	return nil
}
func (this *Edit) Validate() error {
	if nil == this.Ref {
		return github_com_mwitkow_go_proto_validators.FieldError("Ref", fmt.Errorf("message must exist"))
	}
	if this.Ref != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Ref); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Ref", err)
		}
	}
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *EditFilter) Validate() error {
	if this.Filter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Filter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Filter", err)
		}
	}
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *Pong) Validate() error {
	return nil
}

var _regex_OutboundMessage_Channel = regexp.MustCompile(`^.{1,225}$`)

func (this *OutboundMessage) Validate() error {
	if !_regex_OutboundMessage_Channel.MatchString(this.Channel) {
		return github_com_mwitkow_go_proto_validators.FieldError("Channel", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Channel))
	}
	if nil == this.Data {
		return github_com_mwitkow_go_proto_validators.FieldError("Data", fmt.Errorf("message must exist"))
	}
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}

var _regex_Message_Channel = regexp.MustCompile(`^.{1,225}$`)
var _regex_Message_Method = regexp.MustCompile(`^.{1,225}$`)

func (this *Message) Validate() error {
	if !_regex_Message_Channel.MatchString(this.Channel) {
		return github_com_mwitkow_go_proto_validators.FieldError("Channel", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Channel))
	}
	if nil == this.Data {
		return github_com_mwitkow_go_proto_validators.FieldError("Data", fmt.Errorf("message must exist"))
	}
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	if nil == this.User {
		return github_com_mwitkow_go_proto_validators.FieldError("User", fmt.Errorf("message must exist"))
	}
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	if nil == this.Timestamp {
		return github_com_mwitkow_go_proto_validators.FieldError("Timestamp", fmt.Errorf("message must exist"))
	}
	if this.Timestamp != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Timestamp); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Timestamp", err)
		}
	}
	if !_regex_Message_Method.MatchString(this.Method) {
		return github_com_mwitkow_go_proto_validators.FieldError("Method", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Method))
	}
	return nil
}
func (this *Schema) Validate() error {
	if this.Authorizers != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Authorizers); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Authorizers", err)
		}
	}
	if this.Constraints != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Constraints); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Constraints", err)
		}
	}
	if this.Indexes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Indexes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Indexes", err)
		}
	}
	if this.Triggers != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Triggers); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Triggers", err)
		}
	}
	return nil
}
func (this *ExprFilter) Validate() error {
	return nil
}
func (this *RaftCommand) Validate() error {
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	for _, item := range this.SetDocs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("SetDocs", err)
			}
		}
	}
	for _, item := range this.SetConnections {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("SetConnections", err)
			}
		}
	}
	for _, item := range this.DelDocs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("DelDocs", err)
			}
		}
	}
	for _, item := range this.DelConnections {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("DelConnections", err)
			}
		}
	}
	if this.SetIndexes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SetIndexes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SetIndexes", err)
		}
	}
	if this.SetAuthorizers != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SetAuthorizers); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SetAuthorizers", err)
		}
	}
	if this.SetConstraints != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SetConstraints); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SetConstraints", err)
		}
	}
	if this.SendMessage != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SendMessage); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SendMessage", err)
		}
	}
	if this.SetTriggers != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SetTriggers); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SetTriggers", err)
		}
	}
	return nil
}
func (this *Peer) Validate() error {
	return nil
}
func (this *RaftState) Validate() error {
	for _, item := range this.Peers {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Peers", err)
			}
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
