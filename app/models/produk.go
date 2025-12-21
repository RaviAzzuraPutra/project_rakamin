package models

import "time"

type Produk struct {
	ID            *string `json:"id" gorm:"type:char(36);primaryKey"`
	NamaProduk    *string `json:"nama_produk"`
	Slug          *string `json:"slug"`
	HargaReseller *string `json:"harga_reseller"`
	HargaKonsumen *string `json:"harga_konsumen"`
	Stok          *int    `json:"stok"`
	Deskripsi     *string `json:"deskripsi" gorm:"type:text"`
	IDToko        *string `json:"id_toko" gorm:"type:char(36)"`
	IDCategory    *string `json:"id_category" gorm:"type:char(36)"`
	CreatedAt     time.Time
	UpdatedAt     time.Time

	Toko     Toko         `gorm:"foreignKey:IDToko;references:ID"`
	Category Category     `gorm:"foreignKey:IDCategory;references:ID"`
	Foto     []FotoProduk `gorm:"foreignKey:IDProduk;references:ID"`
}
