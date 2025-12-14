package item

import (
	"ecommerce/domain"
	middleware "ecommerce/rest/middlewares"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ReqUpdateItem struct {
	Name              string   `json:"name"`
	Description       *string  `json:"description,omitempty"`
	Quantity          int      `json:"quantity"`
	LowStockThreshold int      `json:"low_stock_threshold"`
	Price             *float64 `json:"price,omitempty"`
	CostPrice         *float64 `json:"cost_price,omitempty"`
	SKU               *string  `json:"sku,omitempty"`
	Barcode           *string  `json:"barcode,omitempty"`
	CategoryID        *int64   `json:"category_id,omitempty"`
	LocationID        *int64   `json:"location_id,omitempty"`
	ImageURLs         []string `json:"image_urls,omitempty"`
	Notes             *string  `json:"notes,omitempty"`
}

func (h *Handler) UpdateItem(w http.ResponseWriter, r *http.Request) {
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

	var req ReqUpdateItem
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		utils.SendError(w, http.StatusBadRequest, "Invalid Request Body")
		return
	}

	// Convert []string to domain.StringArray
	imageURLs := domain.StringArray(req.ImageURLs)
	if imageURLs == nil {
		imageURLs = domain.StringArray{}
	}

	_, err = h.svc.Update(domain.Item{
		ID:                id,
		UserID:            userID,
		Name:              req.Name,
		Description:       req.Description,
		Quantity:          req.Quantity,
		LowStockThreshold: req.LowStockThreshold,
		Price:             req.Price,
		CostPrice:         req.CostPrice,
		SKU:               req.SKU,
		Barcode:           req.Barcode,
		CategoryID:        req.CategoryID,
		LocationID:        req.LocationID,
		ImageURLs:         imageURLs,
		Notes:             req.Notes,
	})

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	utils.SendData(w, http.StatusOK, "Item Successfully Updated")
}
