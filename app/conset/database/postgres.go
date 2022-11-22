//go:build postgres

package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDial(dsn string) gorm.Dialector { //sqlite
	return postgres.Open(dsn)
}
