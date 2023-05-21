package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/yaz-gelsin/internal/entities"
)

func (h *YGHandler) InsertProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body, err := ioutil.ReadAll(r.Body)
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

	product, err = h.uc.InsertProduct(ctx, product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.writeResponse(w, http.StatusOK, product)
}
