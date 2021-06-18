package mimeheader_test

import (
	"testing"

	"github.com/aohorodnyk/mimeheader"
)

func TestAcceptHeader_Negotiate(t *testing.T) {
	t.Parallel()

	for _, prov := range providerAcceptHeaderNegotiate() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			actHeader, actMType, actMatched := prov.ah.Negotiate(prov.ctypes, prov.dtype)
			if actHeader != prov.expHeader {
				t.Errorf("Wrong header matched.\nExpected: %s\nActual: %s", prov.expHeader, actHeader)
			}

			if actMType != prov.expMType {
				t.Errorf("Wrong mime type returned.\nExpected: %s\nActual: %s", prov.expMType, actMType)
			}

			if actMatched != prov.expMatched {
				t.Errorf("Unepected match result.\nExpected: %t\nActual: %t", prov.expMatched, actMatched)
			}
		})
	}
}

type acceptHeaderNegotiate struct {
	name       string
	ah         mimeheader.AcceptHeader
	ctypes     []string
	dtype      string
	expHeader  string
	expMType   string
	expMatched bool
}

func providerAcceptHeaderNegotiate() []acceptHeaderNegotiate {
	return []acceptHeaderNegotiate{
		{
			name:       "Empty ctypes",
			ah:         mimeheader.NewAcceptHeader([]mimeheader.MimeHeader{}),
			ctypes:     []string{},
			dtype:      "text/plain",
			expHeader:  "",
			expMType:   "text/plain",
			expMatched: false,
		},
		{
			name: "Wrong ctypes",
			ah: mimeheader.NewAcceptHeader([]mimeheader.MimeHeader{
				{
					MimeType: mimeheader.MimeType{
						Type:    "*",
						Subtype: "*",
					},
				},
			}),
			ctypes:     []string{"application/json;param="},
			dtype:      "text/plain",
			expHeader:  "",
			expMType:   "text/plain",
			expMatched: false,
		},
		{
			name: "Wildcard ctypes",
			ah: mimeheader.NewAcceptHeader([]mimeheader.MimeHeader{
				{
					MimeType: mimeheader.MimeType{
						Type:    "*",
						Subtype: "*",
					},
				},
			}),
			ctypes:     []string{"application/json;param=1"},
			dtype:      "text/plain",
			expHeader:  "*/*",
			expMType:   "application/json",
			expMatched: true,
		},
		{
			name: "Sorted list of types with the same structure",
			ah: mimeheader.NewAcceptHeader([]mimeheader.MimeHeader{
				{
					MimeType: mimeheader.MimeType{
						Type:    "application",
						Subtype: "json",
						Params:  map[string]string{"q": "1.0"},
					},
					Quality: 1.0,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "*",
						Subtype: "*",
						Params:  map[string]string{"q": "1.0", "param": "wild"},
					},
					Quality: 1.0,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "image",
						Subtype: "png",
						Params:  map[string]string{"q": "1.0", "param": "test"},
					},
					Quality: 1.0,
				},
			}),
			ctypes:     []string{"application/json;param=1", "image/png"},
			dtype:      "text/plain",
			expHeader:  "image/png",
			expMType:   "image/png",
			expMatched: true,
		},
		{
			name: "Sorted list of types with the same structure",
			ah: mimeheader.NewAcceptHeader([]mimeheader.MimeHeader{
				{
					MimeType: mimeheader.MimeType{
						Type:    "application",
						Subtype: "json",
						Params:  map[string]string{"q": "1.0"},
					},
					Quality: 1.0,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "*",
						Subtype: "*",
						Params:  map[string]string{"q": "1.0", "param": "wild"},
					},
					Quality: 1.0,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "image",
						Subtype: "png",
						Params:  map[string]string{"q": "1.0", "param": "test"},
					},
					Quality: 1.0,
				},
			}),
			ctypes:     []string{"application/xml;param=1", "text/plain"},
			dtype:      "text/javascript",
			expHeader:  "*/*",
			expMType:   "application/xml",
			expMatched: true,
		},
		{
			name: "Sorted list of types with the same structure",
			ah: mimeheader.NewAcceptHeader([]mimeheader.MimeHeader{
				{
					MimeType: mimeheader.MimeType{
						Type:    "application",
						Subtype: "*",
						Params:  map[string]string{"q": "1.0"},
					},
					Quality: 1.0,
				},
				{
					MimeType: mimeheader.MimeType{
						Type:    "image",
						Subtype: "*",
						Params:  map[string]string{"q": "1.0", "param": "test"},
					},
					Quality: 1.0,
				},
			}),
			ctypes:     []string{"test/xml;param=1", "text/plain"},
			dtype:      "text/javascript",
			expHeader:  "",
			expMType:   "text/javascript",
			expMatched: false,
		},
	}
}
