package rest

import (
	"encoding/json"
	"github.com/GroVlAn/WBTechL0/internal/service"
	response "github.com/GroVlAn/WBTechL0/internal/tools/resp"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (hh *HttpHandler) ProductHandler() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/", func(r chi.Router) {
		r.Get("/{productID}", hh.Product)
		r.Get("/all/{trackNumber}", hh.All)
		r.Post("/", hh.CreateProduct)
		r.Delete("/{productID}", hh.DeleteProduct)
	})

	return router
}

func (hh *HttpHandler) CreateProduct(w http.ResponseWriter, req *http.Request) {
	var prodRepr service.ProductRepr
	err := json.NewDecoder(req.Body).Decode(&prodRepr)

	if err != nil {
		invDResp := response.InvalidDataResponse{
			Message: "incorrect data",
			Example: service.ExampleProdReq,
		}
		response.Resp(w, hh.log, nil, invDResp, http.StatusBadRequest)
		return
	}

	id, err := hh.prodServ.CreateProduct(prodRepr)

	if err != nil {
		response.Resp(w, hh.log, nil, err.Error(), http.StatusBadRequest)
		return
	}

	prodResp := struct {
		Id int `json:"id"`
	}{
		Id: id,
	}

	response.Resp(w, hh.log, prodResp, nil, http.StatusCreated)
}

func (hh *HttpHandler) All(w http.ResponseWriter, req *http.Request) {
	trNum := chi.URLParam(req, "trackNumber")

	prReps, err := hh.prodServ.All(trNum)

	if err != nil {
		response.Resp(w, hh.log, nil, err.Error(), http.StatusBadRequest)
	}

	response.Resp(w, hh.log, prReps, nil, http.StatusOK)
}

func (hh *HttpHandler) Product(w http.ResponseWriter, req *http.Request) {
	prodId := chi.URLParam(req, "productID")
	id, err := strconv.Atoi(prodId)

	if err != nil {
		response.Resp(w, hh.log, nil, "id must be integer", http.StatusBadRequest)
		return
	}

	prodRepr, errPr := hh.prodServ.Product(id)

	if errPr != nil {
		response.ErrResponse(w, hh.log, errPr)
		return
	}

	response.Resp(w, hh.log, prodRepr, nil, http.StatusOK)
}

func (hh *HttpHandler) DeleteProduct(w http.ResponseWriter, req *http.Request) {
	prodId := chi.URLParam(req, "productID")
	id, err := strconv.Atoi(prodId)

	if err != nil {
		response.Resp(w, hh.log, nil, "id must be integer", http.StatusBadRequest)
		return
	}

	delProdId, errDel := hh.prodServ.DeleteProduct(id)

	if errDel != nil {
		response.ErrResponse(w, hh.log, errDel)
		return
	}

	delProdResp := struct {
		Id int `json:"id"`
	}{
		Id: delProdId,
	}

	response.Resp(w, hh.log, delProdResp, nil, http.StatusOK)
}
