package service

import (
	"errors"
	"fmt"
	"log/slog"
	"math/rand"
)

const charset string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type ShortenService struct {
	baseUrl   string
	shortened map[string]string
}

func NewShortenService(baseUrl string) *ShortenService {
	return &ShortenService{
		baseUrl:   baseUrl,
		shortened: make(map[string]string),
	}
}

func (s *ShortenService) GetUrl(code string) (string, error) {
	slog.Info("trying to find", "code", code)
	v, ok := s.shortened[code]

	if !ok {
		return "", errors.New("code not found")
	}

	return v, nil
}

func (s *ShortenService) Shorten(url string) (code string) {
	code = randomString(5)
	for _, ok := s.shortened[code]; !ok; code = randomString(5) {
		s.shortened[code] = url
		break
	}

	return fmt.Sprintf("%s/%s", s.baseUrl, code)
}

func randomString(n int) string {
	res := make([]byte, 0, n)
	for range n {
		res = append(res, charset[rand.Intn(len(charset))])
	}
	return string(res)
}
