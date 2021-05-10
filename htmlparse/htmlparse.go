package htmlparse

import(
	"fmt"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

type Parser interface {
	GetTitle(html string)(string, error)
}

type parseImpl struct {}

func New() Parser {
	return &parseImpl{}
}

func (p *parseImpl)GetTitle(html string)(string, error){
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return "", fmt.Errorf("failed to initialize reader; %w", err)
	}
	title := doc.Find("title").Text()
	return title, nil
}