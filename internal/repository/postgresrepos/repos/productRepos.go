package repos

import (
	"fmt"
	"github.com/GroVlAn/WBTechL0/internal/core"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
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
		return -1, errors.Wrap(err, "error can not create product")
	}

	return id, nil
}

func (pr *ProductRepos) Product(id int) (core.Product, error) {
	var product core.Product
	query := fmt.Sprintf("SELECT * FROM %s WHERE chrt_id=$1", productTable)
	err := pr.db.Get(&product, query, id)

	return product, errors.Wrap(err, "product not found")
}

func (pr *ProductRepos) Delete(id int) (int, error) {
	var delProdId int
	query := fmt.Sprintf("DELETE FROM %s WHERE chrt_id=$1 RETURNING chrt_id", productTable)
	row := pr.db.QueryRow(query, id)

	if err := row.Scan(&delProdId); err != nil {
		return -1, errors.Wrap(err, "product not found")
	}

	return delProdId, nil
}
