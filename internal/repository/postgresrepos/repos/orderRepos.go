package repos

import (
	"fmt"
	"github.com/GroVlAn/WBTechL0/internal/core"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

const (
	orderTable        = "order"
	orderProductTable = "order_product"
)

type OrderRepos struct {
	db *sqlx.DB
}

func NewOrderRepos(db *sqlx.DB) *OrderRepos {
	return &OrderRepos{
		db: db,
	}
}

func (or *OrderRepos) Create(ord core.Order, dId int, pmtId int, prodId int) (int, error) {
	tx, err := or.db.Begin()

	if err != nil {
		return -1, err
	}

	var id int
	createOrderQuery := fmt.Sprintf("INSERT INTO %s ("+
		"order_uid,"+
		"track_number,"+
		"entry,"+
		"locale,"+
		"customer_id,"+
		"delivery_service,"+
		"shardkey,"+
		"sm_id,"+
		"off_shard,"+
		"date_created,"+
		"delivery_id,"+
		"payment_id"+
		") VALUES "+
		"($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id", orderTable)
	row := or.db.QueryRow(
		createOrderQuery,
		ord.OrderUid,
		ord.TrackNumber,
		ord.Entry,
		ord.Locale,
		ord.InternalSignature,
		ord.CustomerId,
		ord.DeliveryService,
		ord.Shardkey,
		ord.SmId,
		ord.OffShard,
		ord.DateCreated,
		ord.DeliveryId,
		ord.PaymentId,
	)

	if err := row.Scan(&id); err != nil {
		if rolErr := tx.Rollback(); rolErr != nil {
			return -1, rolErr
		}

		return -1, errors.Wrap(err, "error can not create order")
	}

	orderProductQuery := fmt.Sprintf("INSERT INTO %s (order_id, product_id) VALUES ($1, $2)", orderProductTable)
	_, orPrErr := tx.Exec(orderProductQuery, id, prodId)

	if orPrErr != nil {
		if rolErr := tx.Rollback(); rolErr != nil {
			return -1, rolErr
		}
		return -1, errors.Wrap(err, "error can not create order_product")
	}

	return id, tx.Commit()
}

func (or *OrderRepos) Order(id int) (core.Order, error) {
	var order core.Order
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", orderTable)
	err := or.db.Get(&order, query, id)

	return order, errors.Wrap(err, "order not found")
}

func (or *OrderRepos) Delete(id int) (int, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", orderTable)
	_, err := or.db.Exec(query, id)

	if err != nil {
		return -1, errors.Wrap(err, "order not found")
	}

	return id, nil
}
