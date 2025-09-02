package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"
)

func NewFiber() *fiber.App {
	fiberApp := fiber.New(fiber.Config{
		BodyLimit: 10 * 1024 * 1024,
	})
	fiberApp.Use(helmet.New())
	fiberApp.Use(logger.New())
	fiberApp.Use(recover.New())
	fiberApp.Use(cors.New())
	return fiberApp
}
