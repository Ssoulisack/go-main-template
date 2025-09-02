package main

import (
	"fmt"
	"kkl-v2/api/routes"
	"kkl-v2/bootstrap"
	"log"
)

func main() {
	app := bootstrap.App()
	globalEnv := app.Env
	fiber := app.Fiber
	db := app.DB
	rd := app.Redis
	routes.Setup(fiber, db, rd)

	log.Fatal(fiber.Listen(fmt.Sprintf(":%d", globalEnv.App.Port)))
}
