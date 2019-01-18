# lesson_3_go_micro_booking

### 环境准备
安装依赖consul
```
brew install consul
```

安装Go Micro
```
￼go get github.com/micro/go-micro
```

如果您使用代码生成,您还需要使用protoc-gen-go
```
￼go get github.com/micro/protobuf/{proto,protoc-gen-go}
```

### 编译proto
```
protoc --go_out=plugins=micro:. hello.proto
```

### 运行
```
consul agent --dev
go run s.go
go run c.go
```
