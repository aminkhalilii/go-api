package main

import (
	"go-api/config"
	"go-api/internal/routes"
	"log"
	"os"
)

func main() {

	//init mysql
	config.InitMysql()

	// init gin
	router := config.InitGin()

	// register routes
	routes.RegisterAllRoutes(router)

	//run
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server running on port", port)
	router.Run(":" + port)

}
