package main

import (
	"github.com/ilmovilDev/go-auth/config/db"
	"github.com/ilmovilDev/go-auth/config/env"
	"github.com/ilmovilDev/go-auth/internal/router"
)

func main() {

	envs, err := env.LoadConfig()
	if err != nil {
		panic("Failed to load configuration")
	}

	_, err = db.InitializePostgres(envs)
	if err != nil {
		panic("Failed to initialize database")
	}

	router.Initialize(envs)

}
