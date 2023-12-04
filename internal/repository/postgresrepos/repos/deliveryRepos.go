package repos

import (
	"fmt"
	"github.com/GroVlAn/WBTechL0/internal/core"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
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

func (dr *DeliveryRepos) Create(d core.Delivery) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, phone, zip, city, address, region, email)"+
		"VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", deliveryTable)
	row := dr.db.QueryRow(
		query,
		d.Name,
		d.Phone,
		d.Zip,
		d.City,
		d.Address,
		d.Region,
		d.Email,
	)

	if err := row.Scan(&id); err != nil {
		return -1, errors.Wrap(err, "error can not create delivery")
	}

	return id, nil
}

func (dr *DeliveryRepos) Delivery(id int) (core.Delivery, error) {
	var delivery core.Delivery
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", deliveryTable)
	err := dr.db.Get(&delivery, query, id)

	return delivery, errors.Wrap(err, "delivery not found")
}

func (dr *DeliveryRepos) Delete(id int) (int, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", deliveryTable)
	_, err := dr.db.Exec(query, id)

	if err != nil {
		return -1, errors.Wrap(err, "delivery not found")
	}

	return id, nil
}
