package middleware

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
	"strings"
)

// Define custom type for context keys to avoid collisions
// This prevents different packages from accidentally using the same string key
type contextKey string

// Define the specific key we'll use for user_id
// This is exported (starts with capital) so handlers can use it
const UserIDKey contextKey = "user_id"

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

		payloadBytes, err := utils.Base64UrlDecode(jwtPayload)

		if err != nil {
			http.Error(w, "UnAuthorized", http.StatusUnauthorized)
			return
		}

		var payload utils.Payload

		err = json.Unmarshal(payloadBytes, &payload)
		if err != nil {
			http.Error(w, "Unauthorized!", http.StatusUnauthorized)
			return
		}

		// Add user_id to request context
		// Create new context with user_id value using our custom key type
		ctx := context.WithValue(r.Context(), UserIDKey, payload.Sub)

		// Pass the new context to next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
