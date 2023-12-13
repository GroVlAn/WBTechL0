package service

import (
	"github.com/GroVlAn/WBTechL0/internal/core"
	prepos "github.com/GroVlAn/WBTechL0/internal/repository/postgresrepos"
	"github.com/sirupsen/logrus"
)

type PaymentServ struct {
	log   *logrus.Logger
	ch    Cacher
	repos prepos.PaymentRepository
}

func NewPaymentServ(log *logrus.Logger, ch Cacher, repos prepos.PaymentRepository) *PaymentServ {
	return &PaymentServ{
		log:   log,
		ch:    ch,
		repos: repos,
	}
}

func (ps *PaymentServ) Payment(tran string) (core.PaymentRepr, error) {
	pmtCh, err := ps.ch.Payment(tran)

	if err == nil {
		ps.log.Infof("service payment find in cache payment by transaction: %s", tran)
		return pmtCh, nil
	}

	pmt, err := ps.repos.Payment(tran)

	ps.log.Infof("service payment try to find by transactoion: %s", tran)

	return core.PaymentRepr(pmt), err
}

func (ps *PaymentServ) DeletePayment(tran string) (string, error) {
	delPmtId, err := ps.repos.Delete(tran)

	if err != nil {
		ps.log.Errorf("service payment delete: can not delete payment: %s", err.Error())
		return "", err
	}

	ps.ch.DeletePayment(tran)
	ps.log.Infof("service payment delete by transactoion: %s", tran)

	return delPmtId, nil
}
