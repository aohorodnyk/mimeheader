package mimeheader

import (
	"strconv"
	"strings"
)

const DefaultQuality float32 = 1.0

func ParseAcceptHeader(header string) AcceptHeader {
	accepts := strings.Split(header, ",")
	if len(accepts) == 0 {
		return AcceptHeader{}
	}

	mheaders := make([]MimeHeader, 0, len(accepts))

	for _, accept := range accepts {
		mtype, err := ParseMediaType(accept)
		if err != nil {
			continue
		}

		header := MimeHeader{
			MimeType: mtype,
			Quality:  DefaultQuality,
		}

		if qs, ok := header.Params["q"]; ok {
			const floatSize = 32
			quality, err := strconv.ParseFloat(qs, floatSize)

			if err == nil {
				header.Quality = float32(quality)
			}
		}

		mheaders = append(mheaders, header)
	}

	ah := AcceptHeader{}
	ah.Set(mheaders)

	return ah
}
