package pc

import (
	errHandling "urlShortner/error"
	"urlShortner/helper"
	schema "urlShortner/schema"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
)

func ShortenUrl(c *fiber.Ctx) error {
	url := new(schema.UrlReq)

	e := c.BodyParser(url)
	errHandling.HandleErr(e)

	res := schema.NewUrl( helper.ShortUrl(), url.Url, 0)

	e = mgm.Coll(res).Create(res)

	errHandling.HandleErr(e)

	return c.JSON(res)
}


func RedirectToFull(c *fiber.Ctx) error{
	short := c.Params("short")

    err, res := schema.FindByShort(short)

	if err != nil {
		return c.SendString(err.Error())
	}
	res.Clicks++

	err = mgm.Coll(res).Update(res)

	if err != nil {
		return c.SendString(err.Error())
	} 

	if res.Status == true{
		return c.Redirect(res.Full)
	} else {
		return c.JSON(fiber.Map{
			"msg": `Admin hast turned the status to "OFF"`,
		})
	}
	
}