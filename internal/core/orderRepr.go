package core

type DeliveryRepr struct {
	Id      int64  `json:"-" valid:"-"`
	Name    string `json:"name" valid:"type(string), required"`
	Phone   string `json:"phone" valid:"type(string), required"`
	Zip     string `json:"zip" valid:"type(string), required"`
	City    string `json:"city" valid:"type(string), required"`
	Address string `json:"address" valid:"type(string), required"`
	Region  string `json:"region" valid:"type(string), required"`
	Email   string `json:"email" valid:"email, required"`
}

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

type ProductRepr struct {
	Id          int64  `json:"chrt_id" valid:"int, required"`
	TrackNumber string `json:"track_number" valid:"type(string), required"`
	Price       int64  `json:"price" valid:"int, required"`
	Rid         string `json:"rid" valid:"type(string), required"`
	Name        string `json:"name" valid:"type(string), required"`
	Sale        int64  `json:"sale" valid:"int, required"`
	Size        string `json:"size" valid:"type(string), required"`
	TotalPrice  int64  `json:"total_price" valid:"int, required"`
	NmId        int64  `json:"nm_id" valid:"int, required"`
	Brand       string `json:"brand" valid:"type(string), required"`
	Status      int32  `json:"status" valid:"int, required"`
}

type OrderRepr struct {
	Id                int64         `json:"-"`
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
