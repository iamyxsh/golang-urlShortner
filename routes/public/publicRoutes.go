package publicRoutes

import (
	constants "urlShortner/constants"
	pc "urlShortner/controllers/public"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App){
	app.Post(constants.ShortenUrl, pc.ShortenUrl)
	app.Get(constants.RedirectToFull, pc.RedirectToFull)
}