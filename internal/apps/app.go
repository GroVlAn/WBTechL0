package apps

import (
	"github.com/GroVlAn/WBTechL0/internal/apps/producer"
	"github.com/GroVlAn/WBTechL0/internal/apps/subscriber"
)

type Runner interface {
	Run()
}

var applications map[string]Runner = map[string]Runner{
	"producer":   producer.NewProducer(),
	"subscriber": subscriber.NewSubscriber(),
}

func createApp(nameApp string) Runner {
	return applications[nameApp]
}
