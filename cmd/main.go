package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/rommomm123321/go_crud_api/config"
	"github.com/rommomm123321/go_crud_api/pkg/handler"
	"github.com/rommomm123321/go_crud_api/pkg/repository"
	"github.com/rommomm123321/go_crud_api/pkg/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config.LoadEnvVariables()
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!")
	}
	repository := repository.NewRepository(db)
	services := service.NewService(repository)
	handlers := handler.NewHandler(services)

	app := fiber.New()
	app.Mount("/api", handlers.InitRoutes())
	log.Fatal(app.Listen(":8800"))
}
