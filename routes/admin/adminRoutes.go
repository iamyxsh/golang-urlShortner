package adminRoutes

import (
	"urlShortner/constants"
	adminController "urlShortner/controllers/admin"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func AdminRoutes(app *fiber.App){
	app.Post(constants.AdminLogin, adminController.AdminLogin)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("superSecretKey"),
	}))

	app.Post(constants.ChangeUrlStatus, adminController.ChangeUrlStatus)


}