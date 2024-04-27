package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/claudineyveloso/govita.git/cmd/api"
	"github.com/claudineyveloso/govita.git/cmd/db"
	"github.com/claudineyveloso/govita.git/internal/configs"
)

func main() {
	cfg := configs.Config{
		PublicHost: configs.Envs.PublicHost,
		Port:       configs.Envs.Port,
		DBUser:     configs.Envs.DBUser,
		DBPassword: configs.Envs.DBPassword,
		DBName:     configs.Envs.DBName,
	}
	db, err := db.NewPostgresSQLStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(fmt.Sprintf(":%s", configs.Envs.Port), db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB: Successfully connected!")
}
