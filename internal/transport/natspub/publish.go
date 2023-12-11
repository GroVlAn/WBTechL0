package natspub

import (
	"encoding/json"
	"fmt"
	"github.com/GroVlAn/WBTechL0/internal/config"
	"github.com/GroVlAn/WBTechL0/internal/service"
	"github.com/go-faker/faker/v4"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

const (
	ChannelOrd   = "order"
	ChannelProd  = "product"
	DelayPublish = 1 * time.Minute
)

type Publish struct {
	Sc  stan.Conn
	log *logrus.Logger
}

func NewPublish(conf config.Config, log *logrus.Logger) *Publish {
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

	return &Publish{
		Sc:  sc,
		log: log,
	}
}

func (p *Publish) Run() {
	go func() {
		for {
			p.PublishProduct()
			time.Sleep(DelayPublish)
		}
	}()

	go func() {
		for {
			p.PublishOrder()
			time.Sleep(DelayPublish)
		}
	}()
}

func (p *Publish) PublishOrder() {
	var ordFake service.OrderReq

	err := faker.FakeData(&ordFake)

	if err != nil {
		p.log.Errorf("error can not create fake order: %s", err.Error())
		return
	}

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	randTNAddr := random.Intn(5) + 1

	ordFake.TrackNumber = fmt.Sprintf("WBILMTESTTRACK%d", randTNAddr)

	ordB, errMarsh := json.Marshal(ordFake)

	if errMarsh != nil {
		p.log.Errorf("can not marshal fake order: %s", errMarsh.Error())
		return
	}

	errPub := p.Sc.Publish(ChannelOrd, ordB)

	if errPub != nil {
		p.log.Errorf("can not publish order")
		return
	}

	p.log.Infof("publish new order with track number: %s", ordFake.TrackNumber)
}

func (p *Publish) PublishProduct() {
	var prodFake service.ProductRepr

	err := faker.FakeData(&prodFake)

	if err != nil {
		p.log.Errorf("error can not create fake product: %s", err.Error())
		return
	}

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	randTNAddr := random.Intn(5) + 1

	prodFake.TrackNumber = fmt.Sprintf("WBILMTESTTRACK%d", randTNAddr)

	prodB, errMarsh := json.Marshal(prodFake)

	if errMarsh != nil {
		p.log.Errorf("can not marshal fake order: %s", errMarsh.Error())
		return
	}

	errPub := p.Sc.Publish(ChannelProd, prodB)

	if errPub != nil {
		p.log.Errorf("can not publish product")
	}

	p.log.Infof("publish new product with track number: %s", prodFake.TrackNumber)
}
