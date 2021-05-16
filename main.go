package main

import (
	"encoding/json"
	"fmt"

	"github.com/fazarrahman/onlineshop/service"

	"github.com/fazarrahman/onlineshop/schema"

	product_mysql "github.com/fazarrahman/onlineshop/domain/product/repository/mysql"
	"github.com/graphql-go/graphql"
	"github.com/joho/godotenv"

	"github.com/fazarrahman/onlineshop/rest"

	promotion_mysql "github.com/fazarrahman/onlineshop/domain/promotion/repository/mysql"

	db "github.com/fazarrahman/onlineshop/config/mysql"
	"github.com/fazarrahman/onlineshop/lib"

	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	envInit()

	// Initialize mysql database
	dbClient, err := db.New()
	if err != nil {
		log.Println(err)
	}
	log.Println("Database is successfully initialized")

	// initialize repository layer
	promotionMysqlRepo := promotion_mysql.New(dbClient)
	productMysqlRepo := product_mysql.New(dbClient)
	log.Println("Repositories are successfully initialized")

	// initialize service layer
	svc := service.New(promotionMysqlRepo, productMysqlRepo)
	log.Println("Services are successfully initialized")

	router := httprouter.New()
	// initialize rest api
	rest.New(svc).Register(router)

	// initialize graphql
	router.POST("/graphql", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var p postData
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			log.Println(err)
			w.WriteHeader(400)
			return
		}

		var rootObj = make(map[string]interface{})
		// send service layer depedency injection to graphql resolver
		rootObj["service"] = svc
		result := graphql.Do(graphql.Params{
			Context:       r.Context(),
			Schema:        schema.CartSchema,
			RequestString: p.Query,
			RootObject:    rootObj,
		})
		if err := json.NewEncoder(w).Encode(result); err != nil {
			fmt.Printf("could not write result to response: %s", err)
		}
	})

	http.Handle("/", router)

	fmt.Println("Connected to port " + lib.GetEnv("APP_PORT"))
	log.Fatal(http.ListenAndServe(":"+lib.GetEnv("APP_PORT"), router))
}

type postData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

func envInit() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatalln("No .env file found")
	}
}
