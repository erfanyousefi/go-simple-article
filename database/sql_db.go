package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/erfanyousefi/simple-article/config"
)

var SQLDB *sql.DB

func ConnectSqlDB(cfg config.Config) {
	var dns string
	if cfg.Driver == "mysql" {
		// username:password@tcp(host:port)/dbname?parseTime=true
		dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)
	} else {
		// postgres://username:password@host:port/dbname?sslmode=disable
		dns = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)
	}

	var err error
	SQLDB, err := sql.Open(cfg.Driver, dns)
	if err != nil {
		log.Fatal("error in open connection with DB ❌")
	}

	if err := SQLDB.Ping(); err != nil {
		log.Fatal("database connection not available ❌")
	}

	log.Println("connected to database successfully ✅")
}
