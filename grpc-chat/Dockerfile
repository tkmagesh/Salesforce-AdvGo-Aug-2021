FROM golang:alpine as build-env

ENV GO111MODULE=on

RUN apk update && apk add bash ca-certificates git gcc g++ libc-dev

RUN mkdir /grpc-chat
RUN mkdir -p /grpc-chat/proto 

WORKDIR /grpc-chat

COPY ./proto/service.pb.go /grpc-chat/proto
COPY ./server/server.go /grpc-chat

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go build -o grpc-chat .

CMD ./grpc-chat