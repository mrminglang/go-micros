package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"log"

	proto "github.com/mrminglang/go-micros/proto/hello"
)

type MicroFirst struct{}

func (m *MicroFirst) Hello(ctx context.Context, request *proto.HelloRequest, response *proto.HelloResponse) error {
	msg := fmt.Sprintf("hello, name:%s, actions:%s\n", request.Name, request.Actions)
	response.Msg = msg
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("hello"),
		micro.Version("bate"),
	)

	service.Init()

	_ = proto.RegisterMicroHelloHandler(service.Server(), new(MicroFirst))

	if err := service.Run(); err != nil {
		log.Fatalf(err.Error())
	}
}
