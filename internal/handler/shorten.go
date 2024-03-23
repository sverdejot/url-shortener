package handler

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/sverdejot/url-shortener/internal/service"
)

type ShortenHandler struct {
	Service *service.ShortenService
}

func (s *ShortenHandler) Post(w http.ResponseWriter, r *http.Request) {
	req, err := parseReq(r)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	val := s.Service.Shorten(req.LongUrl)

	w.WriteHeader(http.StatusCreated)

	res := struct {
		ShortenedUrl string `json:"shortened_url"`
	}{
		val,
	}

	json.NewEncoder(w).Encode(res)
}

func (h ShortenHandler) Get(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	decode, err := h.Service.GetUrl(code)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add("Location", decode)
	w.WriteHeader(http.StatusFound)
}

type request struct {
	LongUrl string `json:"long_url"`
}

type PostShortenHandler struct {
	BasePath string
}

func parseReq(r *http.Request) (body request, error error) {
	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		return request{}, err
	}

	json.Unmarshal(b, &body)
	return body, nil
}

func (s *ShortenHandler) InitializeRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /shorten", s.Post)
	mux.HandleFunc("GET /{code}", s.Get)
}
