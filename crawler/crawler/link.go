package crawler

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/goware/urlx"
)

func ParseLink(link string) (*url.URL, error) {
	urlParsed, err := urlx.Parse(link)
	if err != nil {
		return nil, err
	}
	normalized, err := urlx.Normalize(urlParsed)
	if err != nil {
		return nil, err
	}

	return url.Parse(normalized)
}

func normalizeLink(base *url.URL, link string) (string, error) {
	parsedLink, err := url.Parse(link)
	if err != nil {
		return "", err
	}
	if parsedLink.Host != "" && parsedLink.Host != base.Host {
		return "", errors.New("external link")
	}

	return base.ResolveReference(parsedLink).String(), nil
}

func validLink(link *url.URL) bool {
	res, err := http.Get(link.String())
	if err != nil || res.StatusCode != 200 {
		return false
	}
	return true
}
