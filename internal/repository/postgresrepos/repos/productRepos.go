package repos

import (
	"database/sql"
	"fmt"
	"github.com/GroVlAn/WBTechL0/internal/core"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	productTable = "product"
)

type ProductRepos struct {
	log *logrus.Logger
	db  *sqlx.DB
}

func NewProductRepos(log *logrus.Logger, db *sqlx.DB) *ProductRepos {
	return &ProductRepos{
		log: log,
		db:  db,
	}
}

func (pr *ProductRepos) Create(prod core.Product) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s ("+
		"track_number,"+
		"price,"+
		"rid,"+
		"name,"+
		"sale,"+
		"size,"+
		"total_price,"+
		"nm_id,"+
		"brand,"+
		"status"+
		") VALUES "+
		"($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING chrt_id", productTable)
	row := pr.db.QueryRow(
		query,
		prod.TrackNumber,
		prod.Price,
		prod.Rid,
		prod.Name,
		prod.Sale,
		prod.Size,
		prod.TotalPrice,
		prod.NmId,
		prod.Brand,
		prod.Status,
	)

	if err := row.Scan(&id); err != nil {
		pr.log.Errorf("error can not create product: %s", err.Error())
		return -1, core.NewCantCreateErr(http.StatusBadRequest, "product")
	}

	pr.log.Infof("create product with id: %d", id)
	return id, nil
}

func (pr *ProductRepos) Product(id int) (core.Product, error) {
	var product core.Product
	query := fmt.Sprintf("SELECT * FROM %s WHERE chrt_id=$1", productTable)
	err := pr.db.Get(&product, query, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			pr.log.Errorf("error get product: not found: %s", err.Error())

			return core.Product{}, core.NewNotFundErr(http.StatusNotFound, "product")
		}

		pr.log.Errorf("error get product: can not find product: %s", err.Error())
		return core.Product{}, core.NewCantCreateErr(http.StatusBadRequest, "product")
	}

	pr.log.Infof("find product by id: %d", id)
	return product, nil
}

func (pr *ProductRepos) FindByTrackNumber(trNumb string) ([]core.Product, error) {
	var prods []core.Product
	query := fmt.Sprintf("SElECT * FROM %s WHERE track_number=$1 ORDER BY chrt_id", productTable)

	err := pr.db.Select(&prods, query, trNumb)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			pr.log.Errorf("error find all products: not found: %s", err.Error())

			return nil, core.NewNotFundErr(http.StatusNotFound, "products")
		}

		pr.log.Errorf("error find all product: can not find product: %s", err.Error())
		return nil, core.NewCantCreateErr(http.StatusBadRequest, "products")
	}

	pr.log.Infof("fidn product by track number: %s", trNumb)
	return prods, nil
}

func (pr *ProductRepos) Delete(id int) (int, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE chrt_id=$1 RETURNING chrt_id", productTable)
	_, err := pr.db.Exec(query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			pr.log.Errorf("error delete product: not found: %s", err.Error())
			return -1, core.NewNotFundErr(http.StatusNotFound, "product")
		}

		pr.log.Errorf("error delete product: can not find product: %s", err.Error())
		return -1, core.NewNotFundErr(http.StatusBadRequest, "product")
	}

	pr.log.Infof("delete product by id: %d", id)
	return id, nil
}
