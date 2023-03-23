package handlers

import (
	"encoding/json"
	"goproduct/config"
	"goproduct/models"
	"net/http"
	"strings"
)

func AddProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := config.Connect()
	defer db.Close()

	db.Create(&product)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func ListProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product

	db := config.Connect()
	defer db.Close()

	sort := r.URL.Query().Get("sort")
	if sort != "" {
		order := "ASC"
		if strings.HasPrefix(sort, "-") {
			order = "DESC"
			sort = strings.TrimPrefix(sort, "-")
		}
		db.Order(sort + " " + order).Find(&products)
	} else {
		db.Order("created_at DESC").Find(&products)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
