package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// type HttpRoute struct {
// 	method     string
// 	path       string
// 	controller func(c *fiber.Ctx) error
// }

// this is a source component, try to came up with a clever solution
func (api *ApiPlugin) HttpListener(port string, domain string, method string, path string /*, controller func() interface{}*/) error {

	api.app.Add(method, path, func(c *fiber.Ctx) error {
		//resp := controller()
		return c.JSON("HELLO WORLD")
	})

	log.Fatal(api.app.Listen(domain + ":" + port))

	return nil
}
