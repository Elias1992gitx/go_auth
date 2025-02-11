package routes

import (
	"github.com/gofiber/fiber/v2"
	"JWT-Authentication-go/controllers"
)

func SetUpRoutes(app *fiber.App){
	app.Get("/", controllers.Hello)
}