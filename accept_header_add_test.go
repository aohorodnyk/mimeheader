package mimeheader_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aohorodnyk/mimeheader"
)

func ExampleAcceptHeader_Add() {
	ah := mimeheader.ParseAcceptHeader("image/png")
	fmt.Println(ah.Match("application/json"))

	ah.Add(mimeheader.MimeHeader{MimeType: mimeheader.MimeType{Type: "application", Subtype: "*"}})
	fmt.Println(ah.Match("application/json"))
	// Output:
	// false
	// true
}

func TestAcceptHeader_Add(t *testing.T) {
	t.Parallel()

	for _, prov := range providerAcceptHeaderAdd() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			prov.act.Add(prov.add)

			if !reflect.DeepEqual(prov.exp, prov.act) {
				t.Fatalf("Add data has unexpected result.\nExpected: %+v\nActual: %+v", prov.exp, prov.act)
			}
		})
	}
}

type acceptHeaderAdd struct {
	name string
	act  mimeheader.AcceptHeader
	add  mimeheader.MimeHeader
	exp  mimeheader.AcceptHeader
}

func providerAcceptHeaderAdd() []acceptHeaderAdd {
	return []acceptHeaderAdd{
		{
			name: "Add empty header to empty AcceptHeader",
			act:  mimeheader.NewAcceptHeaderPlain([]mimeheader.MimeHeader{}),
			add:  mimeheader.MimeHeader{},
			exp:  mimeheader.NewAcceptHeaderPlain([]mimeheader.MimeHeader{}),
		},
		{
			name: "Add wrong subtype header to empty AcceptHeader",
			act:  mimeheader.NewAcceptHeaderPlain([]mimeheader.MimeHeader{}),
			add:  mimeheader.MimeHeader{MimeType: mimeheader.MimeType{Type: "text"}},
			exp:  mimeheader.NewAcceptHeaderPlain([]mimeheader.MimeHeader{}),
		},
		{
			name: "Add wrong type header to empty AcceptHeader",
			act:  mimeheader.NewAcceptHeaderPlain([]mimeheader.MimeHeader{}),
			add:  mimeheader.MimeHeader{MimeType: mimeheader.MimeType{Subtype: "plain"}},
			exp:  mimeheader.NewAcceptHeaderPlain([]mimeheader.MimeHeader{}),
		},
		{
			name: "Add a header to empty AcceptHeader",
			act:  mimeheader.NewAcceptHeaderPlain([]mimeheader.MimeHeader{}),
			add: mimeheader.MimeHeader{
				MimeType: mimeheader.MimeType{
					Type:    "text",
					Subtype: "*",
					Params:  map[string]string{"q": "1.0"},
				},
				Quality: 1.0,
			},
			exp: mimeheader.NewAcceptHeaderPlain([]mimeheader.MimeHeader{
				{
					MimeType: mimeheader.MimeType{
						Type:    "text",
						Subtype: "*",
						Params:  map[string]string{"q": "1.0"},
					},
					Quality: 1.0,
				},
			}),
		},
		{
			name: "Add a header to an AcceptHeader",
			act: mimeheader.NewAcceptHeaderPlain([]mimeheader.MimeHeader{
				{
					MimeType: mimeheader.MimeType{
						Type:    "text",
						Subtype: "javascript",
						Params:  map[string]string{"q": "0.9"},
					},
					Quality: 1.0,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "application",
						Subtype: "json",
						Params:  map[string]string{"q": "0.9"},
					},
					Quality: 0.9,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "application",
						Subtype: "xml",
						Params:  map[string]string{},
					},
					Quality: 1.0,
				},
			}),
			add: mimeheader.MimeHeader{
				MimeType: mimeheader.MimeType{
					Type:    "text",
					Subtype: "*",
					Params:  map[string]string{},
				},
				Quality: 1.0,
			},
			exp: mimeheader.NewAcceptHeaderPlain([]mimeheader.MimeHeader{
				{
					MimeType: mimeheader.MimeType{
						Type:    "text",
						Subtype: "javascript",
						Params:  map[string]string{"q": "0.9"},
					},
					Quality: 1.0,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "application",
						Subtype: "xml",
						Params:  map[string]string{},
					},
					Quality: 1.0,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "text",
						Subtype: "*",
						Params:  map[string]string{},
					},
					Quality: 1.0,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "application",
						Subtype: "json",
						Params:  map[string]string{"q": "0.9"},
					},
					Quality: 0.9,
				},
			}),
		},
	}
}
