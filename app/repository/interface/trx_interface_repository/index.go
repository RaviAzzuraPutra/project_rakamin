package trx_interface_repository

import "last-project/app/models"

type Trx_Interface_Repository interface {
	CreateTrx(trx *models.Trx) error
	CreateLog(log *models.LogProduk) error
	CreateDetailTrx(detail *[]models.DetailTrx) error
	GetByIdTrx(IDTrx string) (*models.Trx, error)
	GetDetailByTrxId(IDTrx string) ([]models.DetailTrx, error)
}
