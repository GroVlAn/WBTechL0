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

func (pr *PaymentRepos) Payment(tran string) (core.Payment, error) {
	var payment core.Payment
	query := fmt.Sprintf("SELECT * FROM %s WHERE transaction=$1", paymentTable)
	err := pr.db.Get(&payment, query, tran)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return core.Payment{}, core.NewNotFundErr(http.StatusNotFound, "payment")
		}

		return core.Payment{}, core.NewCantCreateErr(http.StatusBadRequest, "payment")
	}

	return payment, nil
}

func (pr *PaymentRepos) Delete(tran string) (string, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE transaction=$1", paymentTable)
	_, err := pr.db.Exec(query, tran)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", core.NewNotFundErr(http.StatusNotFound, "payment")
		}

		return "", core.NewCantCreateErr(http.StatusBadRequest, "payment")
	}

	return tran, nil
}
