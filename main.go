package main

import (
	"net"
	"net/http"
	"os"
	"time"

	"github.com/yarikbratashchuk/crawler/crawler"
	"github.com/yarikbratashchuk/crawler/data"
	"github.com/yarikbratashchuk/crawler/linkstorage"
)

func main() {
	conf, err := loadConfig()
	if err != nil {
		os.Exit(1)
	}

	setupLog(os.Stderr, conf.LogLevel)

	parsedDomain, err := crawler.ParseLink(conf.Domain)
	if err != nil {
		log.Errorf("parsing %s: %v", conf.Domain, err)
		os.Exit(1)
	}
	domain := parsedDomain.String()

	fetcher := crawler.NewHTTPFetcher(httpClientWithTimeout(conf.DialTimeout))
	parser, err := crawler.NewNaiveParser(domain)
	if err != nil {
		log.Errorf("creating parser: %v", err)
		os.Exit(1)
	}
	linkStorage := linkstorage.NewInMemory()
	dataProcessor := data.NewMockProcessor(data.WriteLinkToStdout)

	crawler, err := crawler.New(
		domain,
		conf.NumWorkers,
		fetcher,
		parser,
		linkStorage,
		dataProcessor,
	)
	if err != nil {
		log.Criticalf("setting up crawler: %v", err)
		os.Exit(1)
	}

	crawler.Crawl()

	log.Infof("shutting down...")
}

func httpClientWithTimeout(timeout uint) *http.Client {
	tr := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: time.Duration(timeout) * time.Second,
		}).Dial,
	}
	return &http.Client{Transport: tr}
}
