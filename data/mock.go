package data

import "fmt"

type mockProcessor struct {
	process func(Data) error
}

// NewMockProcessor creates mocking data processor
func NewMockProcessor(process func(Data) error) *mockProcessor {
	return &mockProcessor{process}
}

// Process implements Processor
func (p *mockProcessor) Process(data Data) error {
	log.Debug("processing %s", data.Link)

	return p.process(data)
}

func WriteLinkToStdout(data Data) error {
	fmt.Println(data.Link)

	return nil
}
