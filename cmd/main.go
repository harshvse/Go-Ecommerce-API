package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/harshvse/golang-ecommerce-api/cmd/api"
	"github.com/harshvse/golang-ecommerce-api/config"
	"github.com/harshvse/golang-ecommerce-api/db"
)

func main() {
	db, _ := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               "ecom",
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	initStorage(db)

	server := api.NewAPIServer(config.Envs.Port, db)
	if err := server.Run(); err != nil {
		log.Fatal("[-]Server failed to start:", err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal("[-]Failed to Ping database", err)
	}
	log.Printf("[+]Database connected")
}
