package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"example/database"

	_ "github.com/lib/pq"
)

type Product struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Instock bool   `json:"instock"`
	Price   int    `json:"price"`
}

var productsList []Product

func init() {
	productsJSON := `[
		{
			"id": 1,
			"name": "Suka",
			"price": 13000,
			"instock": true
		},
		{
			"id": 2,
			"name": "Bliad",
			"price": 10000,
			"instock": false
		},
		{
			"id": 3,
			"name": "Bliad",
			"price": 10000,
			"instock": false
		}
	]`

	err := json.Unmarshal([]byte(productsJSON), &productsList)
	if err != nil {
		log.Fatal(err)
	}
}

func getNextID() int {
	highestID := -1
	for _, product := range productsList {
		if highestID < product.ID {
			highestID = product.ID
		}
	}
	return highestID + 1
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	results, err := database.DbConn.Query(`SELECT price,
	name,
	id,
	instock
	FROM products`)
	if err != nil {
		log.Fatal(err)
	}
	defer results.Close()

	switch r.Method {
	case http.MethodGet:
		products := make([]Product, 0)
		for results.Next() {
			var product Product
			results.Scan(
				&product.Price,
				&product.Name,
				&product.ID,
				&product.Instock,
			)
			products = append(products, product)
		}

		productsJson, err := json.Marshal(products)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(productsJson)
	case http.MethodPost:
		var newProduct Product
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		err = json.Unmarshal(bodyBytes, &newProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		if newProduct.ID != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		newProduct.ID = getNextID()
		productsList = append(productsList, newProduct)
		w.WriteHeader(http.StatusCreated)
		return
	}
}

type fooHandler struct {
	Message string
}

func main() {
	database.SetupDatabase()
	http.HandleFunc("/products", productsHandler)
	http.ListenAndServe(":8080", nil)
}
