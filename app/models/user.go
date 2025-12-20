package models

import "time"

type User struct {
	ID           *string   `json:"id" gorm:"type:char(36);primaryKey"`
	Nama         *string   `json:"nama"`
	KataSandi    *string   `json:"katasandi"`
	NoTelp       *string   `json:"notelp" gorm:"unique"`
	TanggalLahir time.Time `json:"tanggallahir"`
	JenisKelamin *string   `json:"jeniskelamin"`
	Tentang      *string   `json:"tentang" gorm:"type:text"`
	Pekerjaan    *string   `json:"pekerjaan"`
	Email        *string   `json:"email" gorm:"unique"`
	IDProvinsi   *string   `json:"idprovinsi"`
	IDKota       *string   `json:"idkota"`
	IsAdmin      bool      `json:"isadmin"`
	CreatedAt    time.Time
	UpdatedAt    time.Time

	Alamat []Alamat `gorm:"foreignKey:IDUser"`
	Toko   *Toko    `gorm:"foreignKey:IDUser"`
}
