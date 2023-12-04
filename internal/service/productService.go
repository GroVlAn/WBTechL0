package service

import (
	prepos "github.com/GroVlAn/WBTechL0/internal/repository/postgresrepos"
)

type ProductRepr struct {
	TrackNumber string `json:"track_number"`
	Price       int64  `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int64  `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int64  `json:"total_price"`
	NmId        int64  `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int32  `json:"status"`
}

type ProductServ struct {
	repos *prepos.ProductRepository
}

func NewProductServ(repos *prepos.ProductRepository) *ProductServ {
	return &ProductServ{
		repos: repos,
	}
}

func (pr *ProductServ) CreateProduct(prodRpr ProductRepr) (int, error) {
	return -1, nil
}

func (pr *ProductServ) Product(id int) (ProductRepr, error) {
	return ProductRepr{}, nil
}

func (pr *ProductServ) DeleteProduct(id int) (int, error) {
	return -1, nil
}
