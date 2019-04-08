package main

import (
	"deercoder-chat/api-gateway/routers"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-web"
	"log"
)

//var (
//	UserClient user.UserService
//)

func main() {
	// Create service
	service := web.NewService(
		web.Name("deercoder-chat.api"),
		web.Registry(consul.NewRegistry()),
		web.Address(":8006"),
	)

	_ = service.Init()

	//_ = service.Server().Handle(
	//	service.Server().NewHandler(
	//		&UserServie{
	//			userClient: user.NewUserService("deercoder-chat.user", service.Client()),
	//		},
	//		),
	//)

	//
	//micro.RegisterHandler(user, new(user.UserService))

	// setup user Server Client
	//UserClient = user.NewUserService("deercoder-chat.user", client.DefaultClient)

	// Create RESTful handler (using Gin)
	// Register Handler
	gin.SetMode(gin.DebugMode)
	service.Handle("/", routers.SetRouter())

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
