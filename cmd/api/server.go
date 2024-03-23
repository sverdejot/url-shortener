package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sverdejot/url-shortener/internal/handler"
	"github.com/sverdejot/url-shortener/internal/middleware"
	"github.com/sverdejot/url-shortener/internal/service"
)

const (
	BASE_URL = "http://localhost"
	PORT     = "8080"
)

func Run() {
	mux := http.NewServeMux()

	shortener := service.NewShortenService(fmt.Sprintf("%s:%s", BASE_URL, PORT))

	handler := handler.ShortenHandler{
		Service: shortener,
	}

	handler.InitializeRoutes(mux)

	srv := middleware.Apply(mux, middleware.Timer, middleware.CorrelationId)

	log.Println("starting server")
	if err := http.ListenAndServe(fmt.Sprintf(":%s", PORT), srv); err != nil {
		log.Fatal(err)
	}
}
