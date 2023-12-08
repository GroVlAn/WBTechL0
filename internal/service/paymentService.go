package service

import (
	prepos "github.com/GroVlAn/WBTechL0/internal/repository/postgresrepos"
)

type PaymentRepr struct {
	Id           int64  `json:"-" db:"id"`
	Transaction  string `json:"transaction" valid:"-"`
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

func (ps *PaymentServ) Payment(tran string) (PaymentRepr, error) {
	pmt, err := ps.repos.Payment(tran)

	return PaymentRepr(pmt), err
}

func (ps *PaymentServ) DeletePayment(tran string) (string, error) {
	delPmtId, err := ps.repos.Delete(tran)

	return delPmtId, err
}
