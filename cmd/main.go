package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/ufguff/cmd/api"
	"github.com/ufguff/config"
	"github.com/ufguff/db"
	// http-swagger middleware
)

func main() {
	database, err := db.NewSQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	// проверка связи
	checkDB(database)

	server := api.NewApiServer(":8080", database)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func checkDB(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Succesful connection to DB!")
}
