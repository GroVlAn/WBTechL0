package appargs

import "os"

type Modes = map[string]string

var modes Modes = Modes{
	"--prod": "prod",
	"--dev":  "dev",
}

func Mode() string {
	var mode string

	args := os.Args

	if mode = modes[args[1]]; mode == "" {
		return modes["--dev"]
	}

	return mode
}
