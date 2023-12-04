package repos

import (
	"fmt"
	"github.com/GroVlAn/WBTechL0/internal/core"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

const (
	paymentTable = "payment"
)

type PaymentRepos struct {
	db *sqlx.DB
}

func NewPaymentRepos(db *sqlx.DB) *PaymentRepos {
	return &PaymentRepos{
		db: db,
	}
}

func (pr *PaymentRepos) Create(pmt core.Payment) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s ("+
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
		"($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id", paymentTable)
	row := pr.db.QueryRow(
		query,
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

	if err := row.Scan(&id); err != nil {
		return -1, errors.Wrap(err, "error can not create payment")
	}

	return id, nil
}

func (pr *PaymentRepos) Payment(id int) (core.Payment, error) {
	var payment core.Payment
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", paymentTable)
	err := pr.db.Get(&payment, query, id)

	return payment, errors.Wrap(err, "payment not found")
}

func (pr *PaymentRepos) Delete(id int) (int, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", paymentTable)
	_, err := pr.db.Exec(query, id)

	if err != nil {
		return -1, errors.Wrap(err, "payment not found")
	}

	return id, nil
}
