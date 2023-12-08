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
	Delivery: DeliveryRepr{
		Name:    "Test Testov",
		Phone:   "+9720000000",
		Zip:     "2639809",
		City:    "Kiryat Mozkin",
		Address: "Ploshad Mira 15",
		Region:  "Kraiot",
		Email:   "test@gmail.com",
	},
	Payment: PaymentRepr{
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

type OrderRepr struct {
	OrderUid          string        `json:"order_uid"`
	TrackNumber       string        `json:"track_number"`
	Entry             string        `json:"entry"`
	Delivery          DeliveryRepr  `json:"delivery"`
	Payment           PaymentRepr   `json:"payment"`
	Items             []ProductRepr `json:"items"`
	Locale            string        `json:"locale"`
	InternalSignature string        `json:"internal_signature"`
	CustomerId        string        `json:"customer_id"`
	DeliveryService   string        `json:"delivery_service"`
	Shardkey          string        `json:"shardkey"`
	SmId              int64         `json:"sm_id"`
	OofShard          string        `json:"off_shard"`
	DateCreated       string        `json:"date_created"`
}

type OrderReq struct {
	Id                int          `json:"-" valid:"-"`
	TrackNumber       string       `json:"track_number" valid:"-"`
	Entry             string       `json:"entry" valid:"-"`
	Delivery          DeliveryRepr `json:"delivery" valid:"-"`
	Payment           PaymentRepr  `json:"payment" valid:"-"`
	Locale            string       `json:"locale" valid:"-"`
	InternalSignature string       `json:"internal_signature" valid:"-"`
	CustomerId        string       `json:"customer_id" valid:"-"`
	DeliveryService   string       `json:"delivery_service" valid:"-""`
	Shardkey          string       `json:"shardkey" valid:"-"`
	SmId              int64        `json:"sm_id" valid:"-"`
	OffShard          string       `json:"off_shard" valid:"-"`
}

type OrderServ struct {
	log      *logrus.Logger
	ordRepo  prepos.OrderRepository
	dRepos   prepos.DeliveryRepository
	pmtRepo  prepos.PaymentRepository
	prodRepo prepos.ProductRepository
}

func NewOrderServ(
	log *logrus.Logger,
	repos prepos.OrderRepository,
	dRepos prepos.DeliveryRepository,
	pmtRepo prepos.PaymentRepository,
	prodRepo prepos.ProductRepository,
) *OrderServ {
	return &OrderServ{
		log:      log,
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
		ors.log.Errorln("order create: invalid data")
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

	ors.log.Info("service try to create new order")

	return ordUid, errOrd
}

func (ors *OrderServ) Order(ordUid string) (OrderRepr, error) {
	ord, errOrd := ors.ordRepo.Order(ordUid)

	if errOrd != nil {
		ors.log.Errorf("service order: find order error:: %s", errOrd.Error())
		return OrderRepr{}, errOrd
	}

	d, errD := ors.dRepos.Delivery(ord.DeliveryId)

	if errD != nil {
		ors.log.Errorf("service order: find delivery error: %s", errD.Error())
		return OrderRepr{}, errD
	}

	pmt, errPmt := ors.pmtRepo.Payment(ord.OrderUid)

	if errPmt != nil {
		ors.log.Errorf("service order: find payment error: %s", errPmt.Error())
		return OrderRepr{}, errPmt
	}

	prods, errProds := ors.prodRepo.FindByTrackNumber(ord.TrackNumber)

	if errProds != nil {
		ors.log.Errorf("service order: find products error: %s", errProds.Error())
		return OrderRepr{}, errProds
	}

	prodsReprs := make([]ProductRepr, 0, len(prods))

	for _, prod := range prods {
		prodsReprs = append(prodsReprs, ProductRepr(prod))
	}

	ordRepr := OrderRepr{
		OrderUid:          ord.OrderUid,
		TrackNumber:       ord.TrackNumber,
		Entry:             ord.Entry,
		Delivery:          DeliveryRepr(d),
		Payment:           PaymentRepr(pmt),
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

	ors.log.Infof("service order try delete order by order_uid: %s", ordUid)

	return delOrdUid, err
}
