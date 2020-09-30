package crawler

import (
	"testing"

	"github.com/yarikbratashchuk/crawler/data"
	"github.com/yarikbratashchuk/crawler/linkstorage"
)

// NOTE: very basic test.. I'm running out of time
func TestCrawler(t *testing.T) {
	domain := "http://example.com/"
	numWorkers := uint(3)

	linkdata := map[string]string{
		"http://example.com/":           "<html><body><a href='/hello' /></body></html>",
		"http://example.com/hello":      "<html><body><a href='/another' /></body></html>",
		"http://example.com/another":    "<html><body><a href='/thelastone' /></body></html>",
		"http://example.com/thelastone": "<html><body><a href='/' /></body></html>",
	}

	fetcher := NewInMemFetcher(linkdata)
	parser, _ := NewNaiveParser(domain)
	linkStorage := linkstorage.NewInMemory()

	collectedLinks := []string{}
	dataProcessor := data.NewMockProcessor(func(data data.Data) error {
		collectedLinks = append(collectedLinks, data.Link)
		return nil
	})

	crawler, err := New(
		domain,
		numWorkers,
		fetcher,
		parser,
		linkStorage,
		dataProcessor,
	)
	if err != nil {
		t.Error(err)
	}

	crawler.Crawl()

	if len(collectedLinks) == 0 {
		t.Error("no links collected")
	}

	for _, link := range collectedLinks {
		if _, ok := linkdata[link]; !ok {
			t.Errorf("link %s was not collected", link)
		}
	}
}
