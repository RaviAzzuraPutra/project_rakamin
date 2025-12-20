package models

import "time"

type Alamat struct {
	ID           *string `json:"id" gorm:"type:char(36);primaryKey"`
	IDUser       *string `json:"id_user" gorm:"type:char(36)"`
	JudulAlamat  *string `json:"judul_alamat"`
	NamaPenerima *string `json:"nama_penerima"`
	NoTelp       *string `json:"no_telp"`
	DetailAlamat *string `json:"detail_alamat"`
	CreatedAt    time.Time
	UpdatedAt    time.Time

	User User `gorm:"foreignKey:IDUser;references:ID"`
}
