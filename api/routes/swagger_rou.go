package routes

import (
	"kkl-v2/bootstrap"
	"kkl-v2/core/utilities"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// SetupSwagger configures Swagger documentation for the Fiber app
func SetupSwagger(app *fiber.App, config *bootstrap.SwaggerConfig) {
	// Environment-based Swagger access
	app.Get("/docs/*", utilities.ApiKeyAuthAny, swagger.HandlerDefault)
}
