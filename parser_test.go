package mimeheader_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/aohorodnyk/mimeheader"
)

func ExampleParseMediaType_wildcard() {
	// Parse media type
	mediaType := "application/*; q=1; param=test;"
	mimeType, err := mimeheader.ParseMediaType(mediaType)
	if err != nil {
		panic(err)
	}

	// Print string without params.
	fmt.Println(mimeType.String())
	// Print string with params.
	fmt.Println(mimeType.StringWithParams())
	// Parse input and match it.
	fmt.Println(mimeType.MatchText("application/json; param=test"))
	fmt.Println(mimeType.MatchText("application/xml; q=1"))
	fmt.Println(mimeType.MatchText("text/plain"))

	// Parse mime type.
	tmtype, err := mimeheader.ParseMediaType("application/json;q=0.3")
	if err != nil {
		panic(err)
	}

	// Match mime types.
	fmt.Println(mimeType.Match(tmtype))
	// Output:
	// application/*
	// application/*; param=test; q=1
	// true
	// true
	// false
	// true
}

func ExampleParseMediaType_exact() {
	// Parse media type
	mediaType := "application/json; q=1; param=test;"
	mimeType, err := mimeheader.ParseMediaType(mediaType)
	if err != nil {
		panic(err)
	}

	// Print string without params.
	fmt.Println(mimeType.String())
	// Print string with params.
	fmt.Println(mimeType.StringWithParams())
	// Parse input and match it.
	fmt.Println(mimeType.MatchText("application/json; param=test"))
	fmt.Println(mimeType.MatchText("application/xml; q=1"))
	fmt.Println(mimeType.MatchText("text/plain"))

	// Parse mime type.
	tmtype, err := mimeheader.ParseMediaType("application/json;q=0.3")
	if err != nil {
		panic(err)
	}

	// Match mime types.
	fmt.Println(mimeType.Match(tmtype))
	// Output:
	// application/json
	// application/json; param=test; q=1
	// true
	// false
	// false
	// true
}

func TestParseMediaType(t *testing.T) {
	t.Parallel()

	for _, prov := range providerParseMediaType() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			b, err := mimeheader.ParseMediaType(prov.mtype)
			if !errors.Is(err, prov.expErr) && !errors.As(err, &prov.expErr) {
				t.Errorf("Unexpected error.\nExpected: %#v\nActual: %#v\n", prov.expErr, err)
			}

			if !reflect.DeepEqual(prov.exp, b) {
				t.Fatalf("Unexpected MimType.\nExpected: %+v\nActual: %+v\n", prov.exp, b)
			}
		})
	}
}

type parseMediaType struct {
	name   string
	mtype  string
	expErr error
	exp    mimeheader.MimeType
}

func providerParseMediaType() []parseMediaType {
	return []parseMediaType{
		{
			name:  "Wildcard",
			mtype: "*/*",
			exp: mimeheader.MimeType{
				Type:    "*",
				Subtype: "*",
				Params:  map[string]string{},
			},
		},
		{
			name:  "Wildcard with params",
			mtype: "*/*; q=0.9;param=123;k=m",
			exp: mimeheader.MimeType{
				Type:    "*",
				Subtype: "*",
				Params:  map[string]string{"q": "0.9", "param": "123", "k": "m"},
			},
		},
		{
			name:  "Wildcard subtype",
			mtype: "text/*",
			exp: mimeheader.MimeType{
				Type:    "text",
				Subtype: "*",
				Params:  map[string]string{},
			},
		},
		{
			name:  "Wildcard subtype with delimiter",
			mtype: "text/*;",
			exp: mimeheader.MimeType{
				Type:    "text",
				Subtype: "*",
				Params:  map[string]string{},
			},
		},
		{
			name:  "Wildcard subtype with params",
			mtype: "  text/*  ; q=0.9;param=123;k=m",
			exp: mimeheader.MimeType{
				Type:    "text",
				Subtype: "*",
				Params:  map[string]string{"q": "0.9", "param": "123", "k": "m"},
			},
		},
		{
			name:  "Specific type",
			mtype: "  application/json   ",
			exp: mimeheader.MimeType{
				Type:    "application",
				Subtype: "json",
				Params:  map[string]string{},
			},
		},
		{
			name:  "To lower",
			mtype: "tEXt/plAiN; q=0.9;param=123;k=m",
			exp: mimeheader.MimeType{
				Type:    "text",
				Subtype: "plain",
				Params:  map[string]string{"q": "0.9", "param": "123", "k": "m"},
			},
		},
		{
			name:   "Empty error",
			mtype:  "",
			expErr: mimeheader.MimeParseErr{},
			exp:    mimeheader.MimeType{},
		},
		{
			name:   "Wrong wildcard",
			mtype:  "*/plain",
			expErr: mimeheader.MimeTypeWildcardErr{},
			exp:    mimeheader.MimeType{},
		},
		{
			name:   "Wrong delimiters",
			mtype:  "text/plain;;",
			expErr: mimeheader.MimeParseErr{},
			exp:    mimeheader.MimeType{},
		},
		{
			name:   "Invalid parts number",
			mtype:  "*-plain",
			expErr: mimeheader.MimeTypePartsErr{},
			exp:    mimeheader.MimeType{},
		},
		{
			name:   "Wrong parameter",
			mtype:  "text/plain; p=",
			expErr: mimeheader.MimeTypeWildcardErr{},
			exp:    mimeheader.MimeType{},
		},
	}
}
