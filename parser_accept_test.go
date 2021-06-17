package mimeheader_test

import (
	"reflect"
	"testing"

	"github.com/aohorodnyk/mimeheader"
)

func TestParseAcceptHeader(t *testing.T) {
	t.Parallel()

	for _, prov := range providerParseAcceptHeader() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			act := mimeheader.ParseAcceptHeader(prov.header)
			if !reflect.DeepEqual(prov.exp, act) {
				t.Fatalf("AcceptHeaders are not equal.\nExpected: %+v\nActual: %+v", prov.exp, act)
			}
		})
	}
}

type parseAcceptHeader struct {
	name   string
	header string
	exp    mimeheader.AcceptHeader
}

func providerParseAcceptHeader() []parseAcceptHeader {
	return []parseAcceptHeader{
		{
			name:   "Empty",
			header: "",
			exp:    mimeheader.NewAcceptHeader([]mimeheader.MimeHeader{}),
		},
		{
			name:   "Empty object",
			header: "{}",
			exp:    mimeheader.NewAcceptHeader([]mimeheader.MimeHeader{}),
		},
		{
			name:   "Full wildcard",
			header: "*/*",
			exp: mimeheader.NewAcceptHeader([]mimeheader.MimeHeader{
				{
					MimeType: mimeheader.MimeType{
						Type:    "*",
						Subtype: "*",
						Params:  map[string]string{},
					},
					Quality: 1.0,
				},
			}),
		},
		{
			name:   "Header with sorting",
			header: "*/*; q=0.9; s=1, image/*; q=0.9; s=4, application/json; q=0.9; b=3;, text/plain",
			exp: mimeheader.NewAcceptHeader([]mimeheader.MimeHeader{
				{
					MimeType: mimeheader.MimeType{
						Type:    "text",
						Subtype: "plain",
						Params:  map[string]string{},
					},
					Quality: 1.0,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "application",
						Subtype: "json",
						Params:  map[string]string{"q": "0.9", "b": "3"},
					},
					Quality: 0.9,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "image",
						Subtype: "*",
						Params:  map[string]string{"q": "0.9", "s": "4"},
					},
					Quality: 0.9,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "*",
						Subtype: "*",
						Params:  map[string]string{"q": "0.9", "s": "1"},
					},
					Quality: 0.9,
				},
			}),
		},
		{
			name:   "Lost wildcard",
			header: "*/* q=0.9; s=1, image/*; q=0.9; s=4, application/json; q=0.9; b=3;, text/plain",
			exp: mimeheader.NewAcceptHeader([]mimeheader.MimeHeader{
				{
					MimeType: mimeheader.MimeType{
						Type:    "text",
						Subtype: "plain",
						Params:  map[string]string{},
					},
					Quality: 1.0,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "application",
						Subtype: "json",
						Params:  map[string]string{"q": "0.9", "b": "3"},
					},
					Quality: 0.9,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "image",
						Subtype: "*",
						Params:  map[string]string{"q": "0.9", "s": "4"},
					},
					Quality: 0.9,
				},
			}),
		},
	}
}
