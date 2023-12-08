package service

import (
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
	prodRepos prepos.ProductRepository,
	pmtRepos prepos.PaymentRepository,
	dRepos prepos.DeliveryRepository,
	ordRepos prepos.OrderRepository,
) *Service {
	return &Service{
		ProductService:  NewProductServ(log, prodRepos),
		PaymentService:  NewPaymentServ(log, pmtRepos),
		DeliveryService: NewDeliveryServ(log, dRepos),
		OrderService:    NewOrderServ(log, ordRepos, dRepos, pmtRepos, prodRepos),
	}
}

type ProductService interface {
	CreateProduct(prodRpr ProductRepr) (int, error)
	Product(id int) (ProductRepr, error)
	All(trNum string) ([]ProductRepr, error)
	DeleteProduct(id int) (int, error)
}

type PaymentService interface {
	Payment(tran string) (PaymentRepr, error)
	DeletePayment(tran string) (string, error)
}

type DeliveryService interface {
	Delivery(id int64) (DeliveryRepr, error)
	DeleteDelivery(id int) (int, error)
}

type OrderService interface {
	CreateOrder(orReq OrderReq) (string, error)
	Order(ordUid string) (OrderRepr, error)
	DeleteOrder(ordUid string) (string, error)
}
