package service

import (
	"errors"
	prepos "github.com/GroVlAn/WBTechL0/internal/repository/postgresrepos"
	"github.com/asaskevich/govalidator"
	"github.com/sirupsen/logrus"
)

type PaymentRepr struct {
	Transaction  string `json:"transaction" valid:"type(string), required"`
	RequestId    string `json:"request_id" valid:"type(string)"`
	Currency     string `json:"currency" valid:"type(string), required"`
	Provider     string `json:"provider" valid:"type(string), required"`
	Amount       int64  `json:"amount" valid:"int, required"`
	PaymentDt    int64  `json:"payment_dt" valid:"int, required"`
	Bank         string `json:"bank" valid:"type(string), required"`
	DeliveryCost int64  `json:"delivery_cost" valid:"int, required"`
	GoodsTotal   int64  `json:"goods_total" valid:"int, required"`
	CustomFee    int64  `json:"custom_fee" valid:"int"`
}

type PaymentServ struct {
	repos prepos.PaymentRepository
}

func NewPaymentServ(repos prepos.PaymentRepository) *PaymentServ {
	return &PaymentServ{
		repos: repos,
	}
}

func (ps *PaymentServ) CreatePayment(pmtRepr PaymentRepr) (int, error) {
	result, err := govalidator.ValidateStruct(pmtRepr)

	if err != nil {
		logrus.Errorln(err.Error())
	}

	if !result {
		return -1, errors.New("invalid data")
	}

	return -1, nil
}

func (ps *PaymentServ) Payment(id int) (PaymentRepr, error) {
	return PaymentRepr{}, nil
}

func (ps *PaymentServ) DeletePayment(id int) (int, error) {
	return -1, nil
}
