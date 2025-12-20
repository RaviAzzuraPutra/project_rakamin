package app_config

import "os"

var PORT string

func App_Config() {
	PORT = os.Getenv("APP_PORT")
}
