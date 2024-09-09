package app

import (
	"go-fiber/model/web"
	"go-fiber/model/web/response"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

func NewRouter() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "Raihanmd",
	})

	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(web.WebSuccessResponse[response.User]{
			Code:    http.StatusOK,
			Message: "OK",
			Payload: response.User{
				ID:    1,
				Name:  "Raihanmd",
				Email: "Raihanmd",
			},
		})
	})

	app.Get("/protected", func(c fiber.Ctx) error {
		auth := c.Get("Authorization")

		if auth != "admin" {
			return c.Status(fiber.StatusUnauthorized).JSON(web.WebErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized",
			})
		}

		return c.JSON(web.WebSuccessResponse[fiber.Map]{
			Code:    http.StatusOK,
			Message: "OK",
			Payload: fiber.Map{"message": "Hello, World!"},
		})
	})

	return app
}
