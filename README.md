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

* compile helloworld

```bash
$ protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld
```

* compile book

```bash
$ protoc -I proto/book/ proto/book/book.proto --go_out=plugins=grpc:proto/book
```

## Exec

### hello world

* run server

```bash
$ go run greeter_server/main.go
```

* run client

```bash
$ go run greeter_client/main.go
```

### book

* run server

```bash
$ go run server/book/main.go
```

* run client

```bash
$ go run client/book/main.go
```
