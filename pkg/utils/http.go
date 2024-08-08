package utils

import "net/http"

func GetCookieValueFromRequest(r *http.Request, key string) (string, error) {
	cookie, err := r.Cookie(key)
	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}
