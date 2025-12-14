package item

import (
	middleware "ecommerce/rest/middlewares"
	"ecommerce/utils"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) GetItems(w http.ResponseWriter, r *http.Request) {
	// Get userID from JWT context (set by AuthenticateJWT middleware)
	userID, ok := r.Context().Value(middleware.UserIDKey).(int64)
	if !ok {
		utils.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	reqQuery := r.URL.Query()

	pageAsString := reqQuery.Get("page")
	limitAsString := reqQuery.Get("limit")

	page, _ := strconv.ParseInt(pageAsString, 10, 64)
	limit, _ := strconv.ParseInt(limitAsString, 10, 64)

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	itemList, err := h.svc.List(userID, page, limit)
	if err != nil {
		fmt.Println(err)
		utils.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	itemCount, err := h.svc.Count(userID)
	if err != nil {
		fmt.Println(err)
		utils.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	utils.SendWithPagination(w, itemList, page, limit, itemCount)
}
