package main

import (
    "Fiber/src/main/infra"
    "github.com/gofiber/fiber/v2"

    "Fiber/src/main/app/api/products"

    // "time"

    _ "github.com/go-sql-driver/mysql"

    // jwtware "github.com/gofiber/contrib/jwt"
    // "github.com/golang-jwt/jwt/v5"
)

func main() {
    app := fiber.New()
    infra.ConnectDatabase()

    products.ProductRoutes(app)

    app.Listen(":3000")
}