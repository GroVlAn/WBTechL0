package service

import (
	prepos "github.com/GroVlAn/WBTechL0/internal/repository/postgresrepos"
)

type PaymentRepr struct {
	Transaction  string `json:"transaction"`
	RequestId    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int64  `json:"amount"`
	PaymentDt    int64  `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int64  `json:"delivery_cost"`
	GoodsTotal   int64  `json:"goods_total"`
	CustomFee    int64  `json:"custom_fee"`
}

type PaymentServ struct {
	repos *prepos.PaymentRepository
}

func NewPaymentServ(repos *prepos.PaymentRepository) *PaymentServ {
	return &PaymentServ{
		repos: repos,
	}
}

func (ps *PaymentServ) CreatePayment(pmtRepr PaymentRepr) (int, error) {
	return -1, nil
}

func (ps *PaymentServ) Payment(id int) (PaymentRepr, error) {
	return PaymentRepr{}, nil
}

func (ps *PaymentServ) DeletePayment(id int) (int, error) {
	return -1, nil
}
