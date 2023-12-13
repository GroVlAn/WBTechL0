package service

import (
	"github.com/GroVlAn/WBTechL0/internal/core"
	"github.com/GroVlAn/WBTechL0/internal/repository/cacherepo"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

type TestPaymentRepository struct{}

var MockPayments = []core.Payment{
	{
		Transaction:  "b563feb7b2b84b6test",
		RequestId:    "",
		Currency:     "USD",
		Provider:     "wbpay",
		Amount:       1817,
		PaymentDt:    1637907727,
		Bank:         "alpha",
		DeliveryCost: 1500,
		GoodsTotal:   317,
		CustomFee:    0,
	},
	{
		Transaction:  "b563feb7b2b84b6test2",
		RequestId:    "",
		Currency:     "RUR",
		Provider:     "wbpay",
		Amount:       1818,
		PaymentDt:    1637907728,
		Bank:         "sber",
		DeliveryCost: 1501,
		GoodsTotal:   318,
		CustomFee:    0,
	},
}

func NewTestRepository() *TestPaymentRepository {
	return &TestPaymentRepository{}
}

func (dr *TestPaymentRepository) Payment(tran string) (core.Payment, error) {
	for _, item := range MockPayments {
		if item.Transaction == tran {
			return item, nil
		}
	}

	return core.Payment{}, core.NewNotFundErr(http.StatusNotFound, "payment")
}
func (dr *TestPaymentRepository) Delete(tran string) (string, error) {
	var iPmt int64 = -1

	for key, item := range MockPayments {
		if item.Transaction == tran {
			iPmt = int64(key)
		}
	}

	if iPmt == -1 {
		return "", core.NewNotFundErr(http.StatusNotFound, "payment")
	}

	MockPayments = append(MockPayments[:iPmt], MockPayments[iPmt+1:]...)

	return tran, nil
}
func (dr *TestPaymentRepository) All() ([]core.Payment, error) {
	if len(MockPayments) == 0 {
		return nil, core.NewNotFundErr(http.StatusNotFound, "payment")
	}
	return MockPayments, nil
}

func TestPaymentServ(t *testing.T) {

	logTest := logrus.New()
	repo := NewTestRepository()

	defer CaptureLog(t).Release()
	ch := cacherepo.NewCache(logTest)

	pmtServ := NewPaymentServ(logTest, ch, repo)

	pmtRepr1, errPmt1 := pmtServ.Payment("b563feb7b2b84b6test2")

	assert.NoError(t, errPmt1)
	assert.Equal(t, core.PaymentRepr(MockPayments[1]), pmtRepr1)

	pmtRepr2, errPmt2 := pmtServ.Payment("b563feb7b2b84b6test3")

	assert.Error(t, errPmt2)
	assert.ErrorIs(t, core.NewNotFundErr(http.StatusNotFound, "payment"), errPmt2)
	assert.Equal(t, core.PaymentRepr{}, pmtRepr2)

	pmtDel1, errPmtDel1 := pmtServ.DeletePayment("b563feb7b2b84b6test2")

	assert.Len(t, MockPayments, 1)
	assert.Equal(t, "b563feb7b2b84b6test2", pmtDel1)
	assert.NoError(t, errPmtDel1)

	pmtDel2, errPmtDel2 := pmtServ.DeletePayment("b563feb7b2b84b6test")

	assert.Len(t, MockPayments, 0)
	assert.Equal(t, "b563feb7b2b84b6test", pmtDel2)
	assert.NoError(t, errPmtDel2)

	pmtDel3, errPmtDel3 := pmtServ.DeletePayment("b563feb7b2b84b6test")

	assert.Equal(t, "", pmtDel3)
	assert.Error(t, errPmtDel3)
	assert.ErrorIs(t, core.NewNotFundErr(http.StatusNotFound, "payment"), errPmtDel3)
}
