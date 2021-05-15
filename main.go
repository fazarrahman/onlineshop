package main

import (
	"fmt"

	product_mysql "github.com/fazarrahman/onlineshop/domain/product/repository/mysql"
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

	dbClient, err := db.New()
	if err != nil {
		log.Println(err)
	}
	log.Println("Database is successfully initialized")

	promotionMysqlRepo := promotion_mysql.New(dbClient)
	productMysqlRepo := product_mysql.New(dbClient)
	log.Println("Repositories are successfully initialized")

	router := httprouter.New()
	rest.New(promotionMysqlRepo, productMysqlRepo).Register(router)

	http.Handle("/", router)
	fmt.Println("Connected to port " + lib.GetEnv("APP_PORT"))
	log.Fatal(http.ListenAndServe(":"+lib.GetEnv("APP_PORT"), router))
}

func envInit() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatalln("No .env file found")
	}
}
