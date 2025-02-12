package main

import (
	"JWT-Authentication-go/database"
	"JWT-Authentication-go/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		panic("could not connect to db: " + err.Error())
	}

	fmt.Println("Connection to MongoDB is successful:", db.Name())

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return true // Allows any origin
		},
		AllowMethods:     "GET,POST,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Content-Type,Authorization,Accept,Origin,Access-Control-Request-Method,Access-Control-Request-Headers,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Access-Control-Allow-Methods,Access-Control-Expose-Headers,Access-Control-Max-Age,Access-Control-Allow-Credentials",
		AllowCredentials: true,
	}))

	routes.SetUpRoutes(app)

	err = app.Listen(":8000")
	if err != nil {
		panic("could not start server: " + err.Error())
	}
}
