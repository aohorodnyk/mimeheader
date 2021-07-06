package mimeheader_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aohorodnyk/mimeheader"
)

func ExampleAcceptHeader_Set() {
	ah := mimeheader.ParseAcceptHeader("image/png")
	fmt.Println(ah.Match("application/json"))

	ah.Add(mimeheader.MimeHeader{MimeType: mimeheader.MimeType{Type: "application", Subtype: "*"}})
	fmt.Println(ah.Match("application/json"))
}

func TestAcceptHeader_Set(t *testing.T) {
	t.Parallel()

	for _, prov := range provideAcceptHeaderSet() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			act := mimeheader.NewAcceptHeaderPlain(nil)
			act.Set(prov.headers)

			if !reflect.DeepEqual(prov.exp, act) {
				t.Fatalf("Set data unexpected result.\nExpected: %+v\nActual: %+v", prov.exp, act)
			}
		})
	}
}

type acceptHeaderSet struct {
	name    string
	headers []mimeheader.MimeHeader
	exp     mimeheader.AcceptHeader
}

func provideAcceptHeaderSet() []acceptHeaderSet {
	return []acceptHeaderSet{
		{
			name: "Set complex",
			headers: []mimeheader.MimeHeader{
				{
					MimeType: mimeheader.MimeType{
						Type:    "text",
						Subtype: "*",
						Params:  map[string]string{},
					},
					Quality: 0.9,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "text",
						Subtype: "html",
						Params:  map[string]string{"test": "123"},
					},
					Quality: 1.0,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "application",
						Subtype: "xhtml+xml",
						Params:  map[string]string{},
					},
					Quality: 1.0,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "application",
						Subtype: "xml",
						Params:  map[string]string{},
					},
					Quality: 0.9,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "image",
						Subtype: "webp",
						Params:  map[string]string{},
					},
					Quality: 1.0,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "*",
						Subtype: "*",
						Params:  map[string]string{},
					},
					Quality: 0.8,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "application",
						Subtype: "hjson",
						Params:  map[string]string{},
					},
					Quality: 0.8,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "application",
						Subtype: "json",
						Params:  map[string]string{"test": "tere"},
					},
					Quality: 0.8,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "application",
						Subtype: "*",
						Params:  map[string]string{},
					},
					Quality: 0.8,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "",
						Subtype: "json",
						Params:  map[string]string{},
					},
					Quality: 0.8,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "application",
						Subtype: "",
						Params:  map[string]string{},
					},
					Quality: 0.8,
				},
			},
			exp: mimeheader.NewAcceptHeaderPlain([]mimeheader.MimeHeader{
				{
					MimeType: mimeheader.MimeType{
						Type:    "text",
						Subtype: "html",
						Params:  map[string]string{"test": "123"},
					},
					Quality: 1.0,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "application",
						Subtype: "xhtml+xml",
						Params:  map[string]string{},
					},
					Quality: 1.0,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "image",
						Subtype: "webp",
						Params:  map[string]string{},
					},
					Quality: 1.0,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "application",
						Subtype: "xml",
						Params:  map[string]string{},
					},
					Quality: 0.9,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "text",
						Subtype: "*",
						Params:  map[string]string{},
					},
					Quality: 0.9,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "application",
						Subtype: "json",
						Params:  map[string]string{"test": "tere"},
					},
					Quality: 0.8,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "application",
						Subtype: "hjson",
						Params:  map[string]string{},
					},
					Quality: 0.8,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "application",
						Subtype: "*",
						Params:  map[string]string{},
					},
					Quality: 0.8,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "*",
						Subtype: "*",
						Params:  map[string]string{},
					},
					Quality: 0.8,
				},
			}),
		},
		{
			name: "Set by types",
			headers: []mimeheader.MimeHeader{
				{
					MimeType: mimeheader.MimeType{
						Type:    "text",
						Subtype: "*",
						Params:  map[string]string{},
					},
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "*",
						Subtype: "*",
						Params:  map[string]string{},
					},
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "application",
						Subtype: "*",
						Params:  map[string]string{},
					},
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "application",
						Subtype: "json",
						Params:  map[string]string{},
					},
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "text",
						Subtype: "plain",
						Params:  map[string]string{},
					},
				},
			},
			exp: mimeheader.NewAcceptHeaderPlain([]mimeheader.MimeHeader{
				{
					MimeType: mimeheader.MimeType{
						Type:    "application",
						Subtype: "json",
						Params:  map[string]string{},
					},
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "text",
						Subtype: "plain",
						Params:  map[string]string{},
					},
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "text",
						Subtype: "*",
						Params:  map[string]string{},
					},
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "application",
						Subtype: "*",
						Params:  map[string]string{},
					},
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "*",
						Subtype: "*",
						Params:  map[string]string{},
					},
				},
			}),
		},
		{
			name: "Set by params",
			headers: []mimeheader.MimeHeader{
				{
					MimeType: mimeheader.MimeType{
						Type:    "test",
						Subtype: "1",
						Params:  map[string]string{},
					},
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "test",
						Subtype: "2",
						Params:  map[string]string{"p1": "g3"},
					},
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "test",
						Subtype: "3",
						Params:  map[string]string{"p1": "g3"},
					},
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "test",
						Subtype: "4",
						Params:  map[string]string{"p1": "g3", "p2": "g3"},
					},
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "test",
						Subtype: "5",
						Params:  map[string]string{"p1": "g3", "p2": "g3", "p3": "g3"},
					},
				},
			},
			exp: mimeheader.NewAcceptHeaderPlain([]mimeheader.MimeHeader{
				{
					MimeType: mimeheader.MimeType{
						Type:    "test",
						Subtype: "5",
						Params:  map[string]string{"p1": "g3", "p2": "g3", "p3": "g3"},
					},
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "test",
						Subtype: "4",
						Params:  map[string]string{"p1": "g3", "p2": "g3"},
					},
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "test",
						Subtype: "2",
						Params:  map[string]string{"p1": "g3"},
					},
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "test",
						Subtype: "3",
						Params:  map[string]string{"p1": "g3"},
					},
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "test",
						Subtype: "1",
						Params:  map[string]string{},
					},
				},
			}),
		},
		{
			name: "Set by quality only",
			headers: []mimeheader.MimeHeader{
				{
					MimeType: mimeheader.MimeType{
						Type:    "text",
						Subtype: "*",
						Params:  map[string]string{},
					},
					Quality: 0.3,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "*",
						Subtype: "*",
						Params:  map[string]string{},
					},
					Quality: 1.0,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "application",
						Subtype: "*",
						Params:  map[string]string{},
					},
					Quality: 0.5,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "application",
						Subtype: "json",
						Params:  map[string]string{},
					},
					Quality: 0.1,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "text",
						Subtype: "plain",
						Params:  map[string]string{},
					},
					Quality: 0.9,
				},
			},
			exp: mimeheader.NewAcceptHeaderPlain([]mimeheader.MimeHeader{
				{
					MimeType: mimeheader.MimeType{
						Type:    "*",
						Subtype: "*",
						Params:  map[string]string{},
					},
					Quality: 1.0,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "text",
						Subtype: "plain",
						Params:  map[string]string{},
					},
					Quality: 0.9,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "application",
						Subtype: "*",
						Params:  map[string]string{},
					},
					Quality: 0.5,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "text",
						Subtype: "*",
						Params:  map[string]string{},
					},
					Quality: 0.3,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "application",
						Subtype: "json",
						Params:  map[string]string{},
					},
					Quality: 0.1,
				},
			}),
		},
		{
			name:    "Set empty array",
			headers: []mimeheader.MimeHeader{},
			exp:     mimeheader.NewAcceptHeaderPlain([]mimeheader.MimeHeader{}),
		},
		{
			name: "Issue #6 sort by params",
			headers: []mimeheader.MimeHeader{
				{
					MimeType: mimeheader.MimeType{
						Type:    "application",
						Subtype: "xml",
						Params:  map[string]string{"q": "1.0", "test": "t"},
					},
					Quality: 1,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "application",
						Subtype: "json",
						Params:  map[string]string{"charset": "utf-8", "test": "t"},
					},
					Quality: 1,
				},
			},
			exp: mimeheader.NewAcceptHeaderPlain([]mimeheader.MimeHeader{
				{
					MimeType: mimeheader.MimeType{
						Type:    "application",
						Subtype: "json",
						Params:  map[string]string{"charset": "utf-8", "test": "t"},
					},
					Quality: 1,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "application",
						Subtype: "xml",
						Params:  map[string]string{"q": "1.0", "test": "t"},
					},
					Quality: 1,
				},
			}),
		},
	}
}
