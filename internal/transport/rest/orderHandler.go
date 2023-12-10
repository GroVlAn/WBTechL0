package rest

import (
	"encoding/json"
	"github.com/GroVlAn/WBTechL0/internal/service"
	response "github.com/GroVlAn/WBTechL0/internal/tools/resp"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (hh *HttpHandler) OrderHandler() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/", func(r chi.Router) {
		r.Get("/{orderID}", hh.Order)
		r.Post("/", hh.CreateOrder)
		r.Delete("/{orderID}", hh.DeleteOrder)
	})

	return router
}

func (hh *HttpHandler) CreateOrder(w http.ResponseWriter, req *http.Request) {
	var ordReq service.OrderReq
	err := json.NewDecoder(req.Body).Decode(&ordReq)

	if err != nil {
		invDResp := response.InvalidDataResponse{
			Message: "incorrect data",
			Example: service.ExampleOrderReq,
		}
		response.Resp(w, hh.log, nil, invDResp, http.StatusBadRequest)
		return
	}

	ordUid, errOrd := hh.orServ.CreateOrder(ordReq)

	if errOrd != nil {
		response.ErrResponse(w, hh.log, errOrd)
		return
	}

	ordReps := struct {
		Uid string `json:"order_uid"`
	}{
		Uid: ordUid,
	}

	response.Resp(w, hh.log, ordReps, nil, http.StatusCreated)
}

func (hh *HttpHandler) Order(w http.ResponseWriter, req *http.Request) {
	ordId := chi.URLParam(req, "orderID")

	if ordId == "" {
		response.Resp(w, hh.log, nil, "miss order uid", http.StatusNotFound)
		return
	}

	ordRepr, err := hh.orServ.Order(ordId)

	if err != nil {
		response.ErrResponse(w, hh.log, err)
		return
	}

	response.Resp(w, hh.log, ordRepr, nil, http.StatusOK)
}

func (hh *HttpHandler) DeleteOrder(w http.ResponseWriter, req *http.Request) {
	ordUid := chi.URLParam(req, "orderID")

	if ordUid == "" {
		response.Resp(w, hh.log, nil, "miss order uid", http.StatusNotFound)
		return
	}

	delOrdUid, err := hh.orServ.DeleteOrder(ordUid)

	if err != nil {
		response.ErrResponse(w, hh.log, err)
		return
	}

	ordReps := struct {
		Uid string `json:"order_uid"`
	}{
		Uid: delOrdUid,
	}
	response.Resp(w, hh.log, ordReps, nil, http.StatusOK)
}
