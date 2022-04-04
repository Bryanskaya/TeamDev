package responses

import (
	"encoding/json"
	"net/http"
)

func BadRequest(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(msg)
}

func TokenIsMissed(w http.ResponseWriter) {
	msg := "Missing auth token"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(msg)
}

func JwtAccessDenied(w http.ResponseWriter) {
	msg := "jwt-token is not valid"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusForbidden)
	json.NewEncoder(w).Encode(msg)
}

func TokenExpired(w http.ResponseWriter) {
	msg := "jwt-token expired"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(msg)
}
