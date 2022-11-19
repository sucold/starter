package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hinego/conset/base"
	"github.com/hinego/starter/app/model"
	"time"
)

type AuthLoginReq struct {
	g.Meta   `path:"/auth/login" method:"post" sm:"登录账号" tags:"身份验证"`
	Username string `json:"username" v:"required|email#请输入邮箱|请输入正确的邮箱地址" dc:"邮箱"`
	Password string `json:"password" v:"required|length:8,32#请输入密码|密码长度为{min}到{max}位" dc:"密码"`
}
type AuthLoginRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type AuthUserReq struct {
	g.Meta `path:"/auth/user" method:"get" sm:"用户信息" tags:"身份验证"`
}
type AuthUserRes *model.User

type AuthRegisterReq struct {
	g.Meta   `path:"/auth/register" method:"post" sm:"注册账号" tags:"身份验证"`
	Username string `json:"username" v:"required|email#请输入邮箱|请输入正确的邮箱地址" dc:"邮箱"`
	Password string `json:"password" v:"required|length:8,32#请输入密码|密码长度为{min}到{max}位" dc:"密码"`
	Code     string `json:"code" v:"required|length:6,6#请输入验证码|验证码长度为{min}位" dc:"验证码"`
}
type AuthRegisterRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}
type AuthLogoutReq struct {
	g.Meta `path:"/auth/logout" method:"post" sm:"注销登录" tags:"身份验证"`
}
type AuthLogoutRes struct{}
type AuthSessionReq struct {
	g.Meta `path:"/auth/session" method:"post" sm:"登录记录" tags:"身份验证"`
}
type AuthLoginUser struct {
	ID        int64  `gorm:"column:id;type:integer;primaryKey" json:"id"`
	CreatedAt int64  `gorm:"column:created_at;type:integer;not null" json:"created_at"`
	UpdatedAt int64  `gorm:"column:updated_at;type:integer;not null" json:"updated_at"`
	Name      string `gorm:"column:name;type:text;not null" json:"name"`
	Email     string `gorm:"column:email;type:text;not null" json:"email"`
	Role      string `gorm:"column:role;type:text;not null" json:"role"`
}
type AuthSessionRes struct {
	User    AuthLoginUser `json:"user"`
	Setting base.Setting  `json:"setting"`
}
type AuthSendReq struct {
	g.Meta `path:"/auth/send" method:"post" sm:"发验证码" tags:"身份验证"`
	Mail   string `json:"mail" v:"required|email#请输入邮箱|请输入正确的邮箱地址" dc:"邮箱"` //邮箱
	Type   string `json:"type"`
}
type AuthSendRes struct{}
type AuthRefreshReq struct {
	g.Meta `path:"/auth/refresh" method:"get" sm:"会话续签" tags:"身份验证"`
}
type AuthRefreshRes struct {
	Expire int64 `json:"expire"`
}
type AuthForgetReq struct {
	g.Meta   `path:"/auth/forget" method:"post" sm:"忘记密码" tags:"身份验证"`
	Username string `json:"username" v:"required|email#请输入邮箱|请输入正确的邮箱地址" dc:"邮箱"`
	Password string `json:"password" v:"required|length:8,32#请输入密码|密码长度为{min}到{max}位" dc:"密码"`
	Code     string `json:"code" v:"required|length:6,6#请输入验证码|验证码长度为{min}位" dc:"验证码"`
}
type AuthForgetRes struct{}

type AuthResetReq struct {
	g.Meta   `path:"/auth/reset" method:"post" sm:"重置密码" tags:"身份验证"`
	Password string `json:"password" v:"required|length:8,32#请输入密码|密码长度为{min}到{max}位" dc:"密码"`
}
type AuthResetRes struct{}
