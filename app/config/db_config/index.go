package db_config

import "os"

type Database_Variable struct {
	MYSQL_ROOT_PASSWORD string
	MYSQL_DATABASE      string
	MYSQL_USER          string
	MYSQL_HOST          string
	MYSQL_PORT          string
}

func DB_Config() *Database_Variable {
	return &Database_Variable{
		MYSQL_ROOT_PASSWORD: os.Getenv("MYSQL_ROOT_PASSWORD"),
		MYSQL_DATABASE:      os.Getenv("MYSQL_DATABASE"),
		MYSQL_USER:          os.Getenv("MYSQL_USER"),
		MYSQL_HOST:          os.Getenv("MYSQL_HOST"),
		MYSQL_PORT:          os.Getenv("MYSQL_PORT"),
	}
}
