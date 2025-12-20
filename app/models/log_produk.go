package models

import "time"

type LogProduk struct {
	ID            *string `json:"id" gorm:"type:char(36);primaryKey"`
	IDProduk      *string `json:"id_produk" gorm:"type:char(36)"`
	NamaProduk    *string `json:"nama_produk"`
	Slug          *string `json:"slug"`
	HargaReseller *string `json:"harga_reseller"`
	HargaKonsumen *string `json:"harga_konsumen"`
	Deskripsi     *string `json:"deskripsi" gorm:"type:text"`
	IDToko        *string `json:"id_toko" gorm:"type:char(36)"`
	IDCategory    *string `json:"id_category" gorm:"type:char(36)"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
