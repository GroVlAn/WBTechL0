package writeresp

import "github.com/sirupsen/logrus"

type FnWrite func([]byte) (int, error)

func Write(log *logrus.Logger, b []byte, errText string, fn FnWrite) {
	if _, err := fn(b); err != nil {
		log.Errorf("%s: %s", errText, err.Error())
	}
}
