package main

import (
	"modular-acai-shop/cmd/factory"
)

func main() {
	app := factory.NewApplication()
	app.LoadEnv()
	factory.RegisterRoutes(app)
	defer app.Pg.Close()
	app.RunServer(":8080")

}
