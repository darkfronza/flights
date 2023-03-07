package main

import (
	"fmt"
	"os"

	"github.com/darkfronza/flights/controllers/flights"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/calculate", flights.SortFlights)

	listenAddr := ":8080"
	if port := os.Getenv("API_PORT"); port != "" {
		listenAddr = fmt.Sprintf(":%s", port)
	}

	app.Listen(listenAddr)
}
