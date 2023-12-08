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
	orderTable = "order"
)

type OrderRepos struct {
	db *sqlx.DB
}

func NewOrderRepos(db *sqlx.DB) *OrderRepos {
	return &OrderRepos{
		db: db,
	}
}

func (or *OrderRepos) Create(ord core.Order, d core.Delivery, pmt core.Payment) (string, error) {
	tx, errTx := or.db.Begin()

	if errTx != nil {
		return "", errTx
	}

	var dId int64
	queryD := fmt.Sprintf("INSERT INTO %s (name, phone, zip, city, address, region, email)"+
		"VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", deliveryTable)
	rowD := tx.QueryRow(
		queryD,
		d.Name,
		d.Phone,
		d.Zip,
		d.City,
		d.Address,
		d.Region,
		d.Email,
	)

	if errD := rowD.Scan(&dId); errD != nil {
		errTx := tx.Rollback()
		if errTx != nil {
			logrus.Errorf("Cann not roalback delivery: %s", errTx.Error())
			return "", errD
		}

		return "", core.NewCantCreateErr(http.StatusBadRequest, "order - delivery")
	}

	var tran string
	queryPmt := fmt.Sprintf("INSERT INTO %s ("+
		"transaction,"+
		"request_id,"+
		"currency,"+
		"provider,"+
		"amount,"+
		"payment_dt,"+
		"bank,"+
		"delivery_cost,"+
		"goods_total,"+
		"custom_fee) VALUES "+
		"($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING transaction", paymentTable)
	rowPmt := tx.QueryRow(
		queryPmt,
		pmt.Transaction,
		pmt.RequestId,
		pmt.Currency,
		pmt.Provider,
		pmt.Amount,
		pmt.PaymentDt,
		pmt.Bank,
		pmt.DeliveryCost,
		pmt.GoodsTotal,
		pmt.CustomFee,
	)

	if errPmt := rowPmt.Scan(&tran); errPmt != nil {
		errTx := tx.Rollback()
		if errTx != nil {
			logrus.Errorf("Cann not roalback delivery: %s", errTx.Error())
			return "", errTx
		}
		return "", core.NewCantCreateErr(http.StatusBadRequest, "order - payment")
	}

	ord.DeliveryId = dId
	ord.PaymentTransaction = tran

	var orderUid string
	createOrderQuery := fmt.Sprintf("INSERT INTO \"%s\" ("+
		"order_uid, track_number, entry, locale, internal_signature, customer_id,"+
		"delivery_service, shardkey, sm_id, oof_shard, delivery_id, payment_transaction"+
		") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING order_uid", orderTable)
	rowOrd := tx.QueryRow(
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
		ord.OofShard,
		ord.DeliveryId,
		ord.PaymentTransaction,
	)

	if errOrd := rowOrd.Scan(&orderUid); errOrd != nil {
		errTx := tx.Rollback()
		if errTx != nil {
			logrus.Errorf("Cann not roalback delivery: %s", errTx.Error())
			return "", errTx
		}
		return "", core.NewCantCreateErr(http.StatusBadRequest, "order")
	}

	return orderUid, tx.Commit()
}

func (or *OrderRepos) Order(orderUid string) (core.Order, error) {
	var order core.Order
	query := fmt.Sprintf("SELECT * FROM \"%s\" WHERE order_uid=$1", orderTable)
	fmt.Println(orderUid)
	err := or.db.Get(&order, query, orderUid)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return core.Order{}, core.NewNotFundErr(http.StatusNotFound, "order")
		}

		return core.Order{}, core.NewCantCreateErr(http.StatusBadRequest, "order")
	}

	return order, nil
}

func (or *OrderRepos) Delete(orderUid string) (string, error) {
	query := fmt.Sprintf("DELETE FROM \"%s\" WHERE order_uid=$1", orderTable)
	_, err := or.db.Exec(query, orderUid)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", core.NewNotFundErr(http.StatusNotFound, "order")
		}

		return "", core.NewCantCreateErr(http.StatusBadRequest, "order")
	}

	return orderUid, nil
}
