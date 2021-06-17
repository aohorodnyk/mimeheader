package mimeheader_test

import (
	"errors"
	"fmt"
	"mime"
	"reflect"
	"testing"

	"github.com/aohorodnyk/mimeheader"
)

func TestParseMediaType(t *testing.T) {
	t.Parallel()

	for _, prov := range providerParseMediaType() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			b, err := mimeheader.ParseMediaType(prov.mtype)
			if (prov.expErr == nil) != (err == nil) {
				t.Errorf("Unexpected error.\nExpected: %+v\nActual: %+v\n", prov.expErr, err)
			}

			if err != nil && prov.expErr != nil && prov.expErr.Error() != err.Error() {
				t.Fatalf("Enuxpected error message.\nExpected: %s\nActual: %s\n", prov.expErr.Error(), err.Error())
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
			expErr: fmt.Errorf("%s: %w", mimeheader.ParseMediaTypeErrMsg, errors.New("mime: no media type")),
			exp:    mimeheader.MimeType{},
		},
		{
			name:   "Wrong wildcard",
			mtype:  "*/plain",
			expErr: mimeheader.ErrMimeTypeWildcard,
			exp:    mimeheader.MimeType{},
		},
		{
			name:   "Wrong delimiters",
			mtype:  "text/plain;;",
			expErr: fmt.Errorf("%s: %w", mimeheader.ParseMediaTypeErrMsg, errors.New("mime: invalid media parameter")),
			exp:    mimeheader.MimeType{},
		},
		{
			name:   "Invalid parts number",
			mtype:  "*-plain",
			expErr: mimeheader.ErrMimeTypeParts,
			exp:    mimeheader.MimeType{},
		},
		{
			name:   "Wrong parameter",
			mtype:  "text/plain; p=",
			expErr: fmt.Errorf("%s: %w", mimeheader.ParseMediaTypeErrMsg, mime.ErrInvalidMediaParameter),
			exp:    mimeheader.MimeType{},
		},
	}
}
