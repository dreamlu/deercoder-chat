package main

import (
	"deercoder-chat/chat-srv/controllers/chat"
	"deercoder-chat/chat-srv/proto"
	"github.com/dreamlu/deercoder-gin"
	"github.com/hashicorp/consul/api"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"
	"log"
)

func main() {

	// registry
	registry := consul.NewRegistry(consul.Config(
		&api.Config{
			Address: deercoder.GetDevModeConfig("consul.address"),
			Scheme:  deercoder.GetDevModeConfig("consul.scheme"),
		}))

	service := micro.NewService(
		micro.Name("deercoder-chat.chat"),
		micro.Registry(registry),
		//micro.RegisterTTL(time.Second*30),
		//micro.RegisterInterval(time.Second*10),
		micro.Address(":"+deercoder.GetDevModeConfig("http_port")),
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
