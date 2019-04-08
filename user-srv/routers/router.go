package routers

import (
	"deercoder-chat/user-srv/controllers"
	"github.com/dreamlu/deercoder-gin"
	"github.com/dreamlu/deercoder-gin/util/file"
	"github.com/dreamlu/deercoder-gin/util/lib"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SetRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	router := gin.New()
	deercoder.MaxUploadMemory = router.MaxMultipartMemory
	//router.Use(CorsMiddleware())

	//router.Use(CheckLogin()) //简单登录验证

	// load the casbin model and policy from files, database is also supported.
	//权限中间件
	//e := casbin.NewEnforcer("conf/authz_model.conf", "conf/authz_policy.csv")
	//router.Use(controllers.NewAuthorizer(e))

	//静态目录
	router.Static("api/v1/static", "static")

	// Ping test
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	//组的路由,version
	v1 := router.Group("/api/v1")
	{
		v := v1
		//用户登录
		v.POST("/login/login", controllers.Login)
		//文件上传
		v.POST("/file/upload", file.UpoadFile)

		users := v.Group("/user")
		{
			users.POST("/create", controllers.Create)
			users.PUT("/update", controllers.Update)
			users.DELETE("/delete", controllers.DeleteById)
			users.GET("/search", controllers.GetBySearch)
			//users.GET("/id/:id", controllers.GetById)
		}
	}
	//不存在路由
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"msg":    "接口不存在->('.')/请求方法错误",
		})
	})
	return router
}

/*登录失效验证*/
func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		//小程序端
		if c.Request.Header.Get("Authorization") == "wechat" {
			return
		}
		path := c.Request.URL.String()
		if !strings.Contains(path, "login") && !strings.Contains(path, "/static") {
			_, err := c.Cookie("uid")
			if err != nil {
				c.Abort()
				c.JSON(http.StatusOK, lib.MapNoAuth)
			}
		}
	}
}
