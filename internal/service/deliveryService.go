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
	ch    Cacher
	repos prepos.DeliveryRepository
}

func NewDeliveryServ(log *logrus.Logger, ch Cacher, repos prepos.DeliveryRepository) *DeliveryServ {
	return &DeliveryServ{
		log:   log,
		ch:    ch,
		repos: repos,
	}
}

func (ds *DeliveryServ) Delivery(id int64) (DeliveryRepr, error) {
	dCh, err := ds.ch.Delivery(id)

	if err == nil {
		ds.log.Infof("service delivery find in cache order by order uid: %d", id)
		return dCh, nil
	}

	d, err := ds.repos.Delivery(id)
	ds.log.Infof("service try to find delivery by id: %d", id)

	return DeliveryRepr(d), err
}

func (ds *DeliveryServ) DeleteDelivery(id int64) (int64, error) {
	delDId, err := ds.repos.Delete(id)

	if err != nil {
		ds.log.Errorf("service delivery delete: can not delete delivery: %s", err.Error())
		return -1, err
	}

	ds.ch.DeleteDelivery(id)

	ds.log.Infof("service delete delivery by id: %d", id)

	return delDId, nil
}
