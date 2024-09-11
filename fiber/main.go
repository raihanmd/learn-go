package main

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	if !fiber.IsChild() {
		fmt.Println("I'm the parent process")
	} else {
		fmt.Println("I'm a child process")
	}

	app.Listen(":3000", fiber.ListenConfig{
		EnablePrefork: true,
	})
}
