package app_config

import "os"

var PORT string
var JWT string

func App_Config() {
	PORT = os.Getenv("APP_PORT")
	JWT = os.Getenv("JWT_SECRET")
}
