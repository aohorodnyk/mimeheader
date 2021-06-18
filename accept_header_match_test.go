package mimeheader_test

import (
	"testing"

	"github.com/aohorodnyk/mimeheader"
)

func TestAcceptHeader_Match(t *testing.T) {
	t.Parallel()

	for _, prov := range providerAcceptHeaderMatch() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			ah := mimeheader.ParseAcceptHeader(prov.header)
			act := ah.Match(prov.ctype)
			if act != prov.exp {
				t.Fatalf("Failed match validation.\nExpected: %t\nActual: %t", prov.exp, act)
			}
		})
	}
}

type acceptHeaderMatch struct {
	name   string
	header string
	ctype  string
	exp    bool
}

func providerAcceptHeaderMatch() []acceptHeaderMatch {
	return []acceptHeaderMatch{
		{
			name:   "Empty header",
			header: "",
			ctype:  "application/json",
			exp:    false,
		},
		{
			name:   "Not match header",
			header: "text/plain, text/*, image/*",
			ctype:  "application/json",
			exp:    false,
		},
		{
			name:   "Header with typo in application",
			header: "text/plain, text/*, image/*, aplication/json", //nolint:misspell // Not a typo, this error was done for test purpose.
			ctype:  "application/json",
			exp:    false,
		},
		{
			name:   "JSONP type, instead of JSON",
			header: "text/plain, text/*, image/*, application/jsonp",
			ctype:  "application/json",
			exp:    false,
		},
		{
			name:   "application/*",
			header: "text/plain, text/*, image/*, application/*",
			ctype:  "application/json",
			exp:    true,
		},
		{
			name:   "application/json",
			header: "text/plain, text/*, image/*, application/json",
			ctype:  "application/json",
			exp:    true,
		},
	}
}
