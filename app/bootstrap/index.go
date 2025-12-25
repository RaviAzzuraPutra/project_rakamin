package bootstrap

import (
	"last-project/app/bootstrap/registry"
	"last-project/app/config"
	"last-project/app/config/app_config"
	"last-project/app/database"
	"last-project/app/router/alamat_router"
	"last-project/app/router/auth_router"
	"last-project/app/router/category_router"
	"last-project/app/router/toko_router"
	"last-project/app/router/user_router"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Bootstrap() {
	errEnv := godotenv.Load()

	if errEnv != nil {
		panic("Error loading environment : " + errEnv.Error())
	}

	config.Config()

	database.ConnectToDatabase()

	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"Message": "The application is running well.",
		})
	})

	AuthModules := registry.AuthRegistry()
	TokoModules := registry.Toko_Registry()
	AlamatModules := registry.Alamat_Registry()
	UserModules := registry.User_Registry()
	CategoryModules := registry.Category_Registry()

	auth_router.AuthRouter(app, AuthModules.AuthController)
	toko_router.TokoRouter(app, TokoModules.TokoController)
	alamat_router.AlamatRouter(app, AlamatModules.AlamatController)
	user_router.UserRouter(app, UserModules.UserController)
	category_router.CategoryRouter(app, CategoryModules.CategoryController)

	app.Listen(app_config.PORT)
}
