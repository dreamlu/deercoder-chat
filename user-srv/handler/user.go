package handler

import (
	"context"
	"deercoder-chat/user-srv/models"
	user "deercoder-chat/user-srv/proto"
	"errors"
	"github.com/dreamlu/go-tool"
	"github.com/dreamlu/go-tool/util/lib"
	"log"
	"time"
)

type UserService struct{}

//根据id获得用户获取
func (p *UserService) GetByID(ctx context.Context, req *user.ID, rsp *user.User) error {
	id := req.Id
	//der.DB.AutoMigrate(user.Response)
	//var data = models.User{}
	res := der.GetDataBySQL(&rsp, "select * from `user` where id = ?", id)
	if (res.MapData == lib.MapSuccess) || (res.MapData == lib.MapNoResult) {
		return nil
	}

	return errors.New(res.MapData.Msg.(string))
}

//用户信息分页
func (p *UserService) GetBySearch(ctx context.Context, req *user.Request, rsp *user.SearchData) error {
	params := make(map[string][]string)
	for k, v := range req.Params {
		params[k] = append(params[k], v)
	}
	res := der.GetDataBySearch(models.User{}, &rsp.User, "user", params)
	// sumPage data
	rsp.SumPage = res.Pager.SumPage
	log.Println(res)
	return nil
}

//用户信息删除
func (p *UserService) Delete(ctx context.Context, req *user.ID, rsp *user.Boolean) error {
	id := req.Id
	res := der.DeleteDataBySQL("delete from `user` where id = ?", id)
	rsp.Bool = false
	if res == lib.MapDelete {
		rsp.Bool = true
		return nil
	}
	return errors.New(res.Msg.(string))
}

//用户信息修改
func (p *UserService) Update(ctx context.Context, req *user.Request, rsp *user.Boolean) error {
	params := make(map[string][]string)
	for k, v := range req.Params {
		params[k] = append(params[k], v)
	}

	if _, ok := params["password"]; ok {
		params["password"][0] = der.AesEn(params["password"][0])
	}
	res := der.UpdateData("user", params)
	rsp.Bool = false
	if res == lib.MapUpdate {
		rsp.Bool = true
		return nil
	}
	log.Println(res)
	return errors.New(res.Msg.(string))
}

// 新增用户信息
func (p *UserService) Create(ctx context.Context, req *user.Request, rsp *user.ID) error {

	params := make(map[string][]string)
	for k, v := range req.Params {
		params[k] = append(params[k], v)
	}

	if _, ok := params["password"]; ok {
		params["password"][0] = der.AesEn(params["password"][0])
	}
	params["createtime"] = append(params["createtime"], time.Now().Format("2006-01-02 15:04:05"))
	res := der.CreateData("user", params)
	log.Println(res)
	if res == lib.MapCreate {
		rsp.Id = 0
		return nil
	}
	return errors.New(res.Msg.(string))
}
