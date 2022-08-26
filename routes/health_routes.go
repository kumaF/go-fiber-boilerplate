package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kumaF/go-fiber-boilerplate/controllers"
)

func HealthRoutes(router fiber.Router) {
	router.Get("/health", controllers.CheckHealth)
}