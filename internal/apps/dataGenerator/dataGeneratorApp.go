package dataGeneratorApp

type DataGeneratorApp struct {
	Runner
}

func (s *DataGeneratorApp) Run() {

}

type Runner interface {
	Run()
}
