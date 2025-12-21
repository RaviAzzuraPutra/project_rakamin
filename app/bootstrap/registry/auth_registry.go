package registry

import (
	"last-project/app/controller/auth_controller"
	"last-project/app/repository/auth_repository"
	"last-project/app/service/auth_service"
)

type Auth_Modules struct {
	AuthController *auth_controller.Auth_Controller
}

func AuthRegistry() *Auth_Modules {
	AuthRepository := auth_repository.NewAuthRpositoryRegistry()

	AuthService := auth_service.NewAuthServiceRegistry(AuthRepository)

	AuthController := auth_controller.NewAuthControllerRegistry(AuthService)

	return &Auth_Modules{
		AuthController: AuthController,
	}
}
