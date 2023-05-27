package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func (h *YGHandler) InitRouter(r *mux.Router, db *sqlx.DB) http.Handler {

	h.DB = db

	r.HandleFunc("/insert-product", h.InsertProduct).Methods(http.MethodPost)
	r.HandleFunc("/products", h.GetAllProducts).Methods(http.MethodGet)

	return r
}
