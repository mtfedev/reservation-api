package api

import (
	"github.com/fulltimegodev/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
)

func HandlerGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "James",
		LastName:  "Foo",
	}
	return c.JSON(u)
}

func HandlerGetUser(c *fiber.Ctx) error {
	return c.JSON("James")
}
