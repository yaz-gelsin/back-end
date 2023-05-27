package api

import (
	"net/http"
)

func (h *YGHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	product, err := h.uc.GetAll(ctx)
	if err != nil {
		http.Error(w, "An error occurred: "+err.Error(), http.StatusInternalServerError)
		return
	}

	h.writeResponse(w, http.StatusOK, product)
}
