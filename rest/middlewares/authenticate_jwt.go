package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"ecommerce/utils"
	"net/http"
	"strings"
)

func (m *Middlewares) AuthenticateJWT(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get jwt from header
		header := r.Header.Get("Authorization")

		if header == "" {
			http.Error(w, "UnAuthorized", http.StatusUnauthorized)
			return
		}

		headerArr := strings.Split(header, " ")
		if len(headerArr) != 2 {
			http.Error(w, "UnAuthorized", http.StatusUnauthorized)
			return
		}

		accessToken := headerArr[1]

		tokenParts := strings.Split(accessToken, ".")

		if len(tokenParts) != 3 {
			http.Error(w, "UnAuthorized", http.StatusUnauthorized)
			return
		}

		jwtHeader := tokenParts[0]  // base64
		jwtPayload := tokenParts[1] // base64
		signature := tokenParts[2]  // base64

		message := jwtHeader + "." + jwtPayload
		byteArrSecret := []byte(m.cnf.JwtSecretKey)
		byteArrMessage := []byte(message)

		h := hmac.New(sha256.New, byteArrSecret)
		h.Write(byteArrMessage)
		hash := h.Sum(nil)
		newGeneratedSignature := utils.Base64UrlEncode(hash)

		if signature != newGeneratedSignature {
			http.Error(w, "Unauthorized!", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
