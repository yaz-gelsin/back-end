package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func (h *YGHandler) InitRouter(r *mux.Router, db *sqlx.DB) http.Handler {

	h.DB = db

	// Handler fonksiyonlarını router'a ekleme
	r.HandleFunc("/", h.InsertProduct).Methods(http.MethodPost)

	// db değişkenini kullanmak için router'ı döndürün
	return r
}
