package postgres

import (
	"fmt"
	"github.com/ccb1900/tinyweb/config"
	"github.com/ccb1900/tinyweb/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func dsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		config.Get("db.postgres.host"),
		config.Get("db.postgres.username"),
		config.Get("db.postgres.password"),
		config.Get("db.postgres.database"),
		config.Get("db.postgres.port"),
		config.Get("app.timezone"),
	)
}

func DB() *gorm.DB  {
	db, err := gorm.Open(postgres.Open(dsn()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
