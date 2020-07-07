package main

import (
	"deercoder-chat/chat-srv/handler/chat"
	"deercoder-chat/chat-srv/proto"
	"github.com/dreamlu/go-tool"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"log"
)

func main() {

	// registry
	reg := consul.NewRegistry(
		registry.Addrs(gt.Configger().GetString("app.consul.address")),
	)

	service := micro.NewService(
		micro.Name("deercoder-chat.chat"),
		micro.Registry(reg),
		//micro.RegisterTTL(time.Second*30),
		//micro.RegisterInterval(time.Second*10),
		micro.Address(":"+gt.Configger().GetString("app.port")),
	)

	// service init
	service.Init()

	// Register Handler
	_ = proto.RegisterStreamerHandler(service.Server(), new(chat.Streamer))
	_ = proto.RegisterChatServiceHandler(service.Server(), new(chat.ChatService))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
	// run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
