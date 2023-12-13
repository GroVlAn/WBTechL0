package service

import (
	"fmt"
	"github.com/GroVlAn/WBTechL0/internal/core"
	"github.com/GroVlAn/WBTechL0/internal/repository/cacherepo"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

type TestDeliveryRepository struct{}

var MockDeliveries = []core.Delivery{
	{
		Id:      1,
		Name:    "Test Testov",
		Phone:   "+9720000000",
		Zip:     "2639809",
		City:    "Kiryat Mozkin",
		Address: "Ploshad Mira 15",
		Region:  "Kraiot",
		Email:   "test@gmail.com",
	},
	{
		Id:      2,
		Name:    "Test1 Testov1",
		Phone:   "+9720000002",
		Zip:     "2634209",
		City:    "Moscow",
		Address: "Ploshad Mira 12",
		Region:  "Kraioh",
		Email:   "test2@gmail.com",
	},
}

func NewDeliveryRepo() *TestDeliveryRepository {
	return &TestDeliveryRepository{}
}

func (dr *TestDeliveryRepository) Delivery(id int64) (core.Delivery, error) {
	for _, item := range MockDeliveries {
		if item.Id == id {
			return item, nil
		}
	}

	return core.Delivery{}, core.NewNotFundErr(http.StatusNotFound, "delivery")
}
func (dr *TestDeliveryRepository) Delete(id int64) (int64, error) {
	var iDel int64 = -1

	for key, item := range MockDeliveries {
		if item.Id == id {
			iDel = int64(key)
		}
	}

	if iDel == -1 {
		return -1, core.NewNotFundErr(http.StatusNotFound, "delivery")
	}

	MockDeliveries = append(MockDeliveries[:iDel], MockDeliveries[iDel+1:]...)

	return id, nil
}
func (dr *TestDeliveryRepository) All() ([]core.Delivery, error) {
	if len(MockDeliveries) == 0 {
		return nil, core.NewNotFundErr(http.StatusNotFound, "delivery")
	}
	return MockDeliveries, nil
}

func TestDeliveryServ(t *testing.T) {

	logTest := logrus.New()
	repo := NewDeliveryRepo()

	defer CaptureLog(t).Release()
	ch := cacherepo.NewCache(logTest)
	delServ := NewDeliveryServ(logTest, ch, repo)

	d1, errD1 := delServ.Delivery(1)

	assert.NoError(t, errD1)
	assert.Equal(t, MockDeliveries[0], core.Delivery(d1))

	d2, errD2 := delServ.Delivery(int64(3))
	assert.Error(t, errD2)
	assert.ErrorIs(t, core.NewNotFundErr(http.StatusNotFound, "delivery"), errD2)
	assert.Equal(t, core.Delivery{}, core.Delivery(d2))

	dDel1, errDel1 := delServ.DeleteDelivery(int64(1))
	assert.NoError(t, errDel1)
	assert.Equal(t, dDel1, int64(1))
	assert.Len(t, MockDeliveries, 1)

	dDel2, errDel2 := delServ.DeleteDelivery(int64(2))
	assert.NoError(t, errDel2)
	assert.Equal(t, dDel2, int64(2))
	assert.Len(t, MockDeliveries, 0)

	dDel3, errDel3 := delServ.DeleteDelivery(int64(2))
	fmt.Println(errDel3)
	assert.Error(t, errDel3)
	assert.ErrorIs(t, core.NewNotFundErr(http.StatusNotFound, "delivery"), errDel3)
	assert.Equal(t, dDel3, int64(-1))
}

type logCapturer struct {
	*testing.T
	origOut io.Writer
}

func (tl logCapturer) Write(p []byte) (n int, err error) {
	tl.Logf((string)(p))
	return len(p), nil
}

func (tl logCapturer) Release() {
	logrus.SetOutput(tl.origOut)
}

func CaptureLog(t *testing.T) *logCapturer {
	lc := logCapturer{T: t, origOut: logrus.StandardLogger().Out}
	if !testing.Verbose() {
		logrus.SetOutput(lc)
	}
	return &lc
}
