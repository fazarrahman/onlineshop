package service

import (
	"context"
	"errors"
	"sort"

	promotionEntity "github.com/fazarrahman/onlineshop/domain/promotion/entity"
)

// CartCheckout ...
// Calculate price for the submitted cart
func (s *Svc) CartCheckout(ctx context.Context, skuList []string) (*float64, error) {
	if len(skuList) == 0 {
		return nil, errors.New("SKU list is required")
	}

	// get all active products to get the price
	products, err := s.productRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	// get all active promotions
	promotions, err := s.promotionRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	sort.Strings(skuList)

	prodQtyMap := make(map[string]int64)
	prodPrice := make(map[string]float64)
	prodPromotion := make(map[string]*promotionEntity.Promotion)
	var skuDist []string
	for _, sku := range skuList {
		if prodQtyMap[sku] == 0 {
			// distinct the duplicated sku in request
			skuDist = append(skuDist, sku)

			// set product price to map
			for _, p := range products {
				if sku == p.SKU {
					prodPrice[sku] = p.Price
					break
				}
			}

			// set promotions to map
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
		// if item has promotion and meet minimum required quantity
		if prodPromotion[sku] != nil && prodQtyMap[sku] >= prodPromotion[sku].MinRequiredQty {
			// if item promotion is free other item
			if prodPromotion[sku].FreeItemSKU != nil {
				prPrice = prodPrice[sku]
				// the product price is current product price - free item price
				if prodQtyMap[*prodPromotion[sku].FreeItemSKU] > 0 {
					prPrice -= prodPrice[*prodPromotion[sku].FreeItemSKU] * float64(prodQtyMap[sku])
				}
			} else if prodPromotion[sku].NewPriceQty != nil { // if item promotion is the price with less quantity
				mul := prodQtyMap[sku] / prodPromotion[sku].MinRequiredQty
				prPrice = prodPrice[sku] * float64((prodQtyMap[sku] - mul))
			} else if prodPromotion[sku].DiscountInPercent != nil { // if item promotion is discount
				prPrice = prodPrice[sku] * float64(prodQtyMap[sku]) * (1 - float64(*prodPromotion[sku].DiscountInPercent)/100)
			} else {
				// if the item quantity is less than minimum required quantity required for the promotion. No promotion applied
				prPrice = prodPrice[sku] * float64(prodQtyMap[sku])
			}
		} else {
			// if the item has no promotion. No promotion applied
			prPrice = prodPrice[sku] * float64(prodQtyMap[sku])
		}
		price += prPrice
	}

	return &price, nil
}
