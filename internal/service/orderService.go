package service

import (
	"github.com/GroVlAn/WBTechL0/internal/core"
	prepos "github.com/GroVlAn/WBTechL0/internal/repository/postgresrepos"
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	dateCreatedFormat = "2006-01-02T03:04:05Z"
)

var ExampleOrderReq = OrderReq{
	TrackNumber: "WBILMTESTTRACK",
	Entry:       "WBIL",
	Delivery: core.DeliveryRepr{
		Name:    "Test Testov",
		Phone:   "+9720000000",
		Zip:     "2639809",
		City:    "Kiryat Mozkin",
		Address: "Ploshad Mira 15",
		Region:  "Kraiot",
		Email:   "test@gmail.com",
	},
	Payment: core.PaymentRepr{
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
	Locale:            "en",
	InternalSignature: "",
	CustomerId:        "test",
	DeliveryService:   "meest",
	Shardkey:          "9",
	SmId:              99,
	OffShard:          "1",
}

type OrderReq struct {
	Id                int               `json:"-" valid:"-"`
	TrackNumber       string            `json:"track_number" valid:"type(string), required"`
	Entry             string            `json:"entry" valid:"type(string), required"`
	Delivery          core.DeliveryRepr `json:"delivery" valid:"-"`
	Payment           core.PaymentRepr  `json:"payment" valid:"-"`
	Locale            string            `json:"locale" valid:"type(string), required"`
	InternalSignature string            `json:"internal_signature" valid:"type(string)"`
	CustomerId        string            `json:"customer_id" valid:"type(string), required"`
	DeliveryService   string            `json:"delivery_service" valid:"type(string), required"`
	Shardkey          string            `json:"shardkey" valid:"type(string), required"`
	SmId              int64             `json:"sm_id" valid:"int, required"`
	OffShard          string            `json:"off_shard" valid:"type(string), required"`
}

type OrderServ struct {
	log      *logrus.Logger
	ch       Cacher
	ordRepo  prepos.OrderRepository
	dRepos   prepos.DeliveryRepository
	pmtRepo  prepos.PaymentRepository
	prodRepo prepos.ProductRepository
}

func NewOrderServ(
	log *logrus.Logger,
	ch Cacher,
	repos prepos.OrderRepository,
	dRepos prepos.DeliveryRepository,
	pmtRepo prepos.PaymentRepository,
	prodRepo prepos.ProductRepository,
) *OrderServ {
	return &OrderServ{
		log:      log,
		ch:       ch,
		ordRepo:  repos,
		dRepos:   dRepos,
		pmtRepo:  pmtRepo,
		prodRepo: prodRepo,
	}
}

func (ors *OrderServ) CreateOrder(ordReq OrderReq) (string, error) {
	result, err := govalidator.ValidateStruct(ordReq)

	if err != nil {
		ors.log.Errorln(err.Error())
		return "", err
	}

	if !result {
		ors.log.Errorln("serice order create: invalid data")
		return "", core.NewInvalidDataErr(http.StatusBadRequest, "order", ExampleOrderReq)
	}

	ordUid := uuid.New().String()

	ordReq.Payment.Transaction = ordUid

	d := core.Delivery(ordReq.Delivery)

	pmt := core.Payment(ordReq.Payment)

	ord := core.Order{
		OrderUid:          ordUid,
		TrackNumber:       ordReq.TrackNumber,
		Entry:             ordReq.Entry,
		Locale:            ordReq.Locale,
		InternalSignature: ordReq.InternalSignature,
		CustomerId:        ordReq.CustomerId,
		DeliveryService:   ordReq.DeliveryService,
		Shardkey:          ordReq.Shardkey,
		SmId:              ordReq.SmId,
		OofShard:          ordReq.OffShard,
	}
	ordUid, errOrd := ors.ordRepo.Create(ord, d, pmt)

	if errOrd != nil {
		ors.log.Errorf("service order create: can not create order: %s", errOrd.Error())
		return "", errOrd
	}

	prods, _ := ors.prodRepo.All()

	prodsReprs := make([]core.ProductRepr, 0)

	for _, item := range prods {
		if item.TrackNumber == ord.TrackNumber {
			prodsReprs = append(prodsReprs, core.ProductRepr(item))
		}
	}

	ordRepr := core.OrderRepr{
		OrderUid:          ord.OrderUid,
		TrackNumber:       ord.TrackNumber,
		Entry:             ord.Entry,
		Delivery:          core.DeliveryRepr(d),
		Payment:           core.PaymentRepr(pmt),
		Items:             prodsReprs,
		Locale:            ord.Locale,
		InternalSignature: ord.InternalSignature,
		CustomerId:        ord.CustomerId,
		DeliveryService:   ord.DeliveryService,
		Shardkey:          ord.Shardkey,
		SmId:              ord.SmId,
		OofShard:          ord.OofShard,
		DateCreated:       ord.DateCreated.Format(dateCreatedFormat),
	}

	ors.ch.SetOrder(ordUid, ordRepr)

	ors.log.Info("service create new order")

	return ordUid, nil
}

func (ors *OrderServ) Order(ordUid string) (core.OrderRepr, error) {
	ordCh, err := ors.ch.Order(ordUid)

	if err == nil {
		ors.log.Infof("service order find in cache order by order uid: %s", ordUid)
		return ordCh, nil
	}

	ord, errOrd := ors.ordRepo.Order(ordUid)

	if errOrd != nil {
		ors.log.Errorf("service order: find order error:: %s", errOrd.Error())
		return core.OrderRepr{}, errOrd
	}

	d, errD := ors.dRepos.Delivery(ord.DeliveryId)

	if errD != nil {
		ors.log.Errorf("service order: find delivery error: %s", errD.Error())
		return core.OrderRepr{}, errD
	}

	pmt, errPmt := ors.pmtRepo.Payment(ord.OrderUid)

	if errPmt != nil {
		ors.log.Errorf("service order: find payment error: %s", errPmt.Error())
		return core.OrderRepr{}, errPmt
	}

	prods, errProds := ors.prodRepo.FindByTrackNumber(ord.TrackNumber)

	if errProds != nil {
		ors.log.Errorf("service order: find products error: %s", errProds.Error())
		return core.OrderRepr{}, errProds
	}

	prodsReprs := make([]core.ProductRepr, 0, len(prods))

	for _, prod := range prods {
		prodsReprs = append(prodsReprs, core.ProductRepr(prod))
	}

	ordRepr := core.OrderRepr{
		OrderUid:          ord.OrderUid,
		TrackNumber:       ord.TrackNumber,
		Entry:             ord.Entry,
		Delivery:          core.DeliveryRepr(d),
		Payment:           core.PaymentRepr(pmt),
		Items:             prodsReprs,
		Locale:            ord.Locale,
		InternalSignature: ord.InternalSignature,
		CustomerId:        ord.CustomerId,
		DeliveryService:   ord.DeliveryService,
		Shardkey:          ord.Shardkey,
		SmId:              ord.SmId,
		OofShard:          ord.OofShard,
		DateCreated:       ord.DateCreated.Format(dateCreatedFormat),
	}

	ors.log.Info("service order: return order")

	return ordRepr, nil
}

func (ors *OrderServ) DeleteOrder(ordUid string) (string, error) {
	delOrdUid, err := ors.ordRepo.Delete(ordUid)

	if err != nil {
		ors.log.Errorf("servie order: can not delete order: %s", err.Error())
		return "", err
	}

	ors.log.Infof("service order delete order by order_uid: %s", ordUid)

	ors.ch.DeleteOrder(ordUid)

	return delOrdUid, nil
}

func (ors *OrderServ) All() ([]core.OrderRepr, error) {
	ords, err := ors.ordRepo.All()

	if err != nil {
		ors.log.Errorf("service order: can no find orders: %s", err.Error())
		return nil, err
	}

	ordsReprs := make([]core.OrderRepr, 0)

	for _, item := range ords {
		pmt, errPmt := ors.pmtRepo.Payment(item.PaymentTransaction)

		if errPmt != nil {
			continue
		}

		d, errD := ors.dRepos.Delivery(item.DeliveryId)

		if errD != nil {
			continue
		}

		prods, errProds := ors.prodRepo.FindByTrackNumber(item.TrackNumber)

		if errProds != nil {
			continue
		}

		prodsReprs := make([]core.ProductRepr, 0)

		for _, prod := range prods {
			prodsReprs = append(prodsReprs, core.ProductRepr(prod))
		}

		ordRepr := core.OrderRepr{
			OrderUid:          item.OrderUid,
			TrackNumber:       item.TrackNumber,
			Entry:             item.Entry,
			Delivery:          core.DeliveryRepr(d),
			Payment:           core.PaymentRepr(pmt),
			Items:             prodsReprs,
			Locale:            item.Locale,
			InternalSignature: item.InternalSignature,
			CustomerId:        item.CustomerId,
			DeliveryService:   item.DeliveryService,
			Shardkey:          item.Shardkey,
			SmId:              item.SmId,
			OofShard:          item.OofShard,
			DateCreated:       item.DateCreated.Format(dateCreatedFormat),
		}
		ordsReprs = append(ordsReprs, ordRepr)
	}

	return ordsReprs, nil
}
