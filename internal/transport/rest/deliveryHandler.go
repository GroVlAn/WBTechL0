package rest

import (
	response "github.com/GroVlAn/WBTechL0/internal/tools/resp"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (hh *HttpHandler) DeliveryHandler() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/", func(r chi.Router) {
		r.Get("/{deliveryID}", hh.Delivery)
		r.Delete("/{deliveryID}", hh.DeleteDelivery)
	})

	return router
}

func (hh *HttpHandler) Delivery(w http.ResponseWriter, req *http.Request) {
	dId := chi.URLParam(req, "deliveryID")

	if dId == "" {
		response.Resp(w, hh.log, nil, "miss delivery id", http.StatusNotFound)
		return
	}

	id, err := strconv.Atoi(dId)

	if err != nil {
		response.Resp(w, hh.log, nil, "id must be integer", http.StatusBadRequest)
		return
	}

	dRepr, errD := hh.dServ.Delivery(int64(id))

	if errD != nil {
		response.ErrResponse(w, hh.log, errD)
		return
	}

	response.Resp(w, hh.log, dRepr, nil, http.StatusOK)
}

func (hh *HttpHandler) DeleteDelivery(w http.ResponseWriter, req *http.Request) {
	dId := chi.URLParam(req, "deliveryID")

	if dId == "" {
		response.Resp(w, hh.log, nil, "miss delivery id", http.StatusNotFound)
		return
	}

	id, err := strconv.Atoi(dId)

	if err != nil {
		response.Resp(w, hh.log, nil, "id must be integer", http.StatusBadRequest)
		return
	}

	delDId, errDelD := hh.dServ.DeleteDelivery(int64(id))

	if errDelD != nil {
		response.ErrResponse(w, hh.log, errDelD)
		return
	}

	delDResp := struct {
		Id int64 `json:"id"`
	}{
		Id: delDId,
	}

	response.Resp(w, hh.log, delDResp, nil, http.StatusOK)
}
