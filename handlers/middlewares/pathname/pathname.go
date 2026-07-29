package pathname

import (
	"context"
	"net/http"
)

const PathNameKey = "path-name"

func Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), PathNameKey, r.URL.Path)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
