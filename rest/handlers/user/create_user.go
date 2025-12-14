package user

import (
	"ecommerce/domain"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

type ReqCreateUser struct {
	Email        string  `json:"email"`
	PasswordHash *string `json:"password_hash,omitempty"`
	GoogleID     *string `json:"google_id,omitempty"`
	Name         string  `json:"name"`
	AvatarURL    *string `json:"avatar_url,omitempty"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req ReqCreateUser
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request data")
		return
	}

	user, err := h.svc.Create(domain.User{
		Email:        req.Email,
		PasswordHash: req.PasswordHash,
		GoogleID:     req.GoogleID,
		Name:         req.Name,
		AvatarURL:    req.AvatarURL,
	})

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	utils.SendData(w, http.StatusCreated, user)
}
