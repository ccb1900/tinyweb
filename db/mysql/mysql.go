package mysql

import (
	"fmt"
	"github.com/ccb1900/tinyweb/config"
	"github.com/ccb1900/tinyweb/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Get("db.mysql.username"),
		config.Get("db.mysql.password"),
		config.Get("db.mysql.host"),
		config.Get("db.mysql.port"),
		config.Get("db.mysql.database"),
		)
}

func DB() *gorm.DB  {
	db, err := gorm.Open(mysql.Open(dsn()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
