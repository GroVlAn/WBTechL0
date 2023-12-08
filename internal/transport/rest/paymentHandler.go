package rest

import (
	response "github.com/GroVlAn/WBTechL0/internal/tools/resp"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (hh *HttpHandler) PaymentHandler() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/", func(r chi.Router) {
		r.Get("/{paymentID}", hh.Payment)
		r.Delete("/{paymentID}", hh.DeletePayment)
	})

	return router
}

func (hh *HttpHandler) Payment(w http.ResponseWriter, req *http.Request) {
	pmtId := chi.URLParam(req, "paymentId")

	pmtResp, errPmt := hh.pmtServ.Payment(pmtId)

	if errPmt != nil {
		response.ErrResponse(w, hh.log, errPmt)
		return
	}

	response.Resp(w, hh.log, pmtResp, nil, http.StatusOK)
}

func (hh *HttpHandler) DeletePayment(w http.ResponseWriter, req *http.Request) {
	pmtId := chi.URLParam(req, "paymentId")

	delPmtTran, errDelPmt := hh.pmtServ.DeletePayment(pmtId)

	if errDelPmt != nil {
		response.ErrResponse(w, hh.log, errDelPmt)
		return
	}

	pmtResp := struct {
		Transaction string `json:"transaction"`
	}{
		Transaction: delPmtTran,
	}
	response.Resp(w, hh.log, pmtResp, nil, http.StatusOK)
}
