package alamat_controller

import (
	"last-project/app/request"
	"last-project/app/service/interface/alamat_service_interface"

	"github.com/gofiber/fiber/v2"
)

type Alamat_Controller struct {
	service alamat_service_interface.Alamat_Service_Interface
}

func NewAlamatControllerRegistry(service alamat_service_interface.Alamat_Service_Interface) *Alamat_Controller {
	return &Alamat_Controller{
		service: service,
	}
}

func (c *Alamat_Controller) Create(ctx *fiber.Ctx) error {
	request := new(request.Alamat)

	userID := ctx.Locals("user_id").(string)

	errRequest := ctx.BodyParser(request)

	if errRequest != nil {
		return ctx.JSON(fiber.Map{
			"Status":  400,
			"Message": "Bad Request",
		})
	}

	alamat, err := c.service.Create(request, userID)

	if err != nil {
		return ctx.JSON(fiber.Map{
			"Status":  500,
			"Message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"Status":  200,
		"Message": "Success Add New Addres",
		"Data":    alamat,
	})
}

func (c *Alamat_Controller) GetByIDUserAndIDAlamat(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(string)
	idAlamat := ctx.Params("id")

	alamat, err := c.service.GetByIDUserAndIDAlamat(userID, idAlamat)

	if err != nil {
		return ctx.JSON(fiber.Map{
			"Status":  500,
			"Message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"Status":  200,
		"Message": "Success Get Data",
		"Data":    alamat,
	})
}

func (c *Alamat_Controller) GetByIDUser(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(string)

	alamat, err := c.service.GetByIDUser(userID)

	if err != nil {
		return ctx.JSON(fiber.Map{
			"Status":  500,
			"Message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"Status":  200,
		"Message": "Success Get Data",
		"Data":    alamat,
	})
}

func (c *Alamat_Controller) Update(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(string)
	idAlamat := ctx.Params("id")
	request := new(request.Alamat)

	errRequest := ctx.BodyParser(request)

	if errRequest != nil {
		return ctx.JSON(fiber.Map{
			"Status":  400,
			"Message": "Bad Request",
		})
	}

	NewAlamat, err := c.service.Update(userID, idAlamat, request)

	if err != nil {
		return ctx.JSON(fiber.Map{
			"Status":  500,
			"Message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"Status":  200,
		"Message": "Success Update Data",
		"Data":    NewAlamat,
	})
}

func (c *Alamat_Controller) Delete(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(string)
	idAlamat := ctx.Params("id")

	errDelete := c.service.Delete(userID, idAlamat)

	if errDelete != nil {
		return ctx.JSON(fiber.Map{
			"Status":  500,
			"Message": errDelete.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"Status":  200,
		"Message": "Success Delete Data",
	})
}
