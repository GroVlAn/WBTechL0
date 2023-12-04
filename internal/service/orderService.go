package service

import (
	prepos "github.com/GroVlAn/WBTechL0/internal/repository/postgresrepos"
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
	OrderUid          string    `json:"order_uid"`
	TrackNumber       string    `json:"track_number"`
	Entry             string    `json:"entry"`
	DeliveryID        int       `json:"delivery"`
	PaymentID         int       `json:"payment"`
	ItemsIDs          []int     `json:"items"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerId        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	Shardkey          string    `json:"shardkey"`
	SmId              int64     `json:"sm_id"`
	OffShard          string    `json:"off_shard"`
	DateCreated       time.Time `json:"date_created"`
}

type OrderServ struct {
	repos *prepos.OrderRepository
}

func NewOrderServ(repos *prepos.OrderRepository) *OrderServ {
	return &OrderServ{
		repos: repos,
	}
}

func (ors *OrderServ) CreateOrder(orReq OrderReq) (int, error) {
	return -1, nil
}

func (ors *OrderServ) Order(id int) (OrderRepr, error) {
	return OrderRepr{}, nil
}

func (ors *OrderServ) DeleteOrder(id int) (int, error) {
	return -1, nil
}
