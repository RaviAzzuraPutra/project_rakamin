package models

import "time"

type Trx struct {
	ID               *string `json:"id" gorm:"type:char(36);primaryKey"`
	IDUser           *string `json:"id_user" gorm:"type:char(36)"`
	AlamatPengiriman *string `json:"alamat_pengiriman" gorm:"type:char(36)"`
	HargaTotal       int     `json:"harga_total"`
	KodeInvoice      *string `json:"kode_invoice"`
	MethodBayar      *string `json:"method_bayar"`
	CreatedAt        time.Time
	UpdatedAt        time.Time

	User   User   `gorm:"foreignKey:IDUser;references:ID"`
	Alamat Alamat `gorm:"foreignKey:AlamatPengiriman;references:ID"`
}
