// Code generated by github.com/hinego/gen. DO NOT EDIT.
// Code generated by github.com/hinego/gen. DO NOT EDIT.
// Code generated by github.com/hinego/gen. DO NOT EDIT.

package model

import "github.com/hinego/decimal"

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID        int64           `gorm:"column:id;type:integer;primaryKey" json:"id"`
	CreatedAt int64           `gorm:"column:created_at;type:integer;not null" json:"created_at"`
	UpdatedAt int64           `gorm:"column:updated_at;type:integer;not null" json:"updated_at"`
	Name      string          `gorm:"column:name;type:text;not null" json:"name"`
	Email     string          `gorm:"column:email;type:text;not null" json:"email"`
	Password  string          `gorm:"column:password;type:text;not null" json:"password"`
	Balance   decimal.Decimal `gorm:"column:balance;type:numeric;not null" json:"balance" dc:"用户余额"`
	Refer     int64           `gorm:"column:refer;type:integer;not null" json:"refer"`
	IP        string          `gorm:"column:ip;type:text;not null" json:"ip" dc:"注册IP"`
	Role      string          `gorm:"column:role;type:text;not null" json:"role"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}