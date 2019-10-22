package main

import (
	"deercoder-chat/user-srv/handler"
	user "deercoder-chat/user-srv/proto"
	"github.com/dreamlu/go-tool"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"log"
	"time"
)


func main() {

	// registry
	reg := consul.NewRegistry(
		registry.Addrs(gt.Configger().GetString("app.consul.address")),
	)

	service := micro.NewService(
		micro.Name("deercoder-chat.user"),
		micro.Registry(reg),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Address(":"+gt.Configger().GetString("app.port")),
	)

	// service init
	service.Init()

	// Register Handlers
	// user register
	_ = user.RegisterUserServiceHandler(service.Server(), new(handler.UserService))
	// login register
	_ = user.RegisterLoginServiceHandler(service.Server(), new(handler.LoginService))

	// run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
