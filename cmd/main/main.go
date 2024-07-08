package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/navindu-sachintha/go-bookstore/pkg/controllers"
	"github.com/navindu-sachintha/go-bookstore/pkg/routes"
)

func main() {

	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	handler := controllers.NewHandler()
	router := routes.SetupRouter(handler)

	router.Run()
}
