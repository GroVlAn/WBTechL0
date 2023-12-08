package service

import (
	prepos "github.com/GroVlAn/WBTechL0/internal/repository/postgresrepos"
	"github.com/sirupsen/logrus"
)

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

type DeliveryServ struct {
	log   *logrus.Logger
	repos prepos.DeliveryRepository
}

func NewDeliveryServ(log *logrus.Logger, repos prepos.DeliveryRepository) *DeliveryServ {
	return &DeliveryServ{
		log:   log,
		repos: repos,
	}
}

func (ds *DeliveryServ) Delivery(id int64) (DeliveryRepr, error) {
	d, err := ds.repos.Delivery(id)
	ds.log.Infof("service try to find delivery by id: %d", id)

	return DeliveryRepr(d), err
}

func (ds *DeliveryServ) DeleteDelivery(id int) (int, error) {
	delDId, err := ds.repos.Delete(id)
	ds.log.Infof("service try to delete delivery by id: %d", id)

	return delDId, err
}
