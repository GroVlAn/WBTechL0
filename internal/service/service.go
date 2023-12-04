package service

import prepos "github.com/GroVlAn/WBTechL0/internal/repository/postgresrepos"

type Service struct {
	ProductService
	PaymentService
	DeliveryService
	OrderService
}

func NewService(
	prodRepos *prepos.ProductRepository,
	pmtRepos *prepos.PaymentRepository,
	dRepos *prepos.DeliveryRepository,
	orRepos *prepos.OrderRepository,
) *Service {
	return &Service{
		ProductService:  NewProductServ(prodRepos),
		PaymentService:  NewPaymentServ(pmtRepos),
		DeliveryService: NewDeliveryServ(dRepos),
		OrderService:    NewOrderServ(orRepos),
	}
}

type ProductService interface {
	CreateProduct(prodRpr ProductRepr) (int, error)
	Product(id int) (ProductRepr, error)
	DeleteProduct(id int) (int, error)
}

type PaymentService interface {
	CreatePayment(pmtRepr PaymentRepr) (int, error)
	Payment(id int) (PaymentRepr, error)
	DeletePayment(id int) (int, error)
}

type DeliveryService interface {
	CreateDelivery(dRepr DeliveryRepr) (int, error)
	Delivery(id int) (DeliveryRepr, error)
	DeleteDelivery(id int) (int, error)
}

type OrderService interface {
	CreateOrder(orReq OrderReq) (int, error)
	Order(id int) (OrderRepr, error)
	DeleteOrder(id int) (int, error)
}
