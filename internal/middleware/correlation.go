package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type corrId string

const key corrId = "corr_id"

func CorrelationId(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		corr_id := uuid.NewString()
		ctx = context.WithValue(ctx, key, corr_id)

		r = r.WithContext(ctx)
		f.ServeHTTP(w, r)
	})
}
