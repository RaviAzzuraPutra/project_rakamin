package config

import (
	"last-project/app/config/app_config"
	"last-project/app/config/cloudinary_config"
	"last-project/app/config/db_config"
)

func Config() {
	app_config.App_Config()
	db_config.DB_Config()
	cloudinary_config.CloudinaryConfig()
}
