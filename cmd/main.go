package main

import "modular-acai-shop/pkg/database"

func main() {
	pg, err := database.NewPostgresDB("postgresql://root:example@localhost:5432/db")
	if err != nil {
		panic(err)
	}
	defer pg.Close()

}
