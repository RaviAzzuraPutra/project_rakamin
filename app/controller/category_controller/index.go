package category_controller

import (
	"last-project/app/request"
	"last-project/app/service/interface/category_service_interface"

	"github.com/gofiber/fiber/v2"
)

type Category_Controller struct {
	service category_service_interface.Category_Service_Interface
}

func NewCategoryControllerRegistry(service category_service_interface.Category_Service_Interface) *Category_Controller {
	return &Category_Controller{
		service: service,
	}
}

func (c *Category_Controller) Create(ctx *fiber.Ctx) error {
	request := new(request.Category_Request)

	errRequest := ctx.BodyParser(request)

	if errRequest != nil {
		return ctx.JSON(fiber.Map{
			"Status":  400,
			"Message": "Bad Request",
		})
	}

	category, err := c.service.Create(request)

	if err != nil {
		return ctx.JSON(fiber.Map{
			"Status":  500,
			"Message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"Status":  200,
		"Message": "Success Create Category",
		"Data":    category,
	})
}

func (c *Category_Controller) Get(ctx *fiber.Ctx) error {
	category, err := c.service.Get()

	if err != nil {
		return ctx.JSON(fiber.Map{
			"Status":  500,
			"Message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"Status":  200,
		"Message": "Success Get Category",
		"Data":    category,
	})
}

func (c *Category_Controller) GetById(ctx *fiber.Ctx) error {
	idCategory := ctx.Params("id")

	category, err := c.service.GetById(idCategory)

	if err != nil {
		return ctx.JSON(fiber.Map{
			"Status":  500,
			"Message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"Status":  200,
		"Message": "Success Get Category",
		"Data":    category,
	})
}

func (c *Category_Controller) Update(ctx *fiber.Ctx) error {
	idCategory := ctx.Params("id")

	request := new(request.Category_Request)

	errRequest := ctx.BodyParser(request)

	if errRequest != nil {
		return ctx.JSON(fiber.Map{
			"Status":  400,
			"Message": "Bad Request",
		})
	}

	category, err := c.service.Update(idCategory, request)

	if err != nil {
		return ctx.JSON(fiber.Map{
			"Status":  500,
			"Message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"Status":  200,
		"Message": "Success Update Category",
		"Data":    category,
	})
}

func (c *Category_Controller) Delete(ctx *fiber.Ctx) error {
	idCategory := ctx.Params("id")

	errDelete := c.service.Delete(idCategory)

	if errDelete != nil {
		return ctx.JSON(fiber.Map{
			"Status":  500,
			"Message": errDelete.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"Status":  200,
		"Message": "Success Delete Category",
	})
}
