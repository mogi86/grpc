# grps

## Description

gRPCの学習用リポジトリ。

## Requirement

* Golang 1.11+
* The setting of GOPATH has been completed

## Set up

* install protobuf

```bash
$ brew install protobuf
```

* install gRPC

```bash
$ go get -u google.golang.org/grpc
```

* install protoc-gen-go

```bash
$ go get -u github.com/golang/protobuf/protoc-gen-go

```

## Compile

* compile

```bash
$ protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld
```

## Exec

* run server

```bash
$ go run greeter_server/main.go
```

* run client

```bash
$ go run greeter_client/main.go
```
