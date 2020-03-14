package database

import (
	"dbbook/pkg/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"xorm.io/xorm"
)

func Connect(config config.Database) (DB *xorm.Engine) {
	var e error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
		"utc",
	)

	DB, e = xorm.NewEngine("mysql", dsn)
	if e != nil {
		log.Fatalf("Failed to connect to database: %s", e)
	}

	DB.SetMaxOpenConns(1)
	DB.SetMaxIdleConns(1)
	DB.SetConnMaxLifetime(time.Hour)

	return DB
}
