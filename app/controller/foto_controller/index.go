package foto_controller

import (
	"last-project/app/request"
	"last-project/app/service/interface/foto_service_interface"
	"last-project/app/utils"

	"github.com/gofiber/fiber/v2"
)

type Foto_Controller struct {
	service foto_service_interface.Foto_Service_Interface
}

func NewFotoControllerRegistry(service foto_service_interface.Foto_Service_Interface) *Foto_Controller {
	return &Foto_Controller{
		service: service,
	}
}

func (c *Foto_Controller) Create(ctx *fiber.Ctx) error {
	idProduct := ctx.Params("id")
	userID := ctx.Locals("user_id").(string)
	file, errFoto := ctx.FormFile("url")

	if errFoto != nil {
		return ctx.JSON(fiber.Map{
			"Status":  500,
			"Message": errFoto.Error(),
		})
	}

	request := new(request.Foto_Product_Request)

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

	request.URL = &tempPath

	foto, err := c.service.AddPhoto(request, userID, idProduct)

	if err != nil {
		return ctx.JSON(fiber.Map{
			"Status":  500,
			"Message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"Status":  200,
		"Message": "Success Add New Foto For " + idProduct,
		"Data":    foto,
	})
}

func (c *Foto_Controller) GetById(ctx *fiber.Ctx) error {
	idFoto := ctx.Params("id")

	foto, err := c.service.GetById(idFoto)

	if err != nil {
		return ctx.JSON(fiber.Map{
			"Status":  500,
			"Message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"Status":  200,
		"Message": "Success Get Foto",
		"Data":    foto,
	})
}

func (c *Foto_Controller) GetByIdProduct(ctx *fiber.Ctx) error {
	idProduct := ctx.Params("id")

	foto, err := c.service.GetByIdProduct(idProduct)

	if err != nil {
		return ctx.JSON(fiber.Map{
			"Status":  500,
			"Message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"Status":  200,
		"Message": "Success Get Foto",
		"Data":    foto,
	})
}

func (c *Foto_Controller) Delete(ctx *fiber.Ctx) error {
	idFoto := ctx.Params("id")

	errDelete := c.service.Delete(idFoto)

	if errDelete != nil {
		return ctx.JSON(fiber.Map{
			"Status":  500,
			"Message": errDelete.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"Status":  200,
		"Message": "Succes Delete Foto",
	})
}
