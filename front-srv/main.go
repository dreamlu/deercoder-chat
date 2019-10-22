package main

import (
	"github.com/dreamlu/go-tool"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/web"
	"log"
)

func main() {

	// registry
	reg := consul.NewRegistry(
		registry.Addrs(gt.Configger().GetString("app.consul.address")),
	)

	service := web.NewService(
		web.Name("deercoder-chat.front"),
		web.Registry(reg),
		web.Address(":"+gt.Configger().GetString("app.port")),
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
