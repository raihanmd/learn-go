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

	return app
}
