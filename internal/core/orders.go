package core

import "time"

type Product struct {
	Id          int    `json:"chrt_id" db:"chrt_id"`
	TrackNumber string `json:"track_number" db:"track_number"`
	Price       int64  `json:"price" db:"price"`
	Rid         string `json:"rid" db:"rid"`
	Name        string `json:"name" db:"name"`
	Sale        int64  `json:"sale" db:"sale"`
	Size        string `json:"size" db:"size"`
	TotalPrice  int64  `json:"total_price" db:"total_price"`
	NmId        int64  `json:"nm_id" db:"nm_id"`
	Brand       string `json:"brand" db:"brand"`
	Status      int32  `json:"status" db:"status"`
}

type Payment struct {
	Id           int    `json:"-" db:"id"`
	Transaction  string `json:"transaction" db:"transaction"`
	RequestId    string `json:"request_id" db:"request_id"`
	Currency     string `json:"currency" db:"currency"`
	Provider     string `json:"provider" db:"provider"`
	Amount       int64  `json:"amount" db:"amount"`
	PaymentDt    int64  `json:"payment_dt" db:"payment_dt"`
	Bank         string `json:"bank" db:"bank"`
	DeliveryCost int64  `json:"delivery_cost" db:"delivery_cost"`
	GoodsTotal   int64  `json:"goods_total" db:"goods_total"`
	CustomFee    int64  `json:"custom_fee" db:"custom_fee" db:"custom_fee"`
}

type Delivery struct {
	Id      int    `json:"_" db:"id"`
	Name    string `json:"name" db:"name"`
	Phone   string `json:"phone" db:"phone"`
	Zip     string `json:"zip" db:"zip"`
	City    string `json:"city" db:"city"`
	Address string `json:"address" db:"address"`
	Region  string `json:"region" db:"region"`
	Email   string `json:"email" db:"email"`
}

type Order struct {
	Id                int       `json:"-" db:"id"`
	OrderUid          string    `json:"order_uid" db:"order_uid"`
	TrackNumber       string    `json:"track_number" db:"track_number"`
	Entry             string    `json:"entry" db:"entry"`
	Locale            string    `json:"locale" db:"locale"`
	InternalSignature string    `json:"internal_signature" db:"internal_signature"`
	CustomerId        string    `json:"customer_id" db:"customer_id"`
	DeliveryService   string    `json:"delivery_service" db:"delivery_service"`
	Shardkey          string    `json:"shardkey" id:"shardkey"`
	SmId              int64     `json:"sm_id" db:"sm_id"`
	OffShard          string    `json:"off_shard" db:"off_shard"`
	DateCreated       time.Time `json:"date_created" db:"date_created"`
	DeliveryId        string    `json:"delivery_id" db:"delivery_id"`
	PaymentId         string    `json:"payment_id" db:"payment_id"`
}

type OrderProduct struct {
	Id        int
	OrderId   int
	ProductId int
}
