package main

import (
	"context"
	"flag"
	"log"

	"github.com/fulltimegodev/hotel-reservation/api"
	"github.com/fulltimegodev/hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb//localhost:27017"
const dbname = "hotel-reservation"
const userColl = "users"

func main() {
	listenAddr := flag.String("litebAddr", ":5000", "The listen adders of the API server")
	flag.Parse()

	client, err := mongo.Connect((context.TODO()), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", userHandler.HandlerGetUsers)
	apiv1.Get("/user/id", userHandler.HandlerGetUser)
	app.Listen(*listenAddr)
}
