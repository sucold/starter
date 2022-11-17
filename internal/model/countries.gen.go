// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameCountry = "countries"

// Country mapped from table <countries>
type Country struct {
	ID        int64  `gorm:"column:id;type:integer;primaryKey" json:"id"`
	CreatedAt int64  `gorm:"column:created_at;type:integer;not null" json:"created_at"`
	UpdatedAt int64  `gorm:"column:updated_at;type:integer;not null" json:"updated_at"`
	Country   string `gorm:"column:country;type:text;not null" json:"country"`
	Country2  string `gorm:"column:country2;type:text;not null" json:"country2"`
	Country3  string `gorm:"column:country3;type:text;not null" json:"country3"`
	CountryCn string `gorm:"column:country_cn;type:text;not null" json:"country_cn"`
	Emoji     string `gorm:"column:emoji;type:text;not null" json:"emoji"`
	Logo      string `gorm:"column:logo;type:text;not null" json:"logo"`
	Symbol    string `gorm:"column:symbol;type:text;not null" json:"symbol"`
	Pass      string `gorm:"column:pass;type:text;" json:"pass"`
}

// TableName Country's table name
func (*Country) TableName() string {
	return TableNameCountry
}
