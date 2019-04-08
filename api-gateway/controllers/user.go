package controllers

import (
	"context"
	"deercoder-chat/api-gateway/conf"
	"deercoder-chat/chat-srv/proto"
	"github.com/dreamlu/deercoder-gin/util/lib"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"net/http"
	"strconv"
)

type UserService struct{}

var (
	UserClient user.UserService
)

func init() {
	UserClient = user.NewUserService(conf.UserSrv, client.DefaultClient)
}

//根据id获得用户获取
func (p *UserService) GetByID(u *gin.Context) {
	id, _ := strconv.ParseInt(u.Param("id"), 10, 64)
	response, err := UserClient.GetByID(context.TODO(), &user.Request{
		Id: id,
	})

	if err != nil {
		u.JSON(http.StatusInternalServerError, err)
	}

	u.JSON(http.StatusOK, lib.GetInfoN{
		Status:lib.CodeSuccess,
		Msg: lib.MsgSuccess,
		Data: response,
	})
}

////用户信息分页
//func (p * UserService)GetBySearch(u *gin.Context) {
//	u.Request.ParseForm()
//	values := u.Request.Form //在使用之前需要调用ParseForm方法
//	ss := models.GetBySearch(values)
//	u.JSON(http.StatusOK, ss)
//}
//
////用户信息删除
//func (p * UserService)DeleteById(u *gin.Context) {
//	id := u.Param("id")
//	ss := models.DeleteByid(id)
//	u.JSON(http.StatusOK, ss)
//}
//
////用户信息修改
//func (p * UserService)Update(u *gin.Context) {
//	u.Request.ParseForm()
//	values := u.Request.Form //在使用之前需要调用ParseForm方法
//	ss := models.Update(values)
//	u.JSON(http.StatusOK, ss)
//}
//
////新增用户信息
//func (p * UserService)Create(u *gin.Context) {
//	u.Request.ParseForm()
//	values := u.Request.Form
//	ss := models.Create(values)
//	u.JSON(http.StatusOK, ss)
//}
