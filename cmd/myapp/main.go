package main

import (
	"go-api/config"
	"go-api/internal/controllers"
	"go-api/internal/messageBroker/rabbitmq"
	"go-api/internal/repositories/mysql"
	"go-api/internal/routes"
	"go-api/internal/services"
	"log"
	"os"
)

func main() {

	config.LoadConfig()
	//init mysql
	config.InitDatabase()
	//init redis
	config.InitRedis()
	// init gin
	router := config.InitGin()
	// type go_api_env struct {
	// 	db_driver       string
	// 	redis_url       string
	// 	db_username     string
	// 	db_password     string
	// 	redis_password  string
	// 	rdis_ttl_second int16
	// 	mongo_user      string
	// 	mongo_pass      string
	// }
	// register routes
	// db_driver := os.Getenv("db_driver")
	// if db_driver == "mysql" {

	// 	routes.RegisterUserRoutes(router, userController)
	// }
	//mongodb

	// initialize service (inject repository)

	mysqlRepo := mysql.NewMysqlRepository()
	userService := services.NewUserService(mysqlRepo)
	url := "amqp://admin:secret123@localhost:5672/"
	rabbitmq := rabbitmq.NewRabbitMQ(url)
	authService := services.NewAuthService(mysqlRepo, rabbitmq)
	userController := controllers.NewUserController(userService)
	authController := controllers.NewAuthController(authService)
	routes.RegisterUserRoutes(router, userController, authController)

	productMysqlRepo := mysql.NewProductMysqlRepository()
	productService := services.NewProductService(productMysqlRepo)
	productController := controllers.NewProductController(productService)
	routes.RegisterProductRoutes(router, productController)

	// routes.RegisterAllRoutes(router, userController)

	//run
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server running on port", port)
	router.Run(":" + port)

}
