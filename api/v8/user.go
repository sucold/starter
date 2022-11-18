package v8

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hinego/decimal"
	"github.com/hinego/starter/app/model"
	"github.com/hinego/types"
)

type FetchUser struct {
	ID        int64           `gorm:"column:id;type:integer;primaryKey" json:"id"`
	CreatedAt int64           `gorm:"column:created_at;type:integer;not null" json:"created_at"`
	UpdatedAt int64           `gorm:"column:updated_at;type:integer;not null" json:"updated_at"`
	Name      string          `gorm:"column:name;type:text;not null" json:"name"`
	Email     string          `gorm:"column:email;type:text;not null" json:"email"`
	Balance   decimal.Decimal `gorm:"column:balance;type:numeric;not null" json:"balance" dc:"用户余额"`
	Refer     int64           `gorm:"column:refer;type:integer;not null" json:"refer"`
	IP        string          `gorm:"column:ip;type:text;not null" json:"ip" dc:"注册IP"`
	Role      string          `gorm:"column:role;type:text;not null" json:"role"`
}
type UserFetchReq struct {
	g.Meta `path:"/user/fetch" method:"get" sm:"用户列表" tags:"用户管理"`
	types.PageReq
}
type UserFetchRes struct {
	*types.PageReq
	*types.PageRes
	Data []*FetchUser `json:"data"`
}
type UserUpdateReq struct {
	g.Meta   `path:"/user/update" method:"post" sm:"更新用户" tags:"用户管理"`
	ID       int64            `gorm:"column:id;type:integer;primaryKey" json:"id"`
	Name     string           `gorm:"column:name;type:text;not null" json:"name"`
	Password string           `json:"password" dc:"新密码"`
	Balance  *decimal.Decimal `gorm:"column:balance;type:numeric;not null" json:"balance" dc:"用户余额"`
	Role     string           `gorm:"column:role;type:text;not null" json:"role"`
}
type UserUpdateRes struct {
}

type UserCreateReq struct {
	g.Meta `path:"/user/create" method:"post" sm:"更新用户" tags:"用户管理"`
	model.User
}
type UserCreateRes struct {
}
