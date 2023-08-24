package main

import "github.com/gofiber/fiber/v2"

func handleTest(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"Message": "all good!",
	})
}

func main() {
	app := fiber.New()
	app.Get("/", handleTest)
	app.Listen(":3000")
}
