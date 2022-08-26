package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kumaF/go-fiber-boilerplate/routes"
	"github.com/kumaF/go-fiber-boilerplate/utils"
)

func main() {
	app := fiber.New()
	api := app.Group("/api")

	// v1.0 routes
	v1 := api.Group("/v1.0")
	routes.HealthRoutes(v1)

	utils.StartServer(app)
}