package crawler

import (
	"net/url"
	"testing"
)

func TestParseLink(t *testing.T) {
	cases := []struct {
		name                string
		link                string
		expectedStringified string
		expectedErr         bool
	}{{
		name:                "general case",
		link:                "example.com",
		expectedStringified: "http://example.com",
		expectedErr:         false,
	}, {
		name:                "error case",
		link:                "/hello/world",
		expectedStringified: "",
		expectedErr:         true,
	}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			parsed, err := ParseLink(c.link)
			if err != nil && !c.expectedErr {
				t.Error(err)
				return
			}
			if err == nil && c.expectedErr {
				t.Errorf("should be an error but got %s", parsed)
				return
			}

			if !c.expectedErr && parsed.String() != c.expectedStringified {
				t.Errorf("expected: %s, got: %s", c.expectedStringified, parsed)
			}
		})
	}
}

func TestNormalizeLink(t *testing.T) {
	cases := []struct {
		name         string
		base         string
		link         string
		expectedErr  bool
		expectedLink string
	}{{
		name:         "relative path",
		base:         "http://google.com/",
		link:         "/hello",
		expectedErr:  false,
		expectedLink: "http://google.com/hello",
	}, {
		name:         "relative path without /",
		base:         "http://google.com/",
		link:         "hello/another",
		expectedErr:  false,
		expectedLink: "http://google.com/hello/another",
	}, {
		name:         "with a host in the link",
		base:         "http://google.com",
		link:         "http://google.com/hello",
		expectedErr:  false,
		expectedLink: "http://google.com/hello",
	}, {
		name:         "different hosts",
		base:         "http://google.com",
		link:         "http://mail.google.com/hello",
		expectedErr:  true,
		expectedLink: "",
	}}
	// TODO: more cases
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			base, err := url.Parse(c.base)
			if err != nil {
				t.Error(err)
			}
			got, err := normalizeLink(base, c.link)
			if err != nil && !c.expectedErr {
				t.Errorf(
					"expected err: %v, got: %v",
					c.expectedErr,
					err,
				)
				return
			}
			if err == nil && c.expectedErr {
				t.Errorf("expected error, got link: %s", got)
				return
			}
			if got != c.expectedLink {
				t.Errorf(
					"expected: %s, got: %s",
					c.expectedLink,
					got,
				)
			}
		})
	}
}
