package ordersApp

type OrdersApp struct {
	Runner
}

func (p *OrdersApp) Run() {}

type Runner interface {
	Run()
}
