package apps

import (
	dataGeneratorApp "github.com/GroVlAn/WBTechL0/internal/apps/dataGenerator"
	ordersApp "github.com/GroVlAn/WBTechL0/internal/apps/orders"
)

type Application struct {
	ApplicationsCreator
}

func NewApplication() *Application {
	return &Application{}
}

// CreateOrdersApp create order application
func (a *Application) CreateOrdersApp() *ordersApp.OrdersApp {
	return &ordersApp.OrdersApp{}
}

// CreateDataGeneratorApp create data generator application
func (a *Application) CreateDataGeneratorApp() *dataGeneratorApp.DataGeneratorApp {
	return &dataGeneratorApp.DataGeneratorApp{}
}

type ApplicationsCreator interface {
	CreateOrdersApp() *ordersApp.OrdersApp
	CreateDataGeneratorApp() *dataGeneratorApp.DataGeneratorApp
}

type Runner interface {
	Run(mode string)
}
