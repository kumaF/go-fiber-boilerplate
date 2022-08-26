package utils

import (
	"log"
	"github.com/gofiber/fiber/v2"
)

func StartServer(app *fiber.App) {
	log.Fatal(app.Listen(":3000"))
}