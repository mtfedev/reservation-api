package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("litebAddr", ":5000", "The listen adders of the API server")
	flag.Parse()
	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	app.Get("/foo", handleFoo)

	apiv1.Get("/user", handleUser)
	app.Listen(*listenAddr)
}

func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "OK"})
}

func handleUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "Bob Foo"})
}
