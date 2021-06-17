package mimeheader_test

import (
	"testing"

	"github.com/aohorodnyk/mimeheader"
)

func TestMimeType_MatchText(t *testing.T) {
	t.Parallel()

	for _, prov := range providerMimeTypeMatchText() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			act := prov.b.MatchText(prov.t)

			if prov.exp != act {
				t.Fatalf("Match is not equal to expected value. Expected: %t. Actual: %t", prov.exp, act)
			}
		})
	}
}

type mimeTypeMatchText struct {
	name string
	b    mimeheader.MimeType
	t    string
	exp  bool
}

func providerMimeTypeMatchText() []mimeTypeMatchText {
	return []mimeTypeMatchText{
		{
			name: "Match application/*, positive",
			b: mimeheader.MimeType{
				Type:    "application",
				Subtype: "*",
				Params:  nil,
			},
			t:   "application/json;t=1;q=0.9",
			exp: true,
		},
		{
			name: "Match application/*, positive",
			b: mimeheader.MimeType{
				Type:    "application",
				Subtype: "json",
				Params:  nil,
			},
			t:   "application/*;t=1;q=0.9",
			exp: true,
		},
		{
			name: "Match application/*, positive",
			b: mimeheader.MimeType{
				Type:    "text",
				Subtype: "json",
				Params:  nil,
			},
			t:   "application/*;t=1;q=0.9",
			exp: false,
		},
	}
}
