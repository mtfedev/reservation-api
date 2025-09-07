package main

import (
	"flag"

	"github.com/fulltimegodev/hotel-reservation/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("litebAddr", ":5000", "The listen adders of the API server")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", api.HandlerGetUsers)
	apiv1.Get("/user/id", api.HandlerGetUser)
	app.Listen(*listenAddr)
}
