package rest

import "github.com/gofiber/fiber/v2"

//health check
func CommonRoutes(r fiber.Router) fiber.Router {
	r.Get("/", HealthCheck)
	return r
}

func HandlerRoutes(app *fiber.App) {
	// 404 Handler
	app.Use(NotFound)
}

func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Server is healthy!",
		"data":    nil,
	})
}

func NotFound(c *fiber.Ctx) error {
	return c.SendStatus(404)
}
