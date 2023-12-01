package main

import (
	"github.com/GroVlAn/WBTechL0/internal/apps"
	"github.com/GroVlAn/WBTechL0/scripts/appsArgs"
)

func main() {
	mode := appsArgs.Mode()
	app := apps.NewApplication()
	ordersApplication := app.CreateOrdersApp()
	ordersApplication.Run(mode)
}
