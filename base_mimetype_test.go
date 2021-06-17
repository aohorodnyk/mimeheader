package mimeheader_test

import (
	"testing"

	"github.com/aohorodnyk/mimeheader"
)

func TestBase_MimeType(t *testing.T) {
	t.Parallel()

	for _, prov := range providerBaseMimeType() {
		prov := prov

		act := prov.b.MimeType()
		if prov.exp != act {
			t.Fatalf("Mime type is not equal to expected.\nExpected: %s\nActual: %s\n", prov.exp, act)
		}
	}
}

type baseMimeType struct {
	name string
	b    mimeheader.MimeType
	exp  string
}

func providerBaseMimeType() []baseMimeType {
	return []baseMimeType{
		{
			name: "Empty",
			b: mimeheader.MimeType{
				Type:    "",
				Subtype: "",
				Params:  map[string]string{"param": "a"},
			},
			exp: "*/*",
		},
		{
			name: "Empty type",
			b: mimeheader.MimeType{
				Type:    "",
				Subtype: "plain",
				Params:  map[string]string{"param": "a"},
			},
			exp: "*/plain",
		},
		{
			name: "Empty subtype",
			b: mimeheader.MimeType{
				Type:    "text",
				Subtype: "",
				Params:  map[string]string{"param": "a"},
			},
			exp: "text/*",
		},
		{
			name: "Empty subtype",
			b: mimeheader.MimeType{
				Type:    "text",
				Subtype: "plain",
				Params:  map[string]string{"param": "a"},
			},
			exp: "text/plain",
		},
	}
}
