package middleware

import (
	"context"
	"net/http"

	"github.com/developertom01/library-server/internals/db"
	"github.com/developertom01/library-server/utils"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func AuthenticationMiddleware(db *db.Database) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			h := r.Header["Authorization"]
			if h == nil {
				next.ServeHTTP(w, r)
				return
			}
			jwtToken := utils.ExtractBearerToken(h[0])
			c, err := utils.ValidateToken(jwtToken)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			ctx := context.WithValue(r.Context(), userCtxKey, c)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
