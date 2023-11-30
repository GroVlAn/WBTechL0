package producer

import "github.com/GroVlAn/WBTechL0/internal/apps"

type Producer struct {
	apps.Runner
}

func NewProducer() *Producer {
	return &Producer{}
}

func (p *Producer) Run() {}
