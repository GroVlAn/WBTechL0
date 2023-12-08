package postgresrepos

import (
	"github.com/GroVlAn/WBTechL0/internal/core"
	"github.com/GroVlAn/WBTechL0/internal/repository/postgresrepos/repos"
	"github.com/jmoiron/sqlx"
)

type PostgresRepos struct {
	DeliveryRepository
	PaymentRepository
	ProductRepository
	OrderRepository
}

func NewPostgresRepos(db *sqlx.DB) *PostgresRepos {
	return &PostgresRepos{
		ProductRepository:  repos.NewProductRepos(db),
		PaymentRepository:  repos.NewPaymentRepos(db),
		DeliveryRepository: repos.NewDeliveryRepos(db),
		OrderRepository:    repos.NewOrderRepos(db),
	}
}

type DeliveryRepository interface {
	Delivery(id int64) (core.Delivery, error)
	Delete(id int) (int, error)
}

type PaymentRepository interface {
	Payment(tran string) (core.Payment, error)
	Delete(tran string) (string, error)
}

type ProductRepository interface {
	Create(prod core.Product) (int, error)
	FindByTrackNumber(trNumb string) ([]core.Product, error)
	Product(id int) (core.Product, error)
	Delete(id int) (int, error)
}

type OrderRepository interface {
	Create(ord core.Order, d core.Delivery, pmt core.Payment) (string, error)
	Order(orderUid string) (core.Order, error)
	Delete(orderUid string) (string, error)
}
