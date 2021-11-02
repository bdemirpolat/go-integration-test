package main

import (
	"github.com/bdemirpolat/integration-test/db"
	"github.com/bdemirpolat/integration-test/handlers/userhandler"
	"github.com/bdemirpolat/integration-test/repository"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database := db.ConnectDB()
	userRepo := &repository.UserRepo{DB: database}
	userHandler := &userhandler.UserHandler{Repo: userRepo}
	app := fiber.New()
	app.Post("/users", userHandler.Create)
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
