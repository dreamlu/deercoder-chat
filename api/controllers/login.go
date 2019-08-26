package controllers

import (
	"context"
	"deercoder-chat/api/conf"
	"deercoder-chat/user-srv/proto"
	"github.com/dreamlu/go-tool/util/lib"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"net/http"
)

type LoginService struct{}

var (
	loginClient user.LoginService
)

func init() {
	loginClient = user.NewLoginService(conf.UserSrv, client.DefaultClient)
}

// 登录
func (p *LoginService) Login(u *gin.Context) {
	name := u.PostForm("name")
	password := u.PostForm("password")

	res, err := loginClient.Login(context.TODO(), &user.LoginModel{
		Name:     name,
		Password: password,
	})

	if err != nil {

		if err.Error() == lib.MsgCountErr {
			u.JSON(http.StatusOK, lib.MapNoCount)
			return
		}

		u.JSON(http.StatusOK, lib.GetMapDataError(err.Error()))
		return
	}

	u.JSON(http.StatusOK, lib.GetMapDataSuccess(res))
}
