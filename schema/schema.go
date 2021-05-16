package schema

import (
	"github.com/fazarrahman/onlineshop/service"

	"github.com/graphql-go/graphql"
)

// define cart type
var cartType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Cart",
	Fields: graphql.Fields{
		"sku": &graphql.Field{
			Type: graphql.String,
		},
		"totalAmount": &graphql.Field{
			Type: graphql.Float,
		},
	},
})

// Cart ...
type Cart struct {
	TotalAmount float64 `json:"totalAmount"`
}

// define product type
var productType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Product",
	Fields: graphql.Fields{
		"sku": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"price": &graphql.Field{
			Type: graphql.Float,
		},
		"qty": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

// Product ...
type Product struct {
	SKU   string  `json:"sku"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Qty   int64   `json:"qty"`
}

// root mutation
var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		// this is graphql to checkout cart
		// the parameter are list of scanned item's SKU
		/*
			curl example :
			curl --location --request POST 'http://localhost:4000/graphql' \
			--header 'Content-Type: application/json' \
			--data-raw '{
				"query":"mutation{cartCheckout(sku:[\"43N23P\",\"234234\"]){totalAmount}}"
			}'
		*/
		"cartCheckout": &graphql.Field{
			Type:        cartType, // the return type for this field
			Description: "Checkout the cart",
			Args: graphql.FieldConfigArgument{
				"sku": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				sku, _ := params.Args["sku"].([]interface{})

				skuList := make([]string, len(sku))
				for i, v := range sku {
					skuList[i] = v.(string)
				}

				// get the service depedency injection
				varVal := params.Info.RootValue.(map[string]interface{})
				// call cart checkout service
				price, err := varVal["service"].(service.Service).CartCheckout(params.Context, skuList)
				if err != nil {
					return nil, err
				}

				return Cart{
					TotalAmount: *price,
				}, nil
			},
		},
	},
})

// root query
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		// this is graphql to get all active product
		/*
				curl example :
			   curl --location -g --request POST 'http://localhost:4000/graphql?query={productList{sku}}' \
				--header 'Content-Type: application/json' \
				--data-raw '{
					"query":"{productList{sku, name, price, qty}}"
				}'
		*/
		"productList": &graphql.Field{
			Type:        graphql.NewList(productType),
			Description: "List of products",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// get the service depedency injection
				varVal := p.Info.RootValue.(map[string]interface{})
				// call get all product service
				productList, err := varVal["service"].(service.Service).GetAllProducts(p.Context)
				if err != nil {
					return nil, err
				}

				var products []Product
				for _, p := range productList {
					products = append(products, Product{
						SKU:   p.SKU,
						Name:  p.Name,
						Price: p.Price,
						Qty:   p.Qty,
					})
				}

				return products, nil
			},
		},
	},
})

// define schema, with rootQuery and rootMutation
var CartSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})
