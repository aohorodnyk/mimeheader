package mimeheader

import (
	"mime"
	"strings"
)

// Error messages.
const (
	MimeParseErrMsg        = "error in a parse media type"
	MimeTypePartsErrMsg    = "wrong number of mime type parts"
	MimeTypeWildcardErrMsg = "mimetype cannot be as */plain"
)

// ParseMediaType parses media type to MimeType structure.
func ParseMediaType(mtype string) (MimeType, error) {
	mtype, params, err := mime.ParseMediaType(mtype)
	if err != nil {
		return MimeType{}, MimeParseErr{Err: err, Msg: MimeParseErrMsg}
	}

	mtypes := strings.SplitN(mtype, MimeSeparator, MimeParts)

	if len(mtypes) != MimeParts {
		return MimeType{}, MimeTypePartsErr{Msg: MimeTypePartsErrMsg}
	}

	if mtypes[0] == MimeAny && mtypes[1] != MimeAny {
		return MimeType{}, MimeTypeWildcardErr{Msg: MimeTypeWildcardErrMsg}
	}

	mt := MimeType{
		Type:    mtypes[0],
		Subtype: mtypes[1],
		Params:  params,
	}

	return mt, nil
}
