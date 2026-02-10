package main

import (
	"log"
	"os"

	"github.com/denvyworking/shorten-url-fiber-redis/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	// Загружаем .env (игнорируем ошибку в продакшене)
	_ = godotenv.Load()

	app := fiber.New()
	app.Use(logger.New())
	setupRoutes(app)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	// Слушаем на всех интерфейсах: 0.0.0.0:<port>
	log.Printf("🚀 Starting server on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
