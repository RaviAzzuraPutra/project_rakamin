package request

type Alamat struct {
	JudulAlamat  *string `form:"judul_alamat"`
	NamaPenerima *string `form:"nama_penerima"`
	NoTelp       *string `form:"notelp"`
	DetailAlamat *string `form:"detail_alamat"`
}
