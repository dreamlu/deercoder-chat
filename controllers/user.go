package controllers

import (
	"demo/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

//根据id获得用户获取
func GetUserById(u *gin.Context) {
	id := u.Query("id")
	ss := models.GetUserById(id)
	u.JSON(http.StatusOK, ss)
}

//用户信息分页
func GetUserBySearch(u *gin.Context) {
	u.Request.ParseForm()
	values := u.Request.Form //在使用之前需要调用ParseForm方法
	ss := models.GetUserBySearch(values)
	u.JSON(http.StatusOK, ss)
}

//用户信息删除
func DeleteUserById(u *gin.Context) {
	id := u.Param("id")
	ss := models.DeleteUserByid(id)
	u.JSON(http.StatusOK, ss)
}

//用户信息修改
func UpdateUser(u *gin.Context) {
	u.Request.ParseForm()
	values := u.Request.Form //在使用之前需要调用ParseForm方法
	ss := models.UpdateUser(values)
	u.JSON(http.StatusOK, ss)
}

//新增用户信息
func CreateUser(u *gin.Context) {
	u.Request.ParseForm()
	values := u.Request.Form //在使用之前需要调用ParseForm方法
	ss := models.CreateUser(values)
	u.JSON(http.StatusOK, ss)
}

/*//修改用户账号密码
func UpdateAccount(u *gin.Context) {

}*/
