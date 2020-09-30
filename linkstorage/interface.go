// Package linkstorage holds everything related to storing links during the
// crawling process
package linkstorage

import "errors"

// Storage describes functionality for the link storage
// There are possible 3 states of the link:
// - untouched
// - touched
// - processed
type Storage interface {
	// Add links to the storage.
	// It doesn't matter if some already exists in the storage
	Add(links ...string) error
	// Get returns 'untouched' link and marks it as 'touched'.
	// If there are no 'untouched' links ErrNoLinks is returned.
	Get() (string, error)
	// Processed marks the link as 'processed'
	Processed(string) error
}

var ErrNoLinks = errors.New("no links to visit")
