package trx_controller

import (
	"last-project/app/request"
	"last-project/app/service/interface/trx_service_interface"

	"github.com/gofiber/fiber/v2"
)

type Trx_Controller struct {
	service trx_service_interface.Trx_Service_Interface
}

func NewTrxControllerRegistry(service trx_service_interface.Trx_Service_Interface) *Trx_Controller {
	return &Trx_Controller{
		service: service,
	}
}

func (c *Trx_Controller) Create(ctx *fiber.Ctx) error {
	request := new(request.Create_Trx_Request)
	userID := ctx.Locals("user_id").(string)

	if err := ctx.BodyParser(request); err != nil {
		return ctx.JSON(fiber.Map{
			"Status":  400,
			"Message": "Bad Request",
		})
	}

	trx, err := c.service.Create(request, userID)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"Status":  500,
			"Message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"Status":  200,
		"Message": "Transaction Success",
		"Data":    trx,
	})
}

func (c *Trx_Controller) GetVeryDetailedTrx(ctx *fiber.Ctx) error {

	idAlamat := ctx.Params("id")

	trx, detailTrx, errTrx := c.service.GetVeryDetailedTrx(idAlamat)

	if errTrx != nil {
		return ctx.JSON(fiber.Map{
			"Status":  500,
			"Message": errTrx.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"Status":             200,
		"Message":            "Transaction Success",
		"Transaction":        trx,
		"Detail Transaction": detailTrx,
	})
}
