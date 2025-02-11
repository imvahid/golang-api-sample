package config

import (
	"fmt"
	"github.com/imvahid/golang-api-sample/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 3306
	database = "sample"
	username = "root"
	password = ""
)

func MySQLConnection() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, database,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.ErrorPanic(err)
	return db
}
