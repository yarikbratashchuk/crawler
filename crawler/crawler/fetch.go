package crawler

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"strings"
	"sync"
)

// Fetcher abstracts link requests
type Fetcher interface {
	Get(string) (io.Reader, error)
}

type httpFetcher struct {
	client *http.Client
}

func NewHTTPFetcher(client *http.Client) Fetcher {
	return httpFetcher{client}
}

func (f httpFetcher) Get(link string) (io.Reader, error) {
	res, err := f.client.Get(link)
	if err != nil {
		log.Errorf("failed request to %s: %v", link, err)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Errorf("request to %s finished with code %v", link, res.StatusCode)
		return nil, errors.New("request faiiled")
	}

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, res.Body)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

type inMemFetcher struct {
	sync.Mutex
	linkdata map[string]string
}

func NewInMemFetcher(linkdata map[string]string) Fetcher {
	return &inMemFetcher{linkdata: linkdata}
}

func (f *inMemFetcher) Get(link string) (io.Reader, error) {
	f.Mutex.Lock()
	defer f.Mutex.Unlock()

	data, ok := f.linkdata[link]
	if !ok {
		return nil, errors.New("no data for such link")
	}

	return strings.NewReader(data), nil
}
