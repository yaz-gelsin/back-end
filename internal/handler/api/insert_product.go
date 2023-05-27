package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/yaz-gelsin/internal/entities"
)

func (h *YGHandler) InsertProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var product entities.Product
	err = json.Unmarshal(body, &product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(product)

	product, err = h.uc.InsertProduct(ctx, product)
	if err != nil {
		http.Error(w, "An error occurred: "+err.Error(), http.StatusInternalServerError)
		return
	}

	h.writeResponse(w, http.StatusOK, product)
}
