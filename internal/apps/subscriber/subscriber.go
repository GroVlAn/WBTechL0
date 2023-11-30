package subscriber

import "github.com/GroVlAn/WBTechL0/internal/apps"

type Subscriber struct {
	apps.Runner
}

func NewSubscriber() *Subscriber {
	return &Subscriber{}
}

func (s *Subscriber) Run() {

}
