package repos

import (
	"database/sql"
	"fmt"
	"github.com/GroVlAn/WBTechL0/internal/core"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"net/http"
)

const (
	productTable = "product"
)

type ProductRepos struct {
	db *sqlx.DB
}

func NewProductRepos(db *sqlx.DB) *ProductRepos {
	return &ProductRepos{
		db: db,
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
		return -1, core.NewCantCreateErr(http.StatusBadRequest, "product")
	}

	return id, nil
}

func (pr *ProductRepos) Product(id int) (core.Product, error) {
	var product core.Product
	query := fmt.Sprintf("SELECT * FROM %s WHERE chrt_id=$1", productTable)
	err := pr.db.Get(&product, query, id)

	if errors.Is(err, sql.ErrNoRows) {
		return core.Product{}, core.NewNotFundErr(http.StatusNotFound, "product")
	}

	return product, core.NewCantCreateErr(http.StatusBadRequest, "product")
}

func (pr *ProductRepos) FindByTrackNumber(trNumb string) ([]core.Product, error) {
	var prods []core.Product
	query := fmt.Sprintf("SElECT * FROM %s WHERE track_number=$1 ORDER BY chrt_id", productTable)

	err := pr.db.Select(&prods, query, trNumb)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, core.NewNotFundErr(http.StatusNotFound, "products")
		}

		return nil, core.NewCantCreateErr(http.StatusBadRequest, "products")
	}

	return prods, nil
}

func (pr *ProductRepos) Delete(id int) (int, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE chrt_id=$1 RETURNING chrt_id", productTable)
	_, err := pr.db.Exec(query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return -1, core.NewNotFundErr(http.StatusNotFound, "product")
		}

		return -1, core.NewNotFundErr(http.StatusBadRequest, "product")
	}

	return id, nil
}
