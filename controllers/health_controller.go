package controllers

import (
	"log"
	"github.com/gofiber/fiber/v2"
)

func CheckHealth(c *fiber.Ctx) error {
	log.Println("hello world")

	return c.SendStatus(fiber.StatusOK)
}