package handler

import (
	"database/sql"
	"net/http"

	"github.com/yaz-gelsin/internal/handler/api"
)

func InitRouter(db *sql.DB) http.Handler {
	// Router'ı oluşturun
	router := http.NewServeMux()

	// Handler fonksiyonlarını router'a ekleme
	router.HandleFunc("/", api.GetAll)

	// db değişkenini kullanmak için router'ı döndürün
	return router
}
