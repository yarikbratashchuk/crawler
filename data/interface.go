// Package data holds primitives directly related to the data processing
package data

// Processor abstracts data processing
type Processor interface {
	Process(Data) error
}

type Data struct {
	Link string
	// more fields later
}
