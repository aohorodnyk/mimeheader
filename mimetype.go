package mimeheader

import "mime"

// MimeParts MUST contain two parts <MIME_type>/<MIME_subtype>.
const MimeParts = 2

// MimeSeparator MUST be separated forward slash.
const MimeSeparator = "/"

// MimeAny represented as the asterisk.
const MimeAny = "*"

// MimeType structure for media type (mime type).
type MimeType struct {
	Type    string
	Subtype string
	Params  map[string]string
}

func (mt MimeType) Valid() bool {
	// /plain and text/ are not valid types.
	if mt.Type == "" || mt.Subtype == "" {
		return false
	}

	// */plain is not valid type.
	if mt.Type == MimeAny && mt.Subtype != MimeAny {
		return false
	}

	return true
}

// String builds mime type from type and subtype.
func (mt MimeType) String() string {
	t := mt.Type
	if t == "" {
		t = MimeAny
	}

	st := mt.Subtype
	if st == "" {
		st = MimeAny
	}

	return t + MimeSeparator + st
}

// StringWithParams builds mime type from type and subtype with params.
func (mt MimeType) StringWithParams() string {
	return mime.FormatMediaType(mt.String(), mt.Params)
}

// Match matches current structure with possible wildcards.
// MimeType structure (current) can be wildcard or specific type, like "text/*", "*/*", "text/plain".
func (mt MimeType) Match(target MimeType) bool {
	if !matchMimePart(mt.Type, target.Type) {
		return false
	}

	if !matchMimePart(mt.Subtype, target.Subtype) {
		return false
	}

	return true
}

// MatchText matches current structure with possible wildcards. Target MUST be specific type, like "application/json", "text/plain"
// MimeType structure (current) can be wildcard or specific type, like "text/*", "*/*", "text/plain".
func (mt MimeType) MatchText(target string) bool {
	tmtype, err := ParseMediaType(target)
	if err != nil {
		return false
	}

	return mt.Match(tmtype)
}

func matchMimePart(b, t string) bool {
	if b == t {
		return true
	}

	if b == MimeAny || t == MimeAny {
		return true
	}

	return false
}
