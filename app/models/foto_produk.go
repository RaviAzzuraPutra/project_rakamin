package models

import "time"

type FotoProduk struct {
	ID        *string `json:"id" gorm:"type:char(36);primaryKey"`
	IDProduk  *string `json:"id_produk" gorm:"type:char(36)"`
	URL       *string `json:"url"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Produk Produk `gorm:"foreignKey:IDProduk;references:ID;constraint:OnDelete:CASCADE"`
}
