package item

import (
	middleware "ecommerce/rest/middlewares"
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	// Get userID from JWT context (set by AuthenticateJWT middleware)
	userID, ok := r.Context().Value(middleware.UserIDKey).(int64)
	if !ok {
		utils.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	itemId := r.PathValue("id")

	id, err := strconv.ParseInt(itemId, 10, 64)

	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid Item ID")
		return
	}

	err = h.svc.Delete(id, userID)

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	utils.SendData(w, http.StatusOK, "Successfully Deleted the item")
}
