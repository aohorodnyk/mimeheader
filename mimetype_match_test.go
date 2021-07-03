package mimeheader_test

import (
	"fmt"
	"testing"

	"github.com/aohorodnyk/mimeheader"
)

func ExampleMimeType_Match() {
	// Parse media type
	mediaType := "application/xml; q=1; param=test;"

	mimeType, err := mimeheader.ParseMediaType(mediaType)
	if err != nil {
		panic(err)
	}

	// Parse input and match it.
	fmt.Println(mimeType.Match(mimeheader.MimeType{Type: "application", Subtype: "xml"}))
	fmt.Println(mimeType.Match(mimeheader.MimeType{Type: "application", Subtype: "json"}))
	// Output:
	// true
	// false
}

func TestMimetype_Match(t *testing.T) {
	t.Parallel()

	for _, prov := range providerMimeTypeMatch() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			act := prov.b.Match(prov.t)

			if prov.exp != act {
				t.Fatalf("Match is not equal to expected value. Expected: %t. Actual: %t", prov.exp, act)
			}
		})
	}
}

type mimeTypeMatch struct {
	name string
	b    mimeheader.MimeType
	t    mimeheader.MimeType
	exp  bool
}

func providerMimeTypeMatch() []mimeTypeMatch {
	return []mimeTypeMatch{
		{
			name: "Match text/*, positive",
			b: mimeheader.MimeType{
				Type:    "text",
				Subtype: "*",
				Params:  nil,
			},
			t: mimeheader.MimeType{
				Type:    "text",
				Subtype: "plain",
				Params:  nil,
			},
			exp: true,
		},
		{
			name: "Match specific",
			b: mimeheader.MimeType{
				Type:    "text",
				Subtype: "plain",
				Params:  nil,
			},
			t: mimeheader.MimeType{
				Type:    "text",
				Subtype: "plain",
				Params:  nil,
			},
			exp: true,
		},
		{
			name: "Match specific",
			b: mimeheader.MimeType{
				Type:    "text",
				Subtype: "plain",
				Params:  nil,
			},
			t: mimeheader.MimeType{
				Type:    "text",
				Subtype: "*",
				Params:  nil,
			},
			exp: true,
		},
		{
			name: "Match specific",
			b: mimeheader.MimeType{
				Type:    "text",
				Subtype: "plain",
				Params:  nil,
			},
			t: mimeheader.MimeType{
				Type:    "*",
				Subtype: "*",
				Params:  nil,
			},
			exp: true,
		},
		{
			name: "Match full wildcards",
			b: mimeheader.MimeType{
				Type:    "*",
				Subtype: "*",
				Params:  nil,
			},
			t: mimeheader.MimeType{
				Type:    "text",
				Subtype: "plain",
				Params:  nil,
			},
			exp: true,
		},
		{
			name: "Match pre wildcard",
			b: mimeheader.MimeType{
				Type:    "*",
				Subtype: "plain",
				Params:  nil,
			},
			t: mimeheader.MimeType{
				Type:    "text",
				Subtype: "plain",
				Params:  nil,
			},
			exp: true,
		},
		{
			name: "Match wrong",
			b: mimeheader.MimeType{
				Type:    "text",
				Subtype: "plain",
				Params:  nil,
			},
			t: mimeheader.MimeType{
				Type:    "text",
				Subtype: "plains",
				Params:  nil,
			},
			exp: false,
		},
		{
			name: "Match text/*, negative",
			b: mimeheader.MimeType{
				Type:    "text",
				Subtype: "*",
				Params:  nil,
			},
			t: mimeheader.MimeType{
				Type:    "test",
				Subtype: "*",
				Params:  nil,
			},
			exp: false,
		},
		{
			name: "Empty",
			b: mimeheader.MimeType{
				Type:    "",
				Subtype: "",
				Params:  nil,
			},
			t: mimeheader.MimeType{
				Type:    "test",
				Subtype: "*",
				Params:  nil,
			},
			exp: false,
		},
	}
}
