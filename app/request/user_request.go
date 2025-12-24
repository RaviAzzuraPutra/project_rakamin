package request

import "time"

type User_Request struct {
	Nama         *string   `form:"nama"`
	Password     *string   `form:"password"`
	NoTelp       *string   `form:"notelp"`
	TanggalLahir time.Time `form:"tanggal_lahir"`
	JenisKelamin *string   `form:"jenis_kelamin"`
	Tentang      *string   `form:"tentang"`
	Pekerjaan    *string   `form:"pekerjaan"`
}
