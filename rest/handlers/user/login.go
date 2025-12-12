package user

import (
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var reqLogin ReqLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request data")
		return
	}

	usr, err := h.svc.Find(reqLogin.Email, reqLogin.Password)

	if usr == nil {
		utils.SendError(w, http.StatusForbidden, "Invalid Credentials")
		return
	}

	accessToken, err := utils.CreateJwt(h.cnf.JwtSecretKey, utils.Payload{
		Sub:       usr.ID,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
	})

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internel Server Error")
		return
	}

	utils.SendData(w, http.StatusCreated, accessToken)
}
