package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fazarrahman/onlineshop/service"

	"github.com/julienschmidt/httprouter"
)

type Rest struct {
	svc service.Service
}

func New(svc service.Service) *Rest {
	return &Rest{svc: svc}
}

// Register ...
func (r *Rest) Register(router *httprouter.Router) {
	router.POST("/cart/checkout", r.handleCheckoutCart)
}

// Cart ...
type Cart struct {
	SKUList []string `json:"sku"`
}

// handleCheckoutCart ...
func (rest *Rest) handleCheckoutCart(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var cart Cart
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	err = json.Unmarshal(reqBody, &cart)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error unmarshal")
		return
	}

	price, err := rest.svc.CartCheckout(r.Context(), cart.SKUList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return

	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(price)
}
