package database

import (
	"fmt"
	"log"

	"github.com/erfanyousefi/simple-article/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GORMDB *gorm.DB

func ConnectGORM(cfg config.Config) {
	var dns string
	var dialector gorm.Dialector
	if cfg.Driver == "mysql" {
		dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)
		dialector = mysql.Open(dns)
	} else if cfg.Driver == "postgres" {
		dns = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=diable", cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort)
		// dialector = postgres.Open(dns)
	}

	var err error
	GORMDB, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatal("cannot connected to DB ❌")
	}
	log.Println("connected to DB successfully ✅")
}
