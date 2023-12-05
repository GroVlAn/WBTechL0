package response

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Response struct {
	Result interface{} `json:"result"`
	Err    interface{} `json:"error"`
}

func Resp(w http.ResponseWriter, log *logrus.Logger, result interface{}, errRep interface{}, status ...int) {
	var st int
	for _, s := range status {
		st = s
		break
	}
	if st == 0 {
		st = 200
	}

	resp := Response{
		Result: result,
		Err:    errRep,
	}

	respB, err := json.Marshal(resp)

	if err != nil {
		log.Errorf(errors.Wrap(err, "error can not create response").Error())
		return
	}

	if errRep != nil {
		http.Error(w, string(respB), st)
		return
	}

	w.WriteHeader(st)
	if _, err := w.Write(respB); err != nil {
		log.Errorf("Can not write response: %s", err.Error())
		return
	}
}
