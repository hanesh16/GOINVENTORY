package api

import (
	"database/sql"
	"encoding/json"
	"goinventory/model"
	"net/http"

	"github.com/IBM/sarama"
)

type Handler struct {
	biz IBizlogic
}

func NewHandler(db *sql.DB, producer sarama.SyncProducer) Handler {
	bizlogic := NewBizlogic(db, producer)
	return Handler{
		biz: bizlogic,
	}
}

// POST /products
func (handler Handler) AddProductHandler(w http.ResponseWriter, r *http.Request) {
	var product model.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	createdProduct, err := handler.biz.AddProduct(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdProduct)
}

// GET /products
func (handler Handler) ListProductsHandler(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")

	products, err := handler.biz.ListProducts(search)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
