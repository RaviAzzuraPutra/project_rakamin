package trx_service

import (
	"errors"
	"last-project/app/models"
	"last-project/app/repository/interface/product_interface_repository"
	"last-project/app/repository/interface/trx_interface_repository"
	"last-project/app/request"
	"last-project/app/utils"
	"strconv"
)

type Trx_Service struct {
	repository trx_interface_repository.Trx_Interface_Repository
	product    product_interface_repository.Product_Interface_Repository
}

func NewTrxServiceRegistry(repository trx_interface_repository.Trx_Interface_Repository,
	product product_interface_repository.Product_Interface_Repository) *Trx_Service {
	return &Trx_Service{
		repository: repository,
		product:    product,
	}
}

func (s *Trx_Service) Create(request *request.Create_Trx_Request, IDUser string) (*models.Trx, error) {

	if request.IDAlamat == nil || *request.IDAlamat == "" {
		return nil, errors.New("Address Cannot Be Empty.")
	}

	if request.MethodBayar == nil || *request.MethodBayar == "" {
		return nil, errors.New("Payment Method Cannot Be Empty.")
	}

	if len(request.Items) == 0 {
		return nil, errors.New("Transaction item is empty")
	}

	trxID := utils.GenerateUUID()
	invoice := "INV-" + trxID

	totalHarga := 0
	var detailTrx []models.DetailTrx

	for _, item := range request.Items {
		product, err := s.product.GetByIdAndIdToko(*item.IDProduct, *item.IDToko)

		if err != nil || product == nil {
			return nil, errors.New("Product not found")
		}

		if product.Stok == nil || *product.Stok < *item.Kuantitas {
			return nil, errors.New("Product stock is not sufficient")
		}

		errStock := s.product.ReduceStock(
			*item.IDProduct,
			*item.IDToko,
			*item.Kuantitas,
		)

		if errStock != nil {
			return nil, errStock
		}

		logID := utils.GenerateUUID()

		logProduk := models.LogProduk{
			ID:            &logID,
			IDProduk:      product.ID,
			NamaProduk:    product.NamaProduk,
			Slug:          product.Slug,
			HargaReseller: product.HargaReseller,
			HargaKonsumen: product.HargaKonsumen,
			Deskripsi:     product.Deskripsi,
			IDToko:        product.IDToko,
			IDCategory:    product.IDCategory,
		}

		errLog := s.repository.CreateLog(&logProduk)
		if errLog != nil {
			return nil, errLog
		}

		harga, _ := strconv.Atoi(*product.HargaKonsumen)
		subtotal := harga * (*item.Kuantitas)
		totalHarga += subtotal

		detailID := utils.GenerateUUID()

		detailTrx = append(detailTrx, models.DetailTrx{
			ID:          &detailID,
			IDTrx:       &trxID,
			IDLogProduk: &logID,
			IDToko:      item.IDToko,
			Kuantitas:   item.Kuantitas,
			HargaTotal:  &subtotal,
		})
	}

	trx := models.Trx{
		ID:               &trxID,
		IDUser:           &IDUser,
		AlamatPengiriman: request.IDAlamat,
		HargaTotal:       totalHarga,
		KodeInvoice:      &invoice,
		MethodBayar:      request.MethodBayar,
	}

	errTrx := s.repository.CreateTrx(&trx)

	if errTrx != nil {
		return nil, errTrx
	}

	errDetail := s.repository.CreateDetailTrx(&detailTrx)

	if errDetail != nil {
		return nil, errDetail
	}

	return &trx, nil
}

func (s *Trx_Service) GetVeryDetailedTrx(IDTrx string) (*models.Trx, []models.DetailTrx, error) {

	trx, errGetTrx := s.repository.GetByIdTrx(IDTrx)

	if trx == nil {
		return nil, nil, errors.New("Transaction Not Found")
	}

	if errGetTrx != nil {
		return nil, nil, errors.New("An error occurred while retrieving transaction data. " + errGetTrx.Error())
	}

	detail, errGetDetail := s.repository.GetDetailByTrxId(IDTrx)

	if errGetDetail != nil {
		return nil, nil, errors.New("An error occurred while retrieving transaction detail data. " + errGetDetail.Error())
	}

	if len(detail) == 0 {
		return trx, []models.DetailTrx{}, nil
	}

	return trx, detail, nil
}
