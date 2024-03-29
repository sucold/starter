// Code generated by github.com/hinego/gen. DO NOT EDIT.
// Code generated by github.com/hinego/gen. DO NOT EDIT.
// Code generated by github.com/hinego/gen. DO NOT EDIT.

package model

const TableNameLog = "logs"

// Log mapped from table <logs>
type Log struct {
	ID        int64  `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt int64  `gorm:"column:created_at;type:bigint" json:"created_at"`
	UpdatedAt int64  `gorm:"column:updated_at;type:bigint" json:"updated_at"`
	UserID    int64  `gorm:"column:user_id;type:bigint" json:"user_id"`
	Type      string `gorm:"column:type;type:text" json:"type" dc:"日志类型"`
	LinkID    int64  `gorm:"column:link_id;type:bigint" json:"link_id"`
	Name      string `gorm:"column:name;type:text" json:"name"`
	Content   string `gorm:"column:content;type:text" json:"content"`
}

// TableName Log's table name
func (*Log) TableName() string {
	return TableNameLog
}
