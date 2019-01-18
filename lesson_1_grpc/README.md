## lesson_1_grpc

### 编译环境准备
安装protobuf
```
brew install protobuf
```
安装 gRPC 与 protoc 编译器
```
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
```

### 编译 user.proto 文件
```
cd proto

protoc -I . --go_out=plugins=grpc:. ./user.proto
```

### 运行
```
go run s.go
go run c.go
```


