# lesson_3_go_micro_booking

### 环境准备

```
go get github.com/ziyoubiancheng/lesson_gomicro
￼go get github.com/micro/protobuf/{proto,protoc-gen-go}
go get -u github.com/micro/protoc-gen-micro
go get -u github.com/micro/go-micro
go get -u github.com/hailocab/go-geoindex
```

### 编译proto
在proto目录下执行
```
protoc --micro_out=. --go_out=. auth.proto 
protoc --micro_out=. --go_out=. geo.proto 
protoc --micro_out=. --go_out=. profile.proto 
protoc --micro_out=. --go_out=. rate.proto 
protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. hotel.proto 
```

### 运行
```
consul agent --dev
go run srv/auth/main.go
go run srv/geo/main.go 
go run srv/profile/main.go
go run srv/rate/main.go
go run api/hotel/main.go
micro api
```

### 请求接口

模拟验证token识别的情况
```
curl -H 'Content-Type: application/json' \
    -H "Authorization: Bearer INVALID_TOKEN" \
    -d '{"inDate": "2015-04-09"}' \
    http://localhost:8080/hotel/rates
```

模拟参数错误
```
curl -H 'Content-Type: application/json' \
    -H "Authorization: Bearer VALID_TOKEN" \
    -d '{"inDate": "2015-04-09"}' \
    http://localhost:8080/hotel/rates
```

模拟正常请求
```
curl -H 'Content-Type: application/json' \
    -H "Authorization: Bearer VALID_TOKEN" \
    -d '{"inDate": "2015-04-09", "outDate": "2015-04-10"}' \
    http://localhost:8080/hotel/rates
```


