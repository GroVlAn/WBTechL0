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
		DeliveryRepository: repos.NewDeliveryRepos(db),
		PaymentRepository:  repos.NewPaymentRepos(db),
		ProductRepository:  repos.NewProductRepos(db),
		OrderRepository:    repos.NewOrderRepos(db),
	}
}

type DeliveryRepository interface {
	Create(d core.Delivery) (int, error)
	Delivery(id int) (core.Delivery, error)
	Delete(id int) (int, error)
}

type PaymentRepository interface {
	Create(pmt core.Payment) (int, error)
	Payment(id int) (core.Payment, error)
	Delete(id int) (int, error)
}

type ProductRepository interface {
	Create(prod core.Product) (int, error)
	Product(id int) (core.Product, error)
	Delete(id int) (int, error)
}

type OrderRepository interface {
	Create(ord core.Order, dId int, pmtId int, prodId int) (int, error)
	Order(id int) (core.Order, error)
	Delete(id int) (int, error)
}
