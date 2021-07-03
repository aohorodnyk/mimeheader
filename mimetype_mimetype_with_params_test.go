package mimeheader_test

import (
	"testing"

	"github.com/aohorodnyk/mimeheader"
)

func TestMimeType_MimeTypeWithParams(t *testing.T) {
	t.Parallel()

	for _, prov := range providerMimeTypeWithParams() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			act := prov.b.StringWithParams()
			if prov.exp != act {
				t.Fatalf("Mime type is not equal to expected.\nExpected: %s\nActual: %s\n", prov.exp, act)
			}
		})
	}
}

func providerMimeTypeWithParams() []mimeTypeMimeType {
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
			exp: "*/plain; param=a",
		},
		{
			name: "Empty subtype",
			b: mimeheader.MimeType{
				Type:    "text",
				Subtype: "",
				Params:  map[string]string{"param": "a"},
			},
			exp: "text/*; param=a",
		},
		{
			name: "Empty subtype",
			b: mimeheader.MimeType{
				Type:    "text",
				Subtype: "plain",
				Params:  map[string]string{"param": "a"},
			},
			exp: "text/plain; param=a",
		},
	}
}
