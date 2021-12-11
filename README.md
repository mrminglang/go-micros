# github.com/mrminglang/go-micros
go version go1.14

## go-micro微服务框架实例

#### 介绍
golang使用go-micro

#### 使用说明
- go-micro 依赖go版本为1.13-1.14
- 安装protobuf
```exec
brew install consul
brew install protobuf
go get -u -v github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u -v github.com/micro/protoc-gen-micro
go get github.com/micro/go-micro
```
- 错误解决
    - undefined: grpc.SupportPackageIsVersion6 grpc.ClientConnInterface 错误，由于etcd版本管理的问题，导致etcd的代码和新版本的grpc冲突,会在编译时报错
    ``` 
    go.mod文件
    replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
    ```
  
    - go mod undefined: balancer.PickOptions 错误，protoc-gen-go兼容grpc v1.26.0的最新版本是v1.3.2
   ```
    go get github.com/golang/protobuf/protoc-gen-go@v1.3.2
   ```

#### 特技
1. grpc多语言 [www.cnblogs.com](https://www.cnblogs.com/chenqionghe/p/12394845.html)
