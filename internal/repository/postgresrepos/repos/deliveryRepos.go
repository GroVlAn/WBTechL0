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
	deliveryTable = "delivery"
)

type DeliveryRepos struct {
	db *sqlx.DB
}

func NewDeliveryRepos(db *sqlx.DB) *DeliveryRepos {
	return &DeliveryRepos{
		db: db,
	}
}

func (dr *DeliveryRepos) Delivery(id int64) (core.Delivery, error) {
	var delivery core.Delivery
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", deliveryTable)
	err := dr.db.Get(&delivery, query, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return core.Delivery{}, core.NewNotFundErr(http.StatusNotFound, "delivery")
		}
		return core.Delivery{}, core.NewCantCreateErr(http.StatusBadRequest, "delivery")
	}

	return delivery, nil
}

func (dr *DeliveryRepos) Delete(id int) (int, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", deliveryTable)
	_, err := dr.db.Exec(query, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return -1, core.NewNotFundErr(http.StatusNotFound, "delivery")
		}

		return -1, core.NewCantCreateErr(http.StatusBadRequest, "delivery")
	}

	return id, nil
}
