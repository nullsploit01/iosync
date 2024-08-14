package middlewares

import (
	"context"
	"iosync/internal/services"
	"iosync/pkg/constants"
	"iosync/pkg/utils"
	"net/http"
	"time"
)

func SessionMiddleware(sessionService *services.SessionService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sessionId, err := utils.GetCookieValueFromRequest(r, string(constants.SessionIDKey))

			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			session, err := sessionService.VerifySession(context.Background(), sessionId)

			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			if session.ExpiresAt.Before(time.Now()) {
				http.Error(w, "Session expired", http.StatusUnauthorized)
				return
			}

			if err = sessionService.RefreshSessionExpiry(r.Context(), session); err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), constants.UsernameKey, session.Username)
			ctx = context.WithValue(ctx, constants.SessionIDKey, session.SessionID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
