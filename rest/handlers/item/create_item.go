package item

import (
	"ecommerce/domain"
	middleware "ecommerce/rest/middlewares"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqCreateItem struct {
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

func (h *Handler) CreateItem(w http.ResponseWriter, r *http.Request) {
	// STEP 1: Get user_id from context (set by AuthenticateJWT middleware)
	// r.Context() gets the context from the request
	// .Value(middlewares.UserIDKey) gets the value we stored in middleware
	// .(int64) converts it to int64 type (type assertion)
	userID, ok := r.Context().Value(middleware.UserIDKey).(int64)

	// Check if user_id was found in context
	if !ok {
		utils.SendError(w, http.StatusUnauthorized, "Unauthorized: user_id not found")
		return
	}

	// STEP 2: Parse request body
	var req ReqCreateItem
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		utils.SendError(w, http.StatusBadRequest, "Invalid Request Body")
		return
	}

	// STEP 3: Convert []string to domain.StringArray
	imageURLs := domain.StringArray(req.ImageURLs)
	if imageURLs == nil {
		imageURLs = domain.StringArray{}
	}

	// STEP 4: Create item with userID from context
	createdItem, err := h.svc.Create(domain.Item{
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
	utils.SendData(w, http.StatusCreated, createdItem)
}
