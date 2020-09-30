package linkstorage

import "sync"

type state int

const (
	stateNoLink state = iota
	stateUntouched
	stateTouched
	stateProcessed
)

// inMemStorage is concurrency safe
type inMemStorage struct {
	sync.Mutex
	links map[string]state
}

func NewInMemory() *inMemStorage {
	return &inMemStorage{
		links: make(map[string]state, 1000),
	}
}

func (s *inMemStorage) Add(links ...string) error {
	s.Lock()
	defer s.Unlock()

	for _, l := range links {
		if s.links[l] == stateNoLink {
			s.links[l] = stateUntouched

			log.Debugf("added link: %s", l)
		}

		log.Debugf("skipping link: %s", l)
	}

	return nil
}

func (s *inMemStorage) Get() (string, error) {
	s.Lock()
	defer s.Unlock()

	for link, state := range s.links {
		if state == stateUntouched {
			s.links[link] = stateTouched
			return link, nil
		}
	}

	log.Debug("no untouched links left")

	return "", ErrNoLinks
}

func (s *inMemStorage) Processed(link string) error {
	s.Lock()
	defer s.Unlock()

	_, ok := s.links[link]
	if !ok {
		// TODO: should we care about this?
		return nil
	}

	s.links[link] = stateProcessed

	log.Debugf("processed link: %s", link)

	return nil
}
