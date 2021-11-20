package main

import (
	"flag"
	"log"

	"github.com/seefmitrais/go-rest-api-practice/internal/config"
	"github.com/seefmitrais/go-rest-api-practice/internal/postgresql"
	"github.com/seefmitrais/go-rest-api-practice/internal/repository"
	"github.com/seefmitrais/go-rest-api-practice/internal/rest"
	"github.com/seefmitrais/go-rest-api-practice/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	var envFileName string
	flag.StringVar(&envFileName, "env", "", "Environment Variable filename")
	flag.Parse()
	config.New(envFileName)
	postgresql.CreateConnection()
	DB := postgresql.DB

	app := fiber.New()
	app.Use(cors.New())
	version := app.Group("/api/v1")
	rest.CommonRoutes(version)

	//setup repositories
	userRepository := repository.NewUserRepository(DB)

	//setup services
	userService := service.NewUserService(userRepository)

	//setup rest handlers
	userHandler := rest.NewUserHandler(userService)

	userHandler.RegisterRoutes(version)

	rest.HandlerRoutes(app)
	log.Fatal(app.Listen(":3000"))

	defer DB.Close()

}
