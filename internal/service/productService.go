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
	Id          int64  `json:"chrt_id" valid:"int, required"`
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
	log   *logrus.Logger
	ch    Cacher
	repos prepos.ProductRepository
}

func NewProductServ(log *logrus.Logger, ch Cacher, repos prepos.ProductRepository) *ProductServ {
	return &ProductServ{
		log:   log,
		ch:    ch,
		repos: repos,
	}
}

func (pr *ProductServ) CreateProduct(prodRpr ProductRepr) (int64, error) {
	result, err := govalidator.ValidateStruct(prodRpr)

	if err != nil {
		pr.log.Errorf("service product err: %s", err.Error())
	}

	if !result {
		pr.log.Error("service product err: invalid data")
		return -1, core.NewInvalidDataErr(http.StatusBadRequest, "product", ExampleProdReq)
	}

	prod := core.Product(prodRpr)
	fmt.Println(prod)
	id, err := pr.repos.Create(prod)

	if err != nil {
		pr.log.Errorf("service product: can not create product: %s", err.Error())
		return -1, err
	}

	prodRpr.Id = id

	pr.ch.SetProduct(id, prodRpr)
	pr.log.Info("service product create")

	return id, nil
}

func (pr *ProductServ) All(trNum string) ([]ProductRepr, error) {
	prodsReps := make([]ProductRepr, 0)

	prods, err := pr.repos.FindByTrackNumber(trNum)

	if err != nil {
		pr.log.Errorf("service product err: %s", err.Error())
		return nil, err
	}

	for _, prod := range prods {
		prodsReps = append(prodsReps, ProductRepr(prod))
	}

	pr.log.Info("service product find all")

	return prodsReps, nil
}

func (pr *ProductServ) Product(id int64) (ProductRepr, error) {
	prodCh, err := pr.ch.Product(id)

	if err == nil {
		pr.log.Infof("service product find in cache product by id: %d", id)
		return prodCh, nil
	}

	prod, err := pr.repos.Product(id)

	pr.log.Infof("service product find by id: %d", id)

	return ProductRepr(prod), err
}

func (pr *ProductServ) DeleteProduct(id int64) (int64, error) {
	delProdId, err := pr.repos.Delete(id)

	if err != nil {
		pr.log.Errorf("service product delete: can not delete product: %s", err.Error())
		return -1, err
	}

	pr.ch.DeleteProduct(id)
	pr.log.Infof("service product delete by id: %d", id)

	return delProdId, err
}
