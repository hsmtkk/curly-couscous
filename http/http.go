package http

import (
	"fmt"
	"io"
	"net/http"
)

type Getter interface {
	Get(url string) (string, error)
}

type getterImpl struct {
}

func New() Getter {
	return &getterImpl{}
}

func (g *getterImpl) Get(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("HTTP GET failed; %w", err)
	}
	defer resp.Body.Close()
	code := resp.StatusCode
	if code < 200 || code >= 300 {
		return "", fmt.Errorf("non 2xx HTTP status code; %d", code)
	}
	html, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response; %w", err)
	}
	return string(html), nil
}
