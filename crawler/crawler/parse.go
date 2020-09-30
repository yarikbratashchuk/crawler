package crawler

import (
	"fmt"
	"io"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

// Parser abstracts parsing data and extracting links
type Parser interface {
	Links(io.Reader) ([]string, error)
}

// naiveParser is a proof of concept for this test project :)
type naiveParser struct {
	baseURL *url.URL
}

func NewNaiveParser(baseLink string) (Parser, error) {
	urlParsed, err := ParseLink(baseLink)
	if err != nil {
		return nil, fmt.Errorf("incorrect domain: %s", baseLink)
	}
	if !validLink(urlParsed) {
		return nil, fmt.Errorf("request to %s was unsuccessfull", urlParsed)
	}

	return naiveParser{urlParsed}, nil
}

func (p naiveParser) Links(data io.Reader) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(data)
	if err != nil {
		return nil, err
	}

	links := []string{}
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")

		normalized, err := normalizeLink(p.baseURL, link)
		if err != nil {
			log.Debugf("skipping external link: %s", link)
			return
		}

		links = append(links, normalized)
	})

	return links, nil
}
