package auth_controller

import (
	"last-project/app/request"
	"last-project/app/service/interface/auth_service_interface"

	"github.com/gofiber/fiber/v2"
)

type Auth_Controller struct {
	Service auth_service_interface.Auth_Interface_Service
}

func NewAuthControllerRegistry(service auth_service_interface.Auth_Interface_Service) *Auth_Controller {
	return &Auth_Controller{
		Service: service,
	}
}

func (c *Auth_Controller) Register(ctx *fiber.Ctx) error {
	request := new(request.RegisterRequest)

	errRequest := ctx.BodyParser(request)

	if errRequest != nil {
		return ctx.JSON(fiber.Map{
			"Status":  400,
			"Message": "Bad Request",
		})
	}

	register, errRegister := c.Service.Register(request)

	if errRegister != nil {
		return ctx.JSON(fiber.Map{
			"Status":  500,
			"Message": errRegister.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"Status":  200,
		"Message": "Register Succes",
		"Data":    register,
	})
}

func (c *Auth_Controller) Login(ctx *fiber.Ctx) error {
	request := new(request.LoginRequest)

	errRequest := ctx.BodyParser(request)

	if errRequest != nil {
		return ctx.JSON(fiber.Map{
			"Status":  400,
			"Message": "Bad Request",
		})
	}

	login, errLogin := c.Service.Login(request)

	if errLogin != nil {
		return ctx.JSON(fiber.Map{
			"Status":  500,
			"Message": errLogin.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"Status":  200,
		"Message": "Login Succes",
		"Data":    login,
	})
}
