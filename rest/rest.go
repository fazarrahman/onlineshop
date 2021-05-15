package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"

	product_repository "github.com/fazarrahman/onlineshop/domain/product/repository"
	"github.com/julienschmidt/httprouter"

	promotionEntity "github.com/fazarrahman/onlineshop/domain/promotion/entity"

	promotion_repository "github.com/fazarrahman/onlineshop/domain/promotion/repository"
)

// Rest ...
type Rest struct {
	promotionRepository promotion_repository.Repository
	productRepository   product_repository.Repository
}

// New ...
func New(promotionRepository promotion_repository.Repository, productRepository product_repository.Repository) *Rest {
	return &Rest{promotionRepository: promotionRepository, productRepository: productRepository}
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

	/*
		if strings.Trim(news.Author, " ") == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Authors cannot be blank")
			return
		} else if strings.Trim(news.Body, " ") == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Body cannot be blank")
			return
		}*/

	products, err := rest.productRepository.GetAll(r.Context())
	promotions, err := rest.promotionRepository.GetAll(r.Context())
	sort.Strings(cart.SKUList)

	prodQtyMap := make(map[string]int64)
	prodPrice := make(map[string]float64)
	prodPromotion := make(map[string]*promotionEntity.Promotion)
	var skuDist []string
	for _, sku := range cart.SKUList {
		if prodQtyMap[sku] == 0 {
			skuDist = append(skuDist, sku)
			for _, p := range products {
				if sku == p.SKU {
					prodPrice[sku] = p.Price
					break
				}
			}
			for _, pr := range promotions {
				if pr.ApplyToSKU == sku {
					prodPromotion[sku] = pr
					break
				}
			}
		}
		prodQtyMap[sku]++
	}

	var price float64 = 0
	for _, sku := range skuDist {
		var prPrice float64
		if prodPromotion[sku] != nil && prodQtyMap[sku] >= prodPromotion[sku].MinRequiredQty {
			if prodPromotion[sku].FreeItemSKU != nil {
				prPrice = prodPrice[sku]
				if prodQtyMap[*prodPromotion[sku].FreeItemSKU] > 0 {
					prPrice -= prodPrice[*prodPromotion[sku].FreeItemSKU] * float64(prodQtyMap[sku])
				}
			} else if prodPromotion[sku].NewPriceQty != nil {
				mul := prodQtyMap[sku] / prodPromotion[sku].MinRequiredQty
				prPrice = prodPrice[sku] * float64((prodQtyMap[sku] - mul))
			} else if prodPromotion[sku].DiscountInPercent != nil {
				prPrice = prodPrice[sku] * float64(prodQtyMap[sku]) * (1 - float64(*prodPromotion[sku].DiscountInPercent)/100)
			} else {
				prPrice = prodPrice[sku] * float64(prodQtyMap[sku])
			}
		} else {
			prPrice = prodPrice[sku] * float64(prodQtyMap[sku])
		}
		price += prPrice
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(price)
}
