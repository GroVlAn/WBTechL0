package service

import (
	"errors"
	prepos "github.com/GroVlAn/WBTechL0/internal/repository/postgresrepos"
	"github.com/asaskevich/govalidator"
	"github.com/sirupsen/logrus"
	"time"
)

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
	OffShard          string        `json:"off_shard"`
	DateCreated       time.Time     `json:"date_created"`
}

type OrderReq struct {
	OrderUid          string    `json:"order_uid" valid:"type(string), required"`
	TrackNumber       string    `json:"track_number" valid:"type(string), required"`
	Entry             string    `json:"entry" valid:"type(string), required"`
	DeliveryID        int       `json:"delivery" valid:"int, required"`
	PaymentID         int       `json:"payment" valid:"int, required"`
	ItemsIDs          []int     `json:"items" valid:"type([]string), required"`
	Locale            string    `json:"locale" valid:"type(string), required"`
	InternalSignature string    `json:"internal_signature" valid:"type(string)"`
	CustomerId        string    `json:"customer_id" valid:"type(string), required"`
	DeliveryService   string    `json:"delivery_service" valid:"type(string), required"`
	Shardkey          string    `json:"shardkey" valid:"type(string), required"`
	SmId              int64     `json:"sm_id" valid:"int, required"`
	OffShard          string    `json:"off_shard" valid:"type(string), required"`
	DateCreated       time.Time `json:"date_created" valid:"type(time.Time), required"`
}

type OrderServ struct {
	repos prepos.OrderRepository
}

func NewOrderServ(repos prepos.OrderRepository) *OrderServ {
	return &OrderServ{
		repos: repos,
	}
}

func (ors *OrderServ) CreateOrder(orReq OrderReq) (int, error) {
	result, err := govalidator.ValidateStruct(orReq)

	if err != nil {
		logrus.Errorln(err.Error())
	}

	if !result {
		return -1, errors.New("invalid data")
	}

	return -1, nil
}

func (ors *OrderServ) Order(id int) (OrderRepr, error) {
	return OrderRepr{}, nil
}

func (ors *OrderServ) DeleteOrder(id int) (int, error) {
	return -1, nil
}
