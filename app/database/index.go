package database

import (
	"fmt"
	"last-project/app/config/db_config"
	"last-project/app/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var ErrorConnect error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_config.DB_Config().MYSQL_USER, db_config.DB_Config().MYSQL_ROOT_PASSWORD, db_config.DB_Config().MYSQL_HOST, db_config.DB_Config().MYSQL_PORT, db_config.DB_Config().MYSQL_DATABASE)

	DB, ErrorConnect = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if ErrorConnect != nil {
		panic("An error occurred while attempting to connect to the database. " + ErrorConnect.Error())
	}

	errMigrate := DB.AutoMigrate(&models.User{}, &models.Alamat{}, &models.Category{}, &models.DetailTrx{}, &models.FotoProduk{}, &models.LogProduk{}, &models.Produk{}, &models.Toko{}, &models.Trx{})

	if errMigrate != nil {
		panic("A mistake happened while moving the database. " + errMigrate.Error())
	}

	if DB == nil {
		panic("An error has occurred in the database, mate.")
	}

	log.Printf("Successfully connected to the database! ✅")
}
