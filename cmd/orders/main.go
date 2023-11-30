package main

import (
	"github.com/GroVlAn/WBTechL0/internal/apps"
)

func main() {
	app := apps.NewApplication()
	ordersApplication := app.CreateOrdersApp()
	ordersApplication.Run()
}
