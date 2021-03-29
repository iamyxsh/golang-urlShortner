package main

import (
	"urlShortner/database"
	adminRoutes "urlShortner/routes/admin"
	publicRoutes "urlShortner/routes/public"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	publicRoutes.PublicRoutes(app)
	
	adminRoutes.AdminRoutes(app)

	app.Listen(":5000")
}
