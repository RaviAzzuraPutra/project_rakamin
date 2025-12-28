package product_controller

import (
	"last-project/app/request"
	"last-project/app/service/interface/product_service_interface"

	"github.com/gofiber/fiber/v2"
)

type Product_Controller struct {
	service product_service_interface.Product_Service_Interface
}

func NewProductControllerRegistry(service product_service_interface.Product_Service_Interface) *Product_Controller {
	return &Product_Controller{
		service: service,
	}
}

func (c *Product_Controller) Create(ctx *fiber.Ctx) error {
	request := new(request.Product_Request)

	errRequest := ctx.BodyParser(request)

	if errRequest != nil {
		return ctx.JSON(fiber.Map{
			"Status":  400,
			"Message": "Bad Request",
		})
	}

	userID := ctx.Locals("user_id").(string)

	product, err := c.service.Create(request, userID)

	if err != nil {
		return ctx.JSON(fiber.Map{
			"Status":  500,
			"Message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"Status":  200,
		"Message": "Success Create New Product",
		"Data":    product,
	})

}

func (c *Product_Controller) GetAll(ctx *fiber.Ctx) error {

	product, err := c.service.GetAll()

	if err != nil {
		return ctx.JSON(fiber.Map{
			"Status":  500,
			"Message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"Status":  200,
		"Message": "Success Get Product",
		"Data":    product,
	})
}

func (c *Product_Controller) GetByIdAndIdToko(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(string)
	idProduct := ctx.Params("id")

	product, err := c.service.GetByIdAndIdToko(idProduct, userID)

	if err != nil {
		return ctx.JSON(fiber.Map{
			"Status":  500,
			"Message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"Status":  200,
		"Message": "Success Get Product",
		"Data":    product,
	})
}

func (c *Product_Controller) GetByIdToko(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(string)

	product, err := c.service.GetByIdToko(userID)

	if err != nil {
		return ctx.JSON(fiber.Map{
			"Status":  500,
			"Message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"Status":  200,
		"Message": "Success Get Product",
		"Data":    product,
	})
}

func (c *Product_Controller) GetByCategory(ctx *fiber.Ctx) error {
	search := ctx.Query("search")

	if search == "" {
		return ctx.JSON(fiber.Map{
			"Status":  400,
			"Message": "Search category cannot be empty",
		})
	}

	product, err := c.service.GetByCategory(search)

	if err != nil {
		return ctx.JSON(fiber.Map{
			"Status":  404,
			"Message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"Status":  200,
		"Message": "Success Get Product By Category",
		"Data":    product,
	})
}

func (c *Product_Controller) GetById(ctx *fiber.Ctx) error {
	idProduct := ctx.Params("id")

	product, err := c.service.GetById(idProduct)

	if err != nil {
		return ctx.JSON(fiber.Map{
			"Status":  500,
			"Message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"Status":  200,
		"Message": "Success Get Product",
		"Data":    product,
	})
}

func (c *Product_Controller) Update(ctx *fiber.Ctx) error {
	request := new(request.Product_Request)

	errRequest := ctx.BodyParser(request)

	if errRequest != nil {
		return ctx.JSON(fiber.Map{
			"Status":  400,
			"Message": "Bad Request",
		})
	}

	userID := ctx.Locals("user_id").(string)
	idProduct := ctx.Params("id")

	product, err := c.service.Update(idProduct, userID, request)

	if err != nil {
		return ctx.JSON(fiber.Map{
			"Status":  500,
			"Message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"Status":  200,
		"Message": "Success Update Product",
		"Data":    product,
	})
}

func (c *Product_Controller) Delete(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(string)
	idProduct := ctx.Params("id")

	err := c.service.Delete(idProduct, userID)

	if err != nil {
		return ctx.JSON(fiber.Map{
			"Status":  500,
			"Message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"Status":  200,
		"Message": "Success Delete Product",
	})
}
