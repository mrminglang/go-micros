package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"

	proto "github.com/mrminglang/go-micros/proto/hello"
)

func main() {
	service := micro.NewService(
		micro.Name("hello"),
		micro.Version("bate"),
	)

	service.Init()

	hello := proto.NewMicroHelloService("hello", service.Client())

	response, err := hello.Hello(context.Background(), &proto.HelloRequest{
		Name:    "ming",
		Actions: []string{"111", "222"},
	})
	if err != nil {
		fmt.Println(err)
	}

	// 业务处理
	fmt.Println("code:", response.Code)
	fmt.Println("msg:", response.Msg)
}
