package main

import (
	"Todo-list/routes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("--Error loading .env file--")
	}
	server := fiber.New()
	routes.Router(server)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	server.Listen(":" + port)
}
