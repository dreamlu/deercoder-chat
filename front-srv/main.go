package main

import (
	"github.com/dreamlu/go-tool"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/web"
	"log"
)

func main() {

	//consul
	registry := consul.NewRegistry(consul.Config(
		&api.Config{
			Address: der.GetDevModeConfig("consul.address"),
			Scheme:  der.GetDevModeConfig("consul.scheme"),
		}))

	service := web.NewService(
		web.Name("deercoder-chat.common"),
		web.Registry(registry),
		web.Address(":"+der.GetDevModeConfig("http_port")),
		web.StaticDir("./static"),
	)

	// service init
	_ = service.Init()
	gin.SetMode(gin.DebugMode)
	// 路由
	router := gin.Default()
	router.Static("/", "static")
	// 注册
	service.Handle("/", router)
	// run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
