package main

import (
    "github.com/gofiber/fiber/v2"
	"Fiber/src/main/infra"
)

func main() {
    app := fiber.New()
    infra.ConnectDatabase()

    app.Listen(":3000")
}