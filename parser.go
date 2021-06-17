package mimeheader

import (
	"errors"
	"fmt"
	"mime"
	"strings"
)

const ParseMediaTypeErrMsg = "error in a parse media type"

// Errors triggered by mime type parser.
var (
	// ErrMimeTypeParts is wrong number of mime type parts.
	ErrMimeTypeParts = errors.New("wrong number of mime type parts")
	// ErrMimeTypeWildcard is wrong mimetype format.
	ErrMimeTypeWildcard = errors.New("mimetype cannot be as */plain")
)

// ParseMediaType parses media type to MimeType structure.
func ParseMediaType(mtype string) (MimeType, error) {
	mtype, params, err := mime.ParseMediaType(mtype)
	if err != nil {
		return MimeType{}, fmt.Errorf("%s: %w", ParseMediaTypeErrMsg, err)
	}

	mtypes := strings.SplitN(mtype, MimeSeparator, MimeParts)

	if len(mtypes) != MimeParts {
		return MimeType{}, ErrMimeTypeParts
	}

	if mtypes[0] == MimeAny && mtypes[1] != MimeAny {
		return MimeType{}, ErrMimeTypeWildcard
	}

	mt := MimeType{
		Type:    mtypes[0],
		Subtype: mtypes[1],
		Params:  params,
	}

	return mt, nil
}
