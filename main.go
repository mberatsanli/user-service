package main

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/gofiber/fiber/v2"
	"log"
	"user/database"
	"user/routes"
)

func main() {
	myFigure := figure.NewColorFigure("User Service", "", "green", true)
	myFigure.Print()

	database.ConnectDb()

	log.Println("Starting User Service")

	// Initialize service
	InitService()

	log.Println("Starting Fiber App")
	// Initialize a new Fiber app
	app := fiber.New()
	routes.SetupApiRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	// Start the server on port 8080
	go log.Fatal(app.Listen(":8080"))
}
