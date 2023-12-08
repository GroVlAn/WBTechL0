package response

import (
	"encoding/json"
	"github.com/GroVlAn/WBTechL0/internal/core"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Response struct {
	Result interface{} `json:"result"`
	Err    interface{} `json:"error"`
}

type InvalidDataResponse struct {
	Message string      `json:"message"`
	Example interface{} `json:"example"`
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

func ErrResponse(w http.ResponseWriter, log *logrus.Logger, err error) {
	var errNf core.NotFoundError
	var errCtc core.CantCreateErr
	var errInvD core.InvalidDataError

	okNf := errors.As(err, &errNf)
	if okNf {
		Resp(w, log, nil, errNf.Message, errNf.Status)
		return
	}

	okCtc := errors.As(err, &errCtc)

	if okCtc {
		Resp(w, log, nil, errCtc.Message, errCtc.Status)
		return
	}

	okInvD := errors.As(err, &errInvD)

	if okInvD {
		errInvSt := InvalidDataResponse{
			Message: errInvD.Message,
			Example: errInvD.Example,
		}
		Resp(w, log, nil, errInvSt, errInvD.Status)
		return
	}

	Resp(w, log, nil, "something wrong", http.StatusInternalServerError)
	return
}
