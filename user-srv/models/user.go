package models

import (
	"github.com/dreamlu/go-tool"
)

/*user model*/
type User struct {
	ID         uint               `json:"id" gorm:"primary_key"`
	Name       string             `json:"name"`       //姓名
	Headimg    string             `json:"headimg"`    //头像
	Password   string             `json:"password"`   //密码
	Createtime der.JsonTime `json:"createtime"` //maybe you like util.JsonDate
}