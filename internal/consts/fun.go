package consts

import (
	"github.com/hinego/starter/internal/dao"
	"gorm.io/gen/field"
)

func Sort(sortMap map[string]string, sorter map[string][2]field.Expr) field.Expr {
	for k, v := range sortMap {
		if st, ok := sorter[k]; ok {
			if v == "ascend" {
				return st[0]
			}
			if v == "descend" {
				return st[1]
			}
		}
	}
	return dao.Bin.ID.Desc()
}
