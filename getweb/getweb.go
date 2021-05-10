package getweb

import (
	"fmt"
	"io"
	"net/http"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

type WebGetter interface {
	GetWeb(url string) (string, error)
}

type getterImpl struct {
}

func New() WebGetter {
	return &getterImpl{}
}

func (g *getterImpl) GetWeb(url string) (string, error) {
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
	return convert(string(html))
}

func convert(html string) (string, error) {
	converter := md.NewConverter("", true, nil)
	m, err := converter.ConvertString(html)
	if err != nil {
		return "", fmt.Errorf("failed to convert HTML to markdown; %w", err)
	}
	return m, nil
}
