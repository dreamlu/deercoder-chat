package controllers

import (
	"deercoder-chat/user-srv"
	"github.com/gin-gonic/gin"
	"net/http"
)

//根据id获得用户获取
func GetById(u *gin.Context) {
	id := u.Query("id")
	ss := main.GetById(id)
	u.JSON(http.StatusOK, ss)
}

//用户信息分页
func GetBySearch(u *gin.Context) {
	u.Request.ParseForm()
	values := u.Request.Form //在使用之前需要调用ParseForm方法
	ss := main.GetBySearch(values)
	u.JSON(http.StatusOK, ss)
}

//用户信息删除
func DeleteById(u *gin.Context) {
	id := u.Param("id")
	ss := main.DeleteByid(id)
	u.JSON(http.StatusOK, ss)
}

//用户信息修改
func Update(u *gin.Context) {
	u.Request.ParseForm()
	values := u.Request.Form //在使用之前需要调用ParseForm方法
	ss := main.Update(values)
	u.JSON(http.StatusOK, ss)
}

//新增用户信息
func Create(u *gin.Context) {
	u.Request.ParseForm()
	values := u.Request.Form
	ss := main.Create(values)
	u.JSON(http.StatusOK, ss)
}
