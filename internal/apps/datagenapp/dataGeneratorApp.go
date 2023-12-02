package datagenapp

type DataGeneratorApp struct {
	Runner
}

func (s *DataGeneratorApp) Run(mode string) {

}

type Runner interface {
	Run(mode string)
}
