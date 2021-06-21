package mimeheader

type MimeParseErr struct {
	Err error
	Msg string
}

func (e MimeParseErr) Error() string {
	if e.Err == nil {
		return e.Msg
	}

	return e.Msg + ": " + e.Err.Error()
}

func (e MimeParseErr) Unwrap() error {
	return e.Err
}

type MimeTypePartsErr struct {
	Msg string
}

func (e MimeTypePartsErr) Error() string {
	return e.Msg
}

type MimeTypeWildcardErr struct {
	Msg string
}

func (e MimeTypeWildcardErr) Error() string {
	return e.Msg
}
