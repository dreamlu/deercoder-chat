package main

import (
	"deercoder-chat/chat-srv/controllers/chat"
	"deercoder-chat/chat-srv/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"
	"log"
)

func main() {
	service := micro.NewService(
		micro.Name("deercoder-chat.chat"),
		micro.Registry(consul.NewRegistry()),
		//micro.RegisterTTL(time.Second*30),
		//micro.RegisterInterval(time.Second*10),
		//micro.Address(":8000"),
	)

	// service init
	service.Init()

	// Register Handler
	_ = proto.RegisterStreamerHandler(service.Server(), new(chat.Streamer))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
	// run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
