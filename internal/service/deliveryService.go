package service

import (
	prepos "github.com/GroVlAn/WBTechL0/internal/repository/postgresrepos"
)

type DeliveryRepr struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type DeliveryServ struct {
	repos *prepos.DeliveryRepository
}

func NewDeliveryServ(repos *prepos.DeliveryRepository) *DeliveryServ {
	return &DeliveryServ{
		repos: repos,
	}
}

func (ds *DeliveryServ) CreateDelivery(dRepr DeliveryRepr) (int, error) {
	return -1, nil
}

func (ds *DeliveryServ) Delivery(id int) (DeliveryRepr, error) {
	return DeliveryRepr{}, nil
}

func (ds *DeliveryServ) DeleteDelivery(id int) (int, error) {
	return -1, nil
}
