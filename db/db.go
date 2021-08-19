package db

import (
	"github.com/ccb1900/tinyweb/config"
	"github.com/ccb1900/tinyweb/db/mysql"
	"github.com/ccb1900/tinyweb/db/postgres"
	"github.com/ccb1900/tinyweb/log"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init()  {
	switch config.Get("db.connection") {
	case "mysql":
		db = mysql.DB()
	case "postgres":
		db = postgres.DB()
	default:
		log.Fatal("没有匹配的数据库连接")
	}
}

func GetDB() *gorm.DB  {
	return db
}
