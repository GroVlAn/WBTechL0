package main

import (
	"github.com/GroVlAn/WBTechL0/internal/apps"
	"github.com/GroVlAn/WBTechL0/scripts/appargs"
)

func main() {
	mode := appargs.Mode()
	app := apps.NewApplication()
	clientApplication := app.CreateClientApp()
	clientApplication.Run(mode)
}
