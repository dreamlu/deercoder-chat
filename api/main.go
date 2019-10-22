package main

import (
	"deercoder-chat/api/routers"
	"github.com/dreamlu/go-tool"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/web"
	"log"
	"net/http"
)

//var (
//	UserClient user.UserService
//)

func main() {

	// registry
	reg := consul.NewRegistry(
		registry.Addrs(gt.Configger().GetString("app.consul.address")),
	)

	// Create service
	service := web.NewService(
		// 这里指 所有的http 接口api, 非api网关
		web.Name("deercoder-chat.web.api"),
		web.Registry(reg),
		web.Address(":"+gt.Configger().GetString("app.port")),
	)

	_ = service.Init()

	// Create RESTful handler (using Gin)
	// Register Handler
	gin.SetMode(gin.DebugMode)
	// 路由
	router := routers.SetRouter()
	// 后台配置
	// 注释即可取消
	//back.SetBack(router)
	// 注册
	service.Handle("/", http.StripPrefix("/api",router))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
