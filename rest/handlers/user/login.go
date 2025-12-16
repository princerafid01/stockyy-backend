package user

import (
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var reqLogin UserLoginRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request data")
		return
	}

	// find user by email
	usr, err := h.svc.FindByEmail(reqLogin.Email)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// If user not found or no password hash → invalid credentials
	if usr == nil || usr.PasswordHash == nil {
		utils.SendError(w, http.StatusForbidden, "Invalid Credentials")
		return
	}

	// compare password hash with password
	err = bcrypt.CompareHashAndPassword([]byte(*usr.PasswordHash), []byte(reqLogin.Password))
	if err != nil {
		fmt.Println("error comparing password hash with password ", err)
		utils.SendError(w, http.StatusForbidden, "Invalid Credentials")
		return
	}

	accessToken, err := utils.CreateJwt(h.cnf.JwtSecretKey, utils.Payload{
		Sub:   usr.ID,
		Name:  usr.Name,
		Email: usr.Email,
	})

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internel Server Error")
		return
	}

	resp := UserLoginResponse{
		AccessToken: accessToken,
	}

	utils.SendData(w, http.StatusCreated, resp)
}
