package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Application struct {
	Env   *Env
	Fiber *fiber.App
	DB    *gorm.DB
	Redis *redis.Client
	Swagger *SwaggerConfig
}

var GlobalEnv Env

func App() *Application {
	app := &Application{}
	app.Env = NewEnv()
	GlobalEnv = *NewEnv()
	app.Fiber = NewFiber()
	app.DB = NewDatabaseConnection(app.Env)
	app.Redis = InitializeRedis(app.Env)
	app.Swagger = NewSwaggerConfig()
	return app
}
