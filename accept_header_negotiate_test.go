package mimeheader_test

import (
	"fmt"
	"testing"

	"github.com/aohorodnyk/mimeheader"
)

func ExampleAcceptHeader_Negotiate() {
	header := "image/*; q=0.9; s=4, application/json; q=0.9; b=3;, text/plain,image/png;q=0.9, image/jpeg,image/svg;q=0.8"
	ah := mimeheader.ParseAcceptHeader(header)

	fmt.Println(ah.Negotiate([]string{"application/xml", "image/tiff"}, "text/javascript"))
	fmt.Println(ah.Negotiate([]string{"application/xml", "image/png"}, "text/javascript"))
	fmt.Println(ah.Negotiate([]string{"application/xml", "image/svg"}, "text/javascript"))
	fmt.Println(ah.Negotiate([]string{"text/dart", "application/dart"}, "text/javascript"))
	// Output:
	// image/* image/tiff true
	// image/png image/png true
	// image/* image/svg true
	//  text/javascript false
}

func TestAcceptHeader_Negotiate(t *testing.T) {
	t.Parallel()

	for _, prov := range providerAcceptHeaderNegotiate() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			actHeader, actMType, actMatched := prov.ah.Negotiate(prov.ctypes, prov.dtype)
			if actHeader.String() != prov.expHeader.String() {
				t.Errorf("Wrong header matched.\nExpected: %v\nActual: %v", prov.expHeader, actHeader)
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
	expHeader  mimeheader.MimeHeader
	expMType   string
	expMatched bool
}

func providerAcceptHeaderNegotiate() []acceptHeaderNegotiate {
	return []acceptHeaderNegotiate{
		{
			name:       "Empty ctypes",
			ah:         mimeheader.ParseAcceptHeader(""),
			ctypes:     []string{},
			dtype:      "text/plain",
			expHeader:  mimeheader.MimeHeader{},
			expMType:   "text/plain",
			expMatched: false,
		},
		{
			name:       "Wrong ctypes",
			ah:         mimeheader.ParseAcceptHeader("*/*"),
			ctypes:     []string{"application/json;param="},
			dtype:      "text/plain",
			expHeader:  mimeheader.MimeHeader{},
			expMType:   "text/plain",
			expMatched: false,
		},
		{
			name:       "Wildcard ctypes",
			ah:         mimeheader.ParseAcceptHeader("*/*"),
			ctypes:     []string{"application/json;param=1"},
			dtype:      "text/plain",
			expHeader:  mimeheader.MimeHeader{MimeType: mimeheader.MimeType{Type: "*", Subtype: "*"}},
			expMType:   "application/json",
			expMatched: true,
		},
		{
			name:       "Sorted list of types with the same structure image/png",
			ah:         mimeheader.ParseAcceptHeader("application/json;q=1.0,*/*;q=1.0; param=wild,image/png;q=1.0;param=test"),
			ctypes:     []string{"application/json;param=1", "image/png"},
			dtype:      "text/plain",
			expHeader:  mimeheader.MimeHeader{MimeType: mimeheader.MimeType{Type: "image", Subtype: "png"}},
			expMType:   "image/png",
			expMatched: true,
		},
		{
			name:       "Sorted list of types with the same structure */*",
			ah:         mimeheader.ParseAcceptHeader("application/json;q=1.0,*/*;q=1.0; param=wild,image/png;q=1.0;param=test"),
			ctypes:     []string{"application/xml;param=1", "text/plain"},
			dtype:      "text/javascript",
			expHeader:  mimeheader.MimeHeader{MimeType: mimeheader.MimeType{Type: "*", Subtype: "*"}},
			expMType:   "application/xml",
			expMatched: true,
		},
		{
			name:       "Sorted list of types with the same structure",
			ah:         mimeheader.ParseAcceptHeader("application/json;q=1.0,image/*;q=1.0;param=test"),
			ctypes:     []string{"test/xml;param=1", "text/plain"},
			dtype:      "text/javascript",
			expHeader:  mimeheader.MimeHeader{},
			expMType:   "text/javascript",
			expMatched: false,
		},
	}
}
