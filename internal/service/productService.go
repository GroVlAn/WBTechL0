package service

import (
	"fmt"
	"github.com/GroVlAn/WBTechL0/internal/core"
	prepos "github.com/GroVlAn/WBTechL0/internal/repository/postgresrepos"
	"github.com/asaskevich/govalidator"
	"github.com/sirupsen/logrus"
	"net/http"
)

var ExampleProdReq = ProductRepr{
	TrackNumber: "WBILMTESTTRACK",
	Price:       453,
	Rid:         "ab4219087a764ae0btest",
	Name:        "Mascaras",
	Sale:        30,
	Size:        "0",
	TotalPrice:  317,
	NmId:        2389212,
	Brand:       "Vivienne Sabo",
	Status:      202,
}

type ProductRepr struct {
	Id          int    `json:"chrt_id" valid:"int, required"`
	TrackNumber string `json:"track_number" valid:"type(string), required"`
	Price       int64  `json:"price" valid:"int, required"`
	Rid         string `json:"rid" valid:"type(string), required"`
	Name        string `json:"name" valid:"type(string), required"`
	Sale        int64  `json:"sale" valid:"int, required"`
	Size        string `json:"size" valid:"type(string), required"`
	TotalPrice  int64  `json:"total_price" valid:"int, required"`
	NmId        int64  `json:"nm_id" valid:"int, required"`
	Brand       string `json:"brand" valid:"type(string), required"`
	Status      int32  `json:"status" valid:"int, required"`
}

type ProductServ struct {
	repos prepos.ProductRepository
}

func NewProductServ(repos prepos.ProductRepository) *ProductServ {
	return &ProductServ{
		repos: repos,
	}
}

func (pr *ProductServ) CreateProduct(prodRpr ProductRepr) (int, error) {
	result, err := govalidator.ValidateStruct(prodRpr)

	if err != nil {
		logrus.Errorln(err.Error())
	}

	if !result {
		return -1, core.NewInvalidDataErr(http.StatusBadRequest, "product", ExampleProdReq)
	}

	prod := core.Product(prodRpr)
	fmt.Println(prod)
	id, err := pr.repos.Create(prod)

	return id, err
}

func (pr *ProductServ) All(trNum string) ([]ProductRepr, error) {
	prodsReps := make([]ProductRepr, 0)

	prods, err := pr.repos.FindByTrackNumber(trNum)

	for _, prod := range prods {
		prodsReps = append(prodsReps, ProductRepr(prod))
	}

	return prodsReps, err
}

func (pr *ProductServ) Product(id int) (ProductRepr, error) {
	prod, err := pr.repos.Product(id)

	return ProductRepr(prod), err
}

func (pr *ProductServ) DeleteProduct(id int) (int, error) {
	delProdId, err := pr.repos.Delete(id)

	return delProdId, err
}
