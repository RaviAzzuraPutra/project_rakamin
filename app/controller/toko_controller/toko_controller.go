package toko_controller

import (
	"last-project/app/request"
	"last-project/app/service/interface/toko_service_interface"
	"last-project/app/utils"

	"github.com/gofiber/fiber/v2"
)

type Toko_Controller struct {
	service toko_service_interface.Toko_Service_Interface
}

func NewTokoControllerRegistry(service toko_service_interface.Toko_Service_Interface) *Toko_Controller {
	return &Toko_Controller{
		service: service,
	}
}

func (c *Toko_Controller) GetToko(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(string)

	toko, err := c.service.GetToko(userID)

	if err != nil {
		return ctx.JSON(fiber.Map{
			"Status":  404,
			"Message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"Status":  200,
		"Message": "Success Get Store Data",
		"Data":    toko,
	})
}

func (c *Toko_Controller) UpdateToko(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(string)
	request := new(request.Toko_Request)

	file, err := ctx.FormFile("url_foto")

	errRequest := ctx.BodyParser(request)

	if errRequest != nil {
		return ctx.JSON(fiber.Map{
			"Status":  400,
			"Message": "Bad Request",
		})
	}

	tempPath := utils.GenerateUUID()

	if err := ctx.SaveFile(file, tempPath); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"Status":  500,
			"Message": "Failed to save file",
		})
	}

	request.UrlFoto = &tempPath

	toko, err := c.service.UpdateToko(userID, request)

	if err != nil {
		return ctx.JSON(fiber.Map{
			"Status":  500,
			"Message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"Status":  200,
		"Message": "Berhasil Upload Foto",
		"Data":    toko,
	})
}
