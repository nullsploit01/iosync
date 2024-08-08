package middlewares

import (
	"context"
	"iosync/pkg/constants"
	"iosync/pkg/utils"
	"net/http"
)

func SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionId, err := utils.GetCookieValueFromRequest(r, constants.SessionIDCookieKey)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), constants.SessionIDCookieKey, sessionId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
