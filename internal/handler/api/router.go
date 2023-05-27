package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func (h *YGHandler) InitRouter(r *mux.Router, db *sqlx.DB) http.Handler {

	h.DB = db

	r.HandleFunc("/", h.InsertProduct).Methods(http.MethodPost)

	return r
}
