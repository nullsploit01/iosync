package middlewares

import (
	"context"
	"iosync/pkg/utils"
	"net/http"
)

func SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionId, err := utils.GetCookieValueFromRequest(r, "session_id")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "session_id", sessionId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
