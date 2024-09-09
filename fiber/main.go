package main

import (
	"go-fiber/app"
)

func main() {
	app := app.NewRouter()

	app.Listen(":3000")
}
