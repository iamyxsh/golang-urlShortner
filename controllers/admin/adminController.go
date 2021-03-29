package adminController

import (
	"fmt"
	errHandling "urlShortner/error"
	"urlShortner/helper"
	"urlShortner/schema"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
)


func AdminLogin(c *fiber.Ctx) error {
	loginReq := new(schema.AdminLoginReq)

	e := c.BodyParser(loginReq)
	errHandling.HandleErr(e)

	fmt.Println(loginReq.Username)

	err, res:= schema.FindAdmin(loginReq.Username)

	if err != nil{
		return c.SendString(err.Error())
	}

	if loginReq.Password != res.Password {
		return c.SendString("Admin Not Found.")
	}

	token, err := helper.CreateToken(res.IDField.ID)

	errHandling.HandleErr(err)

	return c.JSON(fiber.Map{
		"user" : res,
		"token": token, 
	})
}

func ChangeUrlStatus(c *fiber.Ctx) error {
	short := c.Params("short")
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	user_id := claims["user_id"].(string)

	if user_id != "6061eb7ce08fd7f347ea92a8"{
		return c.SendString("UnAuthorized Access")
	}

	err, res := schema.FindByShort(short)

	if err != nil {
		return c.SendString("Short Not Found.")
	}

	res.Status = !res.Status
	err = mgm.Coll(res).Update(res)

	if err != nil {
		return c.SendString("Short Not Found.")
	}

	return c.JSON(fiber.Map{
		"msg": "Status Changed Successfully",
		"status": res.Status,
	})
	
}