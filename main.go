package main

import (
	"deercder-chat/routers"
	"github.com/Dreamlu/deercoder-gin"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	router := routers.SetRouter()
	// Listen and Server in 0.0.0.0:8080
	router.Run(":" + deercoder.GetConfigValue("http_port"))
}
