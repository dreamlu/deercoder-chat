package main

import (
	"deercoder-chat/user-srv/controllers"
	user "deercoder-chat/user-srv/proto"
	"github.com/dreamlu/deercoder-gin"
	"github.com/hashicorp/consul/api"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"
	"log"
	"time"
)


func main() {

	// registry
	registry := consul.NewRegistry(consul.Config(
		&api.Config{
			Address: deercoder.GetDevModeConfig("consul.address"),
			Scheme:  deercoder.GetDevModeConfig("consul.scheme"),
		}))

	service := micro.NewService(
		micro.Name("deercoder-chat.user"),
		micro.Registry(registry),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Address(":"+deercoder.GetDevModeConfig("http_port")),
	)

	// service init
	service.Init()

	// Register Handlers
	// user register
	_ = user.RegisterUserServiceHandler(service.Server(), new(controllers.UserService))
	// login register
	_ = user.RegisterLoginServiceHandler(service.Server(), new(controllers.LoginService))

	// run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
