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
	deliveryTable = "delivery"
)

type DeliveryRepos struct {
	log *logrus.Logger
	db  *sqlx.DB
}

func NewDeliveryRepos(log *logrus.Logger, db *sqlx.DB) *DeliveryRepos {
	return &DeliveryRepos{
		log: log,
		db:  db,
	}
}

func (dr *DeliveryRepos) Delivery(id int64) (core.Delivery, error) {
	var delivery core.Delivery
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", deliveryTable)
	err := dr.db.Get(&delivery, query, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			dr.log.Errorf("error get delivery: not found: %s", err.Error())
			return core.Delivery{}, core.NewNotFundErr(http.StatusNotFound, "delivery")
		}

		dr.log.Errorf("error can not create delivery: %s", err.Error())
		return core.Delivery{}, core.NewCantCreateErr(http.StatusBadRequest, "delivery")
	}

	dr.log.Infof("find delivery by id: %d", id)
	return delivery, nil
}

func (dr *DeliveryRepos) Delete(id int) (int, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", deliveryTable)
	_, err := dr.db.Exec(query, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			dr.log.Errorf("error delete delivery: not foud: %s", err.Error())
			return -1, core.NewNotFundErr(http.StatusNotFound, "delivery")
		}

		dr.log.Errorf("error can not delete delivery: %s", err.Error())
		return -1, core.NewCantCreateErr(http.StatusBadRequest, "delivery")
	}

	dr.log.Infof("delete delivery with id: %d", id)
	return id, nil
}
