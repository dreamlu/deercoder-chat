package controllers

import (
	"context"
	"deercoder-chat/user-srv/models"
	user "deercoder-chat/user-srv/proto"
	"github.com/dreamlu/deercoder-gin"
	"log"
	"time"
)

type UserService struct{}

//根据id获得用户获取
func (p *UserService) GetByID(ctx context.Context, req *user.ID, rsp *user.User) error {
	id := req.Id
	//deercoder.DB.AutoMigrate(user.Response)
	//var data = models.User{}
	deercoder.GetDataBySQL(&rsp, "select * from `user` where id = ?", id)
	//rsp.Name = data.Name
	return nil
}

//用户信息分页
func (p *UserService) GetBySearch(ctx context.Context, req *user.Request, rsp *user.SearchData) error {
	params := make(map[string][]string)
	for k, v := range req.Params {
		params[k] = append(params[k], v)
	}
	res := deercoder.GetDataBySearch(models.User{}, &rsp.User,"user", params)
	// test
	// will fix it
	rsp.SumPage = 10
	log.Println(res)
	return nil
}

//用户信息删除
func (p *UserService) Delete(ctx context.Context, req *user.ID, rsp *user.Boolean) error {
	id := req.Id
	deercoder.DeleteDataBySQL("delete from `user` where id = ?", id)
	rsp.Bool = true
	return nil
}

//用户信息修改
func (p *UserService) Update(ctx context.Context, req *user.Request, rsp *user.Boolean) error {
	params := make(map[string][]string)
	for k, v := range req.Params {
		params[k] = append(params[k], v)
	}

	if _,ok := params["password"]; ok {
		params["password"][0] = deercoder.AesEn(params["password"][0])
	}
	res := deercoder.UpdateData("user", params)
	rsp.Bool = true
	log.Println(res)
	return nil
}

//新增用户信息
func (p *UserService) Create(ctx context.Context, req *user.Request, rsp *user.ID) error {

	params := make(map[string][]string)
	for k, v := range req.Params {
		params[k] = append(params[k], v)
	}

	if _,ok := params["password"]; ok {
		params["password"][0] = deercoder.AesEn(params["password"][0])
	}
	params["createtime"] = append(params["createtime"], time.Now().Format("2006-01-02 15:04:05"))
	res := deercoder.CreateData("user", params)
	log.Println(res)
	return nil
}
