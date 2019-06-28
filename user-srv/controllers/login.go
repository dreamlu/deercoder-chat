package controllers

import (
	"context"
	user "deercoder-chat/user-srv/proto"
	"errors"
	"github.com/dreamlu/go-tool"
	"github.com/dreamlu/go-tool/util/lib"
)

type LoginModel struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginService struct{}

//登录
func (p *LoginService) Login(ctx context.Context, req *user.LoginModel, rsp *user.LoginModel) error {
	var login LoginModel
	var sql string

	name := req.Name
	password := req.Password

	sql = "SELECT id,password FROM `user` WHERE name = ?"
	dba := der.DB.Raw(sql, name).Scan(&login)
	num := dba.RowsAffected
	if dba.Error == nil && num > 0 {
		password = der.AesEn(password)
		if login.Password == password {
			rsp.Id = login.ID

			var cache der.CacheManager = new(der.RedisManager)
			_ = cache.NewCache()
			//cacheModel := der.CacheModel{}
			// redis 存储用户信息
			_ = cache.Set(login.ID, der.CacheModel{
				Time: 30 * 2 * 24, // 有效期 一天
				Data: login.ID,
			})

			return nil
		} else {
			return errors.New(lib.MsgCountErr)
		}
	}
	return errors.New(lib.MsgError)
}
