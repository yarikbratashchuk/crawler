package crawler

import (
	"errors"
	"time"

	"github.com/yarikbratashchuk/crawler/data"
	"github.com/yarikbratashchuk/crawler/linkstorage"
)

type Crawler struct {
	domain string

	fetcher Fetcher
	parser  Parser
	links   linkstorage.Storage
	data    data.Processor

	workers chan struct{}
}

func New(
	domain string,
	numWorkers uint,
	fetcher Fetcher,
	parser Parser,
	links linkstorage.Storage,
	data data.Processor,
) (*Crawler, error) {
	log.Debugf("setting up crawler: numWorkers: %d", numWorkers)

	if numWorkers == 0 {
		return nil, errors.New("number of workers should be greater than 0")
	}

	if fetcher == nil || parser == nil || links == nil || data == nil {
		return nil, errors.New("nil interfaces are not allowed")
	}

	return &Crawler{
		domain: domain,

		fetcher: fetcher,
		parser:  parser,
		links:   links,
		data:    data,

		workers: make(chan struct{}, numWorkers),
	}, nil
}

func (c *Crawler) Crawl() {
	log.Infof("crawling %s", c.domain)

	if err := c.links.Add(c.domain); err != nil {
		return
	}

	for {
		c.workers <- struct{}{}

		link, err := c.links.Get()
		if err != nil {
			if err == linkstorage.ErrNoLinks {
				// new links will not appear
				if len(c.workers) == 1 {
					break
				}
				time.Sleep(3 * time.Second)
			} else {
				log.Error(err)
			}
			<-c.workers
			continue
		}

		go func(link string) {
			c.crawl(link)
			<-c.workers
		}(link)
	}
}

// TODO: refactor with a for loop construction
func (c *Crawler) crawl(link string) {
	res, err := c.fetcher.Get(link)
	if err != nil {
		return
	}

	links, err := c.parser.Links(res)
	if err != nil {
		log.Errorf("finding links: %v", err)
		return
	}
	if err := c.links.Add(links...); err != nil {
		return
	}

	data := data.Data{Link: link}
	if err := c.data.Process(data); err != nil {
		return
	}

	if err := c.links.Processed(link); err != nil {
		return
	}
}
