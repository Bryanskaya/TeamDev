package auth

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"

	"api/teamdev/utils"
)

func LoginFromCookie(r *http.Request) (string, error) {
	c, err := r.Cookie("token")
	if err != nil {
		return "", err
	}

	tokenStr := c.Value
	tk := new(Token)

	_, err = jwt.ParseWithClaims(tokenStr, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.Config.TokenPassword), nil
	})
	return tk.Login, err
}
