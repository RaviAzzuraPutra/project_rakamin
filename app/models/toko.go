package models

import "time"

type Toko struct {
	ID        *string `json:"id" gorm:"type:char(36);primaryKey"`
	IDUser    *string `json:"id_user" gorm:"type:char(36)"`
	NamaToko  *string `json:"nama_toko"`
	URLFoto   *string `json:"url_foto"`
	CreatedAt time.Time
	UpdatedAt time.Time

	User   *User    `gorm:"foreignKey:IDUser;references:ID"`
	Produk []Produk `gorm:"foreignKey:IDToko;references:ID"`
}
