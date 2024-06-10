package main

import (
	"log"
	"os"

	"apimatriz/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	routes.MatrizRoutes(app)

	app.Get("/api-docs/*", swagger.HandlerDefault)

	app.Get("/api-docs/*", func(c *fiber.Ctx) error {
        return c.SendFile("./swagger.yaml")
    })

	port := os.Getenv("PORT")
	if port == "" {
		port = "2000"
	}

	log.Fatal(app.Listen(":" + port))
}
