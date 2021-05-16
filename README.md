# onlineshop

## Please follow these steps to run the backend api

### A. Prepare mysql database
1. Please execute script in sqlquery.sql
2. Please add the connection to .env files, for example :
APP_PORT=4000
DB_SERVER=127.0.0.1:3306
DB_DATABASE_NAME=onlineshopdb
DB_USERNAME=root
DB_PASSWORD=thepassword

### B. Run with this command : go run main.go

### C. Curl to get the product list :
curl --location -g --request POST 'http://localhost:4000/graphql?query={productList{sku}}' \
--header 'Content-Type: application/json' \
--data-raw '{
    "query":"{productList{sku, name, price, qty}}"
}'

### D. Curl to calculate the cart price :
curl --location --request POST 'http://localhost:4000/cart/checkout' \
--header 'Content-Type: application/json' \
--data-raw '{
    "sku": ["234234","234234","120P90","43N23P"]
}'
