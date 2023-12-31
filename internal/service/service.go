package service

import (
	"github.com/GroVlAn/WBTechL0/internal/core"
	prepos "github.com/GroVlAn/WBTechL0/internal/repository/postgresrepos"
	"github.com/sirupsen/logrus"
)

type Service struct {
	ProductService
	PaymentService
	DeliveryService
	OrderService
}

func NewService(
	log *logrus.Logger,
	ch Cacher,
	prodRepos prepos.ProductRepository,
	pmtRepos prepos.PaymentRepository,
	dRepos prepos.DeliveryRepository,
	ordRepos prepos.OrderRepository,
) *Service {
	return &Service{
		ProductService:  NewProductServ(log, ch, prodRepos),
		PaymentService:  NewPaymentServ(log, ch, pmtRepos),
		DeliveryService: NewDeliveryServ(log, ch, dRepos),
		OrderService:    NewOrderServ(log, ch, ordRepos, dRepos, pmtRepos, prodRepos),
	}
}

type ProductService interface {
	CreateProduct(prodRpr core.ProductRepr) (int64, error)
	Product(id int64) (core.ProductRepr, error)
	All(trNum string) ([]core.ProductRepr, error)
	DeleteProduct(id int64) (int64, error)
}

type PaymentService interface {
	Payment(tran string) (core.PaymentRepr, error)
	DeletePayment(tran string) (string, error)
}

type DeliveryService interface {
	Delivery(id int64) (core.DeliveryRepr, error)
	DeleteDelivery(id int64) (int64, error)
}

type OrderService interface {
	CreateOrder(ordReq OrderReq) (string, error)
	Order(ordUid string) (core.OrderRepr, error)
	DeleteOrder(ordUid string) (string, error)
	All() ([]core.OrderRepr, error)
}

type Cacher interface {
	SetDelivery(id int64, dRepr core.DeliveryRepr)
	Delivery(id int64) (core.DeliveryRepr, error)
	DeleteDelivery(id int64)
	SetOrder(ordUid string, ordRepr core.OrderRepr)
	Order(ordUid string) (core.OrderRepr, error)
	DeleteOrder(ordUid string)
	SetPayment(tran string, pmtRepr core.PaymentRepr)
	Payment(tran string) (core.PaymentRepr, error)
	DeletePayment(tran string)
	SetProduct(id int64, prodRepr core.ProductRepr)
	Product(id int64) (core.ProductRepr, error)
	DeleteProduct(id int64)
}
