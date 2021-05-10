package html2md

import (
	"fmt"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

type Converter interface {
	Convert(html string) (string, error)
}

type converterImpl struct{}

func New() Converter {
	return &converterImpl{}
}

func (c *converterImpl) Convert(html string) (string, error) {
	converter := md.NewConverter("", true, nil)
	m, err := converter.ConvertString(html)
	if err != nil {
		return "", fmt.Errorf("failed to convert HTML to markdown; %w", err)
	}
	return m, nil
}
