package natssub

import (
	"encoding/json"
	"fmt"
	"github.com/GroVlAn/WBTechL0/internal/config"
	"github.com/GroVlAn/WBTechL0/internal/core"
	"github.com/GroVlAn/WBTechL0/internal/service"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

const (
	ChannelOrd  = "order"
	ChannelProd = "product"
)

type Subscribe struct {
	Sc       stan.Conn
	log      *logrus.Logger
	prodServ service.ProductService
	orServ   service.OrderService
}

func NewSubscribe(
	conf config.Config,
	log *logrus.Logger,
	prodServ service.ProductService,
	orServ service.OrderService,
) *Subscribe {
	url := stan.NatsURL(fmt.Sprintf("nats://%s:%s", conf.NatsConfig.NatsServer, conf.NatsConfig.Port))
	conWait := stan.ConnectWait(conf.NatsConfig.ConnectionWait)
	sc, err := stan.Connect(conf.NatsConfig.ClusterID, conf.NatsConfig.ClientID, url, conWait)

	if err != nil {
		log.Errorf(`can not connect to nats streaming:
							cluster id: %s, error: %s`,
			conf.NatsConfig.ClusterID,
			err.Error(),
		)
		return nil
	}

	return &Subscribe{
		Sc:       sc,
		log:      log,
		prodServ: prodServ,
		orServ:   orServ,
	}
}

func (s *Subscribe) Run() {
	if s == nil {
		return
	}
	go func() {
		s.SubProduct()
	}()

	go func() {
		s.SubOrder()
	}()
}

func (s *Subscribe) SubProduct() {
	_, err := s.Sc.QueueSubscribe(ChannelProd, "product", func(m *stan.Msg) {
		var prodRepr core.ProductRepr
		errUnm := json.Unmarshal(m.Data, &prodRepr)

		if errUnm != nil {
			s.log.Errorf("can not unmarshal product: %s", errUnm.Error())
		} else {
			id, errCprod := s.prodServ.CreateProduct(prodRepr)

			if errCprod != nil {
				s.log.Errorf("can not create product: %s", errCprod.Error())
			} else {
				s.log.Infof("product by id %d is created", id)
			}
		}
	})

	if err != nil {
		s.log.Errorf("can not subscibe product channel: %s", err.Error())
		return
	}
}

func (s *Subscribe) SubOrder() {
	_, err := s.Sc.QueueSubscribe(ChannelOrd, "order", func(m *stan.Msg) {
		var ordReq service.OrderReq
		errUnm := json.Unmarshal(m.Data, &ordReq)

		if errUnm != nil {
			s.log.Errorf("can not unmarshal order: %s", errUnm.Error())
		} else {
			id, errCOrd := s.orServ.CreateOrder(ordReq)

			if errCOrd != nil {
				s.log.Errorf("can not create order: %s", errCOrd.Error())
			} else {
				s.log.Infof("order by id %s is created", id)
			}
		}
	}, stan.DurableName("order"))

	if err != nil {
		s.log.Errorf("can not subscibe order channel: %s", err.Error())
		return
	}
}
