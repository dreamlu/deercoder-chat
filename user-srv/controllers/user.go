package controllers

import (
	"context"
	"deercoder-chat/user-srv/models"
	user "deercoder-chat/user-srv/proto"
	"github.com/dreamlu/deercoder-gin"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserService struct{}

//根据id获得用户获取
func (p *UserService) GetByID(ctx context.Context, req *user.Request, rsp *user.Response) error {
	id := req.Id
	//deercoder.DB.AutoMigrate(user.Response)
	//var data = models.User{}
	deercoder.GetDataBySql(&rsp, "select * from `user` where id = ?", id)
	//rsp.Name = data.Name
	return nil
}

//用户信息分页
func GetBySearch(u *gin.Context) {
	u.Request.ParseForm()
	values := u.Request.Form //在使用之前需要调用ParseForm方法
	ss := models.GetBySearch(values)
	u.JSON(http.StatusOK, ss)
}

//用户信息删除
func DeleteById(u *gin.Context) {
	id := u.Param("id")
	ss := models.DeleteByid(id)
	u.JSON(http.StatusOK, ss)
}

//用户信息修改
func Update(u *gin.Context) {
	u.Request.ParseForm()
	values := u.Request.Form //在使用之前需要调用ParseForm方法
	ss := models.Update(values)
	u.JSON(http.StatusOK, ss)
}

//新增用户信息
func Create(u *gin.Context) {
	u.Request.ParseForm()
	values := u.Request.Form
	ss := models.Create(values)
	u.JSON(http.StatusOK, ss)
}
