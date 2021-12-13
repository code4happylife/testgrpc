# testgrpc

### Use protoc to generate chat.pb.go
```shell
 protoc --go_out=plugins=grpc:chat chat.proto
```

### Run server.go
```shell
go run server.go
```

### Run client.go
```shell
go run client.go
```

### Test result
```shell
testgrpc % go run server.go
2021/12/13 09:40:56 Received message body from client:Hello from client !

 testgrpc % go run client.go
2021/12/13 09:40:56 Response from Server: Hello from the server!

```