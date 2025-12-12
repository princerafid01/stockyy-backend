package product

import (
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	id, err := strconv.Atoi(productId)

	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Please give a valid product ID")
		return
	}

	product, err := h.svc.Get(id)

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	if product == nil {
		utils.SendError(w, http.StatusNotFound, "Product Not Found")
		return
	}

	utils.SendData(w, http.StatusCreated, product)

}
