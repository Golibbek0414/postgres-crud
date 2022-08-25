package main

import (
	"postgres-gin-crud/config"
	"postgres-gin-crud/postgres"
	"postgres-gin-crud/postgres/migrations"
	"postgres-gin-crud/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	db, err := postgres.Connect(cfg)
	if err != nil {
		panic(err)
	}
	repo := migrations.NewDb(db)
    server.NewRouter(repo)
}
