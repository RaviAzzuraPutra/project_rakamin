package user_controller

import (
	"last-project/app/request"
	"last-project/app/service/interface/user_service_interface"

	"github.com/gofiber/fiber/v2"
)

type User_Controller struct {
	service user_service_interface.User_Service_Interface
}

func NewUserControllerRegistry(service user_service_interface.User_Service_Interface) *User_Controller {
	return &User_Controller{
		service: service,
	}
}

func (c *User_Controller) Get(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(string)

	user, err := c.service.Get(userID)

	if err != nil {
		return ctx.JSON(fiber.Map{
			"Status":  500,
			"Message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"Status":  200,
		"Message": "Success Get User",
		"Data":    user,
	})
}

func (c *User_Controller) Update(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(string)

	request := new(request.User_Request)

	errRequest := ctx.BodyParser(request)

	if errRequest != nil {
		return ctx.JSON(fiber.Map{
			"Status":  400,
			"Message": "Bad Request",
		})
	}

	user, errUpdate := c.service.Update(userID, request)

	if errUpdate != nil {
		return ctx.JSON(fiber.Map{
			"Status":  500,
			"Message": errUpdate.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"Status":  200,
		"Message": "Succes Update User",
		"Data":    user,
	})
}
