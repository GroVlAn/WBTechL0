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
	paymentTable = "payment"
)

type PaymentRepos struct {
	log *logrus.Logger
	db  *sqlx.DB
}

func NewPaymentRepos(log *logrus.Logger, db *sqlx.DB) *PaymentRepos {
	return &PaymentRepos{
		log: log,
		db:  db,
	}
}

func (pr *PaymentRepos) Payment(tran string) (core.Payment, error) {
	var payment core.Payment
	query := fmt.Sprintf("SELECT * FROM %s WHERE transaction=$1", paymentTable)
	err := pr.db.Get(&payment, query, tran)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			pr.log.Errorf("error get payment: not found: %s", err.Error())
			return core.Payment{}, core.NewNotFundErr(http.StatusNotFound, "payment")
		}

		pr.log.Errorf("error get payment: can not find payment: %s", err.Error())
		return core.Payment{}, core.NewCantCreateErr(http.StatusBadRequest, "payment")
	}

	pr.log.Infof("find payment by transaction: %s", tran)
	return payment, nil
}

func (pr *PaymentRepos) Delete(tran string) (string, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE transaction=$1", paymentTable)
	_, err := pr.db.Exec(query, tran)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			pr.log.Errorf("error delete payment: not found: %s", err.Error())

			return "", core.NewNotFundErr(http.StatusNotFound, "payment")
		}

		pr.log.Errorf("error delete payment: can not delete payment: %s", err.Error())
		return "", core.NewCantCreateErr(http.StatusBadRequest, "payment")
	}

	pr.log.Infof("delete payment with transaction: %s", tran)
	return tran, nil
}
