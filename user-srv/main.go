package main

import (
	"deercoder-chat/user-srv/controllers"
	user "deercoder-chat/user-srv/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"
	"log"
	"time"
)

func main() {

	service := micro.NewService(
		micro.Name("deercoder-chat.user"),
		micro.Registry(consul.NewRegistry()),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Address(":8000"),
	)

	// service init
	service.Init()

	// Register Handlers
	_ = user.RegisterUserServiceHandler(service.Server(), new(controllers.UserService))

	// run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
