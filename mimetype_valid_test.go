package mimeheader_test

import (
	"testing"

	"github.com/aohorodnyk/mimeheader"
)

func TestMimeType_Valid(t *testing.T) {
	t.Parallel()

	for _, prov := range providerMimeTypeValid() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			act := prov.mtype.Valid()
			if prov.exp != act {
				t.Fatalf("Wrong validation result.\nExpected: %t,Actual: %t", prov.exp, act)
			}
		})
	}
}

type mimeTypeValid struct {
	name  string
	mtype mimeheader.MimeType
	exp   bool
}

func providerMimeTypeValid() []mimeTypeValid {
	return []mimeTypeValid{
		{
			name:  "Empty",
			mtype: mimeheader.MimeType{},
			exp:   false,
		},
		{
			name:  "Empty type",
			mtype: mimeheader.MimeType{Subtype: "plain"},
			exp:   false,
		},
		{
			name:  "Empty subtype",
			mtype: mimeheader.MimeType{Type: "text"},
			exp:   false,
		},
		{
			name:  "Type with wildcard",
			mtype: mimeheader.MimeType{Type: "*", Subtype: "plain"},
			exp:   false,
		},
		{
			name:  "Wildcards",
			mtype: mimeheader.MimeType{Type: "*", Subtype: "*"},
			exp:   true,
		},
		{
			name:  "Subtype with wildcard",
			mtype: mimeheader.MimeType{Type: "*", Subtype: "*"},
			exp:   true,
		},
		{
			name:  "text/plain",
			mtype: mimeheader.MimeType{Type: "text", Subtype: "plain"},
			exp:   true,
		},
	}
}
