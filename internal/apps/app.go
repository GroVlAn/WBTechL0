package apps

import (
	"github.com/GroVlAn/WBTechL0/internal/apps/clientapp"
	dataGeneratorApp "github.com/GroVlAn/WBTechL0/internal/apps/datagenapp"
	ordersApp "github.com/GroVlAn/WBTechL0/internal/apps/orderapp"
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

// CreateClientApp create client application
func (a *Application) CreateClientApp() *clientapp.ClientApp {
	return &clientapp.ClientApp{}
}

type ApplicationsCreator interface {
	CreateOrdersApp() *ordersApp.OrdersApp
	CreateDataGeneratorApp() *dataGeneratorApp.DataGeneratorApp
}

type Runner interface {
	Run(mode string)
}
