package controllers

import (
	"context"
	user "deercoder-chat/user-srv/proto"
	"errors"
	"github.com/dreamlu/deercoder-gin"
	"github.com/dreamlu/deercoder-gin/util/lib"
)

type LoginModel struct {
	ID       int64    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginService struct{}

//登录
func (p *LoginService) Login(ctx context.Context, req *user.LoginModel, rsp *user.LoginModel) error {
	var user LoginModel
	var sql string

	name := req.Name
	password := req.Password

	sql = "SELECT id,password FROM `user` WHERE name = ?"
	dba := deercoder.DB.Raw(sql, name).Scan(&user)
	num := dba.RowsAffected
	if dba.Error == nil && num > 0 {
		password = deercoder.AesEn(password)
		if user.Password == password {
			rsp.Id = user.ID
			return nil
			//strID := strconv.Itoa(user.ID)
			//key := "-Iloveyouchinese"
			//user_id, err := deercoder.Encrypt([]byte(strID + string([]byte(key)[:len(key)-len(strID)])))
			//if err != nil {
			//	fmt.Println("cookie加密错误Encrypt: ", err)
			//	return
			//}
			//u.SetCookie("uid", string(user_id), 24*3600, "/", "*", false, true)
			//info = map[string]interface{}{"status": lib.CodeSuccess, "msg": "请求成功", "uid": strconv.Itoa(user.ID)}
		} else {
			return errors.New(lib.MsgCountErr)
		}
	}
	return errors.New(lib.MsgError)
}
