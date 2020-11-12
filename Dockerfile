FROM golang:1.14.2-alpine3.11 as build-env

RUN mkdir /graphik
RUN apk --update add ca-certificates
RUN apk add make git
WORKDIR /graphik
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN mkdir -p bin && cd graphik; go build -o ../bin/graphik
ENTRYPOINT ["./bin/graphik"]