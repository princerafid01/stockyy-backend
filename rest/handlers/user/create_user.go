package user

import (
	"ecommerce/domain"
	"ecommerce/utils"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req UserCreateRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request data")
		return
	}

	usr, err := h.svc.FindByEmail(req.Email)

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	if usr != nil {
		utils.SendError(w, http.StatusBadRequest, "User already exists")
		return
	}

	var hashedPassword *string

	if req.Password != nil && *req.Password != "" {
		hased, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
		if err != nil {
			utils.SendError(w, http.StatusInternalServerError, "Failed to process password")
			return
		}
		hashStr := string(hased)
		hashedPassword = &hashStr
	}

	user, err := h.svc.Create(domain.User{
		Email:        req.Email,
		PasswordHash: hashedPassword,
		GoogleID:     req.GoogleID,
		Name:         req.Name,
		AvatarURL:    req.AvatarURL,
	})

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	resp := UserCreateResponse{
		ID:        user.ID,
		Email:     user.Email,
		GoogleID:  user.GoogleID,
		Name:      user.Name,
		AvatarURL: user.AvatarURL,
	}

	utils.SendData(w, http.StatusCreated, resp)
}
