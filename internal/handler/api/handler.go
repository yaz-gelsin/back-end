package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/yaz-gelsin/internal/entities"
)

type YGUseCase interface {
	InsertProduct(ctx context.Context, product entities.Product) (entities.Product, error)
	GetAll(ctx context.Context) ([]entities.Product, error)
}

type YGHandler struct {
	uc YGUseCase
	DB *sqlx.DB
}

func NewHandler(uc YGUseCase, db *sqlx.DB) *YGHandler {
	return &YGHandler{
		uc: uc,
		DB: db,
	}
}

func (h *YGHandler) writeResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	payload, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(payload)
}
