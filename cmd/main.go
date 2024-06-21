package main

import (
	"fmt"
	"modular-acai-shop/cmd/factory"
	"net/http"
)

func main() {
	app := factory.NewApplication()
	app.LoadEnv()
	factory.RegisterRoutes(app)

	defer app.Pg.Close()

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)

}
