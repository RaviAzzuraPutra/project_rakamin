package models

import "time"

type DetailTrx struct {
	ID          *string `json:"id" gorm:"type:char(36);primaryKey"`
	IDTrx       *string `json:"id_trx" gorm:"type:char(36)"`
	IDLogProduk *string `json:"id_log_produk" gorm:"type:char(36)"`
	IDToko      *string `json:"id_toko" gorm:"type:char(36)"`
	Kuantitas   int     `json:"kuantitas"`
	HargaTotal  int     `json:"harga_total"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
