package mimeheader_test

import (
	"testing"

	"github.com/aohorodnyk/mimeheader"
)

func TestMimeType_MimeType(t *testing.T) {
	t.Parallel()

	for _, prov := range providerMimeTypeMimeType() {
		prov := prov

		act := prov.b.String()
		if prov.exp != act {
			t.Fatalf("Mime type is not equal to expected.\nExpected: %s\nActual: %s\n", prov.exp, act)
		}
	}
}

type mimeTypeMimeType struct {
	name string
	b    mimeheader.MimeType
	exp  string
}

func providerMimeTypeMimeType() []mimeTypeMimeType {
	return []mimeTypeMimeType{
		{
			name: "Empty",
			b: mimeheader.MimeType{
				Type:    "",
				Subtype: "",
				Params:  map[string]string{"param": "a"},
			},
			exp: "",
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
