package controllers

import (
	"fmt"
	"github.com/Dreamlu/deercoder-gin"
	"github.com/Dreamlu/deercoder-gin/util/lib"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Account struct {
	ID       int    `json:"id"`
	Account  string `json:"account"`
	Password string `json:"password"`
	ParentID int    `json:"parent_id"`
}

//登录
func Login(u *gin.Context) {
	var info interface{}
	var user Account
	var sql string

	account := u.PostForm("account")
	password := u.PostForm("password")

	sql = "SELECT id,password,parent_id FROM `teacher` WHERE account = ?"
	dba := deercoder.DB.Raw(sql, account).Scan(&user)
	num := dba.RowsAffected
	if dba.Error == nil && num > 0 {
		password = deercoder.AesEn(password)
		if user.Password == password {
			//name	string	cookie_key
			//value	string	cookie_val
			//maxAge	int	生存期（秒）
			//path	string	有效域
			//domain	string	有效域名
			//secure	bool	是否安全传输 是则只走https
			//httpOnly	bool	是否仅网络使用 是则js无法获取
			// encript
			//字符为16的倍数
			strID := strconv.Itoa(user.ID)
			key := "-Iloveyouchinese"
			user_id, err := deercoder.Encrypt([]byte(strID + string([]byte(key)[:len(key)-len(strID)])))
			if err != nil {
				fmt.Println("cookie加密错误Encrypt: ", err)
				return
			}
			//u.SetCookie("uid", strconv.Itoa(user.Id), 1800, "/", "*", false, true)
			u.SetCookie("uid", string(user_id), 24*3600, "/", "*", false, true)
			//this.SetSecureCookie(beego.AppConfig.String("secertkey"), "uid", strconv.Itoa(v.Id))
			if user.ParentID != 0 {
				user.ParentID = 1
			}
			info = map[string]interface{}{"status": lib.CodeSuccess, "msg": "请求成功", "uid": strconv.Itoa(user.ID), "flag": user.ParentID}
		} else {
			info = lib.MapCountErr
		}
	} else {
		info = lib.MapNoCount
	}

	u.JSON(http.StatusOK, info)
}
