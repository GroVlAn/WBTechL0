package service

import (
	"errors"
	prepos "github.com/GroVlAn/WBTechL0/internal/repository/postgresrepos"
	"github.com/asaskevich/govalidator"
	"github.com/sirupsen/logrus"
)

type DeliveryRepr struct {
	Name    string `json:"name" valid:"type(string), required"`
	Phone   string `json:"phone" valid:"type(string), required"`
	Zip     string `json:"zip" valid:"type(string), required"`
	City    string `json:"city" valid:"type(string), required"`
	Address string `json:"address" valid:"type(string), required"`
	Region  string `json:"region" valid:"type(string), required"`
	Email   string `json:"email" valid:"email, required"`
}

type DeliveryServ struct {
	repos prepos.DeliveryRepository
}

func NewDeliveryServ(repos prepos.DeliveryRepository) *DeliveryServ {
	return &DeliveryServ{
		repos: repos,
	}
}

func (ds *DeliveryServ) CreateDelivery(dRepr DeliveryRepr) (int, error) {
	result, err := govalidator.ValidateStruct(dRepr)

	if err != nil {
		logrus.Errorln(err.Error())
	}

	if !result {
		return -1, errors.New("invalid data")
	}

	return -1, nil
}

func (ds *DeliveryServ) Delivery(id int) (DeliveryRepr, error) {
	return DeliveryRepr{}, nil
}

func (ds *DeliveryServ) DeleteDelivery(id int) (int, error) {
	return -1, nil
}
