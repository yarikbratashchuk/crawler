package crawler_test

import (
	"strings"
	"testing"

	"github.com/monzo/crawler/crawler"
)

func TestNaiveParser(t *testing.T) {
	data := `
	<html>
		<head></head>
			<a href="/hello" />
			<a href="/another/hello" />
		<body></body>
	</html>
	`

	parser, err := crawler.NewNaiveParser("somedomain.com")
	if err != nil {
		t.Error(err)
	}
	links, err := parser.Links(strings.NewReader(data))
	if err != nil {
		t.Error(err)
	}

	if len(links) != 2 {
		t.Fatal("should be 2 links")
	}
	// we don't check link normalization here
}
