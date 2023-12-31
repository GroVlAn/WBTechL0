package postgresrepos

import (
	"github.com/GroVlAn/WBTechL0/internal/core"
	"github.com/GroVlAn/WBTechL0/internal/repository/postgresrepos/repos"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type PostgresRepos struct {
	log *logrus.Logger
	DeliveryRepository
	PaymentRepository
	ProductRepository
	OrderRepository
}

func NewPostgresRepos(log *logrus.Logger, db *sqlx.DB) *PostgresRepos {
	return &PostgresRepos{
		ProductRepository:  repos.NewProductRepos(log, db),
		PaymentRepository:  repos.NewPaymentRepos(log, db),
		DeliveryRepository: repos.NewDeliveryRepos(log, db),
		OrderRepository:    repos.NewOrderRepos(log, db),
	}
}

type DeliveryRepository interface {
	Delivery(id int64) (core.Delivery, error)
	Delete(id int64) (int64, error)
	All() ([]core.Delivery, error)
}

type PaymentRepository interface {
	Payment(tran string) (core.Payment, error)
	Delete(tran string) (string, error)
	All() ([]core.Payment, error)
}

type ProductRepository interface {
	Create(prod core.Product) (int64, error)
	FindByTrackNumber(trNumb string) ([]core.Product, error)
	Product(id int64) (core.Product, error)
	Delete(id int64) (int64, error)
	All() ([]core.Product, error)
}

type OrderRepository interface {
	Create(ord core.Order, d core.Delivery, pmt core.Payment) (string, error)
	Order(orderUid string) (core.Order, error)
	Delete(orderUid string) (string, error)
	All() ([]core.Order, error)
}
