package mimeheader

import "sort"

type MimeHeader struct {
	MimeType
	Quality float32
}

type AcceptHeader struct {
	mheaders []MimeHeader
}

func NewAcceptHeader(mheaders []MimeHeader) AcceptHeader {
	return AcceptHeader{mheaders: mheaders}
}

// Len function for sort.Interface interface.
func (ah AcceptHeader) Len() int {
	return len(ah.mheaders)
}

// Less function for sort.Interface interface.
func (ah AcceptHeader) Less(i, j int) bool {
	// Sort accepts by quality value. If quality is not equal, then return is i less than j
	if ah.mheaders[i].Quality != ah.mheaders[j].Quality {
		return ah.mheaders[i].Quality < ah.mheaders[j].Quality
	}

	// '*' value has less priority than a specific type
	// If i contains '*' and j has specific type, then i less than j
	if ah.mheaders[i].Type == MimeAny && ah.mheaders[j].Type != MimeAny {
		return true
	}

	// If i contains a specific type and j contains '*' then i greater than j
	if ah.mheaders[i].Type != MimeAny && ah.mheaders[j].Type == MimeAny {
		return false
	}

	// '*' value has less priority than a specific type
	// If i contains '*' and j has specific type, then i less than j
	if ah.mheaders[i].Subtype == MimeAny && ah.mheaders[j].Subtype != MimeAny {
		return true
	}

	// If i contains a specific type and j contains '*' then i greater than j
	if ah.mheaders[i].Subtype != MimeAny && ah.mheaders[j].Subtype == MimeAny {
		return false
	}

	// More specific params has greater priority
	return len(ah.mheaders[i].Params) < len(ah.mheaders[j].Params)
}

// Swap function for sort.Interface interface.
func (ah *AcceptHeader) Swap(i, j int) {
	ah.mheaders[i], ah.mheaders[j] = ah.mheaders[j], ah.mheaders[i]
}

func (ah *AcceptHeader) Add(mh MimeHeader) {
	ah.mheaders = append(ah.mheaders, mh)

	ah.sort()
}

func (ah *AcceptHeader) Set(mhs []MimeHeader) {
	ah.mheaders = mhs

	ah.sort()
}

// Negotiate return appropriate type fot current accept list from supported (common) mime types.
// Second parameter shows matched common type or default type applied.
func (ah AcceptHeader) Negotiate(ctypes []string, dtype string) (header, mimType string, matched bool) {
	if len(ctypes) == 0 || len(ah.mheaders) == 0 {
		return "", dtype, false
	}

	for _, ctype := range ctypes {
		for _, header := range ah.mheaders {
			mtype, err := ParseMediaType(ctype)
			if err != nil {
				continue
			}

			if header.Match(mtype) {
				return header.String(), mtype.String(), true
			}
		}
	}

	return "", dtype, false
}

func (ah AcceptHeader) Match(mtype string) bool {
	_, _, matched := ah.Negotiate([]string{mtype}, "")

	return matched
}

func (ah *AcceptHeader) sort() {
	sort.Sort(sort.Reverse(ah))
}
