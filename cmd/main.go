package main

import (
	"context"
	"log"
	"os"

	"github.com/devdutt6/ipo-tracker-go/handlers"
	"github.com/devdutt6/ipo-tracker-go/mongoutil"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Could not load env File")
	}

	defer mongoutil.GetConnection().Disconnect(context.Background())
	// auth apis
	app.Post("/register", handlers.RegisterHandler)
	app.Post("/login", handlers.LoginHandler)

	// middleware authenticate
	app.Use(handlers.Authenticate)

	// pan routes
	app.Route("/pan", func(router fiber.Router) {
		router.Get("", handlers.GetPanHandler)
		router.Post("", handlers.AddPanHandler)
		router.Delete("", handlers.DeletePanHandler)
	}, "panRouter")

	// check allotment
	app.Get("/checkAllotments/:companyId", handlers.CheckAllotmentHandler)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}

/**
ABDPN8414B
AEWPM4993E
CEWPS4143E
ADCPS7802F
*/
