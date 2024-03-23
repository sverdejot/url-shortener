package middleware

import (
	"net/http"
)

type middleware func(h http.Handler) http.Handler

func Apply(h http.Handler, middlewares ...middleware) (applied http.Handler) {
	for _, m := range middlewares {
		h = m(h)
	}
	return h
}
