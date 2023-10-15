package middleware

import (
	"context"
	"net/http"

	"github.com/developertom01/library-server/internals/db"
	"github.com/developertom01/library-server/utils"
)

type contextKey struct {
	name string
}

func throwUnAuthorizeError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, `{"message":"UnAuthorized"}`, http.StatusUnauthorized)
	return
}

func AuthenticationMiddleware(db *db.Database) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			h := r.Header["Authorization"]
			if h == nil {
				throwUnAuthorizeError(w, r)
				return
			}
			jwtToken, err := utils.ExtractBearerToken(h[0])
			if err != nil {
				throwUnAuthorizeError(w, r)
				return
			}
			c, err := utils.ValidateToken(jwtToken)
			if err != nil {
				throwUnAuthorizeError(w, r)
				return
			}
			ctx := context.WithValue(r.Context(), "user", c)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
