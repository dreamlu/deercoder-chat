package back

import (
	"deercoder-chat/api-gateway/back/datamodel"
	_ "github.com/chenhg5/go-admin/adapter/gin"
	"github.com/chenhg5/go-admin/engine"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/dreamlu/deercoder-gin"
	"github.com/gin-gonic/gin"
)

// 后台配置
func SetBack(router *gin.Engine) *gin.Engine {
	// 实例化一个go-admin引擎对象
	eng := engine.Default()
	// go-admin全局配置
	cfg := config.Config{
		DATABASE: []config.Database{
			{
				HOST:         deercoder.GetDevModeConfig("mysql.host"),
				PORT:         deercoder.GetDevModeConfig("mysql.port"),
				USER:         deercoder.GetDevModeConfig("db.user"),
				PWD:          deercoder.GetDevModeConfig("db.password"),
				NAME:         deercoder.GetDevModeConfig("db.name"),
				MAX_IDLE_CON: 50,
				MAX_OPEN_CON: 150,
				DRIVER:       "mysql",
			},
		},
		DOMAIN: deercoder.GetDevModeConfig("domain"), // 是cookie相关的，访问网站的域名
		PREFIX: "admin",
		// STORE 必须设置且保证有写权限，否则增加不了新的管理员用户
		STORE: config.Store{
			PATH:   "./uploads",
			PREFIX: "uploads",
		},
		LANGUAGE: "cn",
	}

	// Generators： 详见 https://github.com/chenhg5/go-admin/blob/master/examples/datamodel/tables.go
	adminPlugin := admin.NewAdmin(datamodel.Generators)

	// 增加配置与插件，使用Use方法挂载到Web框架中
	_ = eng.AddConfig(cfg).AddPlugins(adminPlugin).Use(router)
	return router
}
