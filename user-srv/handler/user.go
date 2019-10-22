package handler

import (
	"context"
	"deercoder-chat/user-srv/models"
	user "deercoder-chat/user-srv/proto"
	"github.com/dreamlu/go-tool"
	"time"
)

type UserService struct{}

//根据id获得用户获取
func (p *UserService) GetByID(ctx context.Context, req *user.ID, rsp *user.User) error {
	id := req.Id
	//der.DB.AutoMigrate(user.Response)
	//var data = models.User{}
	return gt.DBTooler().GetDataBySQL(&rsp, "select * from `user` where id = ?", id)
}

//用户信息分页
func (p *UserService) GetBySearch(ctx context.Context, req *user.Request, rsp *user.SearchData) error {
	params := make(map[string][]string)
	for k, v := range req.Params {
		params[k] = append(params[k], v)
	}
	pager, err := gt.DBTooler().GetDataBySearch(models.User{}, &rsp.User, "user", params)
	// sumPage data
	rsp.SumPage = pager.TotalNum
	//log.Println(res)
	return err
}

//用户信息删除
func (p *UserService) Delete(ctx context.Context, req *user.ID, rsp *user.Boolean) error {
	id := req.Id
	err := gt.DBTooler().DeleteDataBySQL("delete from `user` where id = ?", id)
	rsp.Bool = false
	if err == nil {
		rsp.Bool = true
		//return nil
	}
	return err
}

//用户信息修改
func (p *UserService) Update(ctx context.Context, req *user.Request, rsp *user.Boolean) error {
	params := make(map[string][]string)
	for k, v := range req.Params {
		params[k] = append(params[k], v)
	}

	if _, ok := params["password"]; ok {
		params["password"][0] = gt.AesEn(params["password"][0])
	}
	err := gt.DBTooler().UpdateData("user", params)
	rsp.Bool = false
	if err == nil {
		rsp.Bool = true
		//return nil
	}
	//log.Println(res)
	return err
}

// 新增用户信息
func (p *UserService) Create(ctx context.Context, req *user.Request, rsp *user.ID) error {

	params := make(map[string][]string)
	for k, v := range req.Params {
		params[k] = append(params[k], v)
	}

	if _, ok := params["password"]; ok {
		params["password"][0] = gt.AesEn(params["password"][0])
	}
	params["createtime"] = append(params["createtime"], time.Now().Format("2006-01-02 15:04:05"))
	err := gt.DBTooler().CreateData("user", params)
	//log.Println(res)
	if err == nil {
		rsp.Id = 0
		return nil
	}
	return err
}
