package item

import (
	middleware "ecommerce/rest/middlewares"
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func (h *Handler) GetItem(w http.ResponseWriter, r *http.Request) {
	// Get userID from JWT context (set by AuthenticateJWT middleware)
	userID, ok := r.Context().Value(middleware.UserIDKey).(int64)
	if !ok {
		utils.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	itemId := r.PathValue("id")

	id, err := strconv.ParseInt(itemId, 10, 64)

	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Please give a valid item ID")
		return
	}

	item, err := h.svc.Get(id, userID)

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	if item == nil {
		utils.SendError(w, http.StatusNotFound, "Item Not Found")
		return
	}

	utils.SendData(w, http.StatusOK, item)
}
