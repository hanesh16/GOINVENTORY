package api

import (
	"database/sql"
	"net/http"

	"github.com/IBM/sarama"
)

func RegisterRoutes(mux *http.ServeMux, db *sql.DB, producer sarama.SyncProducer) {
	handler := NewHandler(db, producer)

	// POST /products
	// GET  /products
	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handler.AddProductHandler(w, r)
			return
		}

		if r.Method == http.MethodGet {
			handler.ListProductsHandler(w, r)
			return
		}

		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})
}
