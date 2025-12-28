package request

type Product_Request struct {
	NamaProduk    *string `form:"nama_produk"`
	HargaReseller *string `form:"harga_reseller"`
	HargaKonsumen *string `form:"harga_konsumen"`
	Stok          *int    `form:"stok"`
	Deskripsi     *string `form:"deskripsi"`
	IDCategory    *string `form:"IDCategory"`
}

type Foto_Product_Request struct {
	URL *string `form:"url"`
}
