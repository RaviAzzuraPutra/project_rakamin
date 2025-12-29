package request

type Create_Trx_Request struct {
	IDAlamat    *string           `json:"id_alamat"`
	MethodBayar *string           `json:"method_bayar"`
	Items       []Create_Trx_Item `json:"items"`
}

type Create_Trx_Item struct {
	IDProduct *string `json:"id_product"`
	IDToko    *string `json:"id_toko"`
	Kuantitas *int    `json:"kuantitas"`
}
