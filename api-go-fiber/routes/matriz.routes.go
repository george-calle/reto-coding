package routes

import (
	"apimatriz/controllers"

	"github.com/gofiber/fiber/v2"
)

func MatrizRoutes(app *fiber.App) {
	app.Post("/matriz", controllers.PostMatriz)
	app.Post("/matriz-sinapi", controllers.PostMatriz_SinAPI)
}
