package writeresp

import "github.com/sirupsen/logrus"

type fnWrite func([]byte) (int, error)

func Write(log *logrus.Logger, b []byte, errText string, fn fnWrite) {
	if _, err := fn([]byte(b)); err != nil {
		log.Errorf("%s: %s", errText, err.Error())
	}
}
