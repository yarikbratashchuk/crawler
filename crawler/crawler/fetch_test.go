package crawler_test

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/monzo/crawler/crawler"
)

func TestHTTPFetcher(t *testing.T) {
	link := "http://example.com"

	fetcher := crawler.NewHTTPFetcher(http.DefaultClient)
	data, err := fetcher.Get(link)
	// TODO: use assert lib
	if err != nil {
		t.Error(err)
	}

	bytes, err := ioutil.ReadAll(data)
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(string(bytes), "Example Domain") {
		t.Errorf("weird response: %v", string(bytes))
	}
}
