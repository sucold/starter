package v7

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserUpdateReq struct {
	g.Meta    `path:"/user/update" method:"post" sm:"更新信息" tags:"用户管理"`
	Name      string `gorm:"column:name;type:text;not null" json:"name"`
	Password  string `json:"password" dc:"原密码" v:"required-with:Password1|length:8,30#请输入新密码|密码长度为8-30位"`
	Password1 string `json:"password1" dc:"新密码" v:"required-with:Password|length:8,30#请输入旧密码|密码长度为8-30位"`
}
type UserUpdateRes struct {
}

type UserGetReq struct {
	g.Meta `path:"/user/get" method:"get" sm:"获取信息" tags:"用户管理"`
}
type UserGetRes struct {
	ID        int64  `gorm:"column:id;type:integer;primaryKey" json:"id"`
	CreatedAt int64  `gorm:"column:created_at;type:integer;not null" json:"created_at"`
	UpdatedAt int64  `gorm:"column:updated_at;type:integer;not null" json:"updated_at"`
	Name      string `gorm:"column:name;type:text;not null" json:"name"`
	Email     string `gorm:"column:email;type:text;not null" json:"email"`
	Role      string `gorm:"column:role;type:text;not null" json:"role"`
}
