package api

import (
	"github.com/gofiber/fiber/v2"
)

type ApiPlugin struct {
	app *fiber.App
}

func Config() ApiPlugin {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	return ApiPlugin{
		app,
	}
}
