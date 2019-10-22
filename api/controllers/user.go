package controllers

import (
	"context"
	"deercoder-chat/api/conf"
	user "deercoder-chat/user-srv/proto"
	"github.com/dreamlu/go-tool/tool/result"
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
	res, err := UserClient.GetByID(context.TODO(), &user.ID{
		Id: id,
	})

	if err != nil {
		u.JSON(http.StatusOK, result.GetError(err.Error()))
		return
	}

	if res == nil {
		u.JSON(http.StatusOK, result.MapNoResult)
		return
	}

	u.JSON(http.StatusOK, result.GetSuccess(res))
}

//用户信息分页
func (p *UserService) GetBySearch(u *gin.Context) {
	_ = u.Request.ParseForm()
	values := u.Request.Form
	params := make(map[string]string)
	for k, v := range values {
		params[k] = v[0]
	}
	res, err := UserClient.GetBySearch(context.TODO(), &user.Request{
		Params: params,
	})

	if err != nil {
		u.JSON(http.StatusOK, result.GetError(err.Error()))
		return
	}

	if res == nil {
		u.JSON(http.StatusOK, result.MapNoResult)
		return
	}

	clientPage, _ := strconv.ParseInt(params["clientPage"], 10, 64)
	everyPage, _ := strconv.ParseInt(params["everyPage"], 10, 64)
	u.JSON(http.StatusOK, result.GetSuccessPager(res.User, result.Pager{
		ClientPage: clientPage,
		EveryPage:  everyPage,
		TotalNum:   res.SumPage,
	}))
}

//用户信息删除
func (p *UserService) Delete(u *gin.Context) {
	id, _ := strconv.ParseInt(u.Param("id"), 10, 64)
	_, err := UserClient.Delete(context.TODO(), &user.ID{
		Id: id,
	})

	if err != nil {
		u.JSON(http.StatusOK, result.GetError(err.Error()))
		return
	}

	u.JSON(http.StatusOK, result.MapUpdate)
}

//用户信息修改
func (p *UserService) Update(u *gin.Context) {
	_ = u.Request.ParseForm()
	values := u.Request.Form
	params := make(map[string]string)
	for k, v := range values {
		params[k] = v[0]
	}
	_, err := UserClient.Update(context.TODO(), &user.Request{
		Params: params,
	})

	if err != nil {
		u.JSON(http.StatusOK, result.GetError(err.Error()))
		return
	}

	u.JSON(http.StatusOK, result.MapUpdate)
}

//新增用户信息
func (p *UserService) Create(u *gin.Context) {
	_ = u.Request.ParseForm()
	values := u.Request.Form
	params := make(map[string]string)
	for k, v := range values {
		params[k] = v[0]
	}
	_, err := UserClient.Create(context.TODO(), &user.Request{
		Params: params,
	})

	if err != nil {
		u.JSON(http.StatusOK, result.GetError(err.Error()))
		return
	}

	u.JSON(http.StatusOK, result.MapCreate)
}
