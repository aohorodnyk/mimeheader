package mimeheader

import "sort"

type MimeHeader struct {
	MimeType
	Quality float32
}

type AcceptHeader struct {
	MHeaders []MimeHeader
}

func NewAcceptHeaderPlain(mheaders []MimeHeader) AcceptHeader {
	return AcceptHeader{MHeaders: mheaders}
}

func NewAcceptHeader(mheaders []MimeHeader) AcceptHeader {
	ah := AcceptHeader{MHeaders: mheaders}
	ah.sort()

	return ah
}

// Len function for sort.Interface interface.
func (ah AcceptHeader) Len() int {
	return len(ah.MHeaders)
}

// Less function for sort.Interface interface.
func (ah AcceptHeader) Less(i, j int) bool {
	// Sort accepts by quality value. If quality is not equal, then return is i less than j
	if ah.MHeaders[i].Quality != ah.MHeaders[j].Quality {
		return ah.MHeaders[i].Quality < ah.MHeaders[j].Quality
	}

	less, done := ah.lessWildcard(i, j)
	if done {
		return less
	}

	return ah.lessParams(i, j)
}

func (ah AcceptHeader) lessParams(i, j int) bool {
	li := len(ah.MHeaders[i].Params)
	_, ok := ah.MHeaders[i].Params["q"]

	if ok {
		li--
	}

	lj := len(ah.MHeaders[j].Params)
	_, ok = ah.MHeaders[j].Params["q"]

	if ok {
		lj--
	}

	return li < lj
}

func (ah AcceptHeader) lessWildcard(i, j int) (less, done bool) {
	// '*' value has less priority than a specific type
	// If i contains '*' and j has specific type, then i less than j
	if ah.MHeaders[i].Type == MimeAny && ah.MHeaders[j].Type != MimeAny {
		return true, true
	}

	// If i contains a specific type and j contains '*' then i greater than j
	if ah.MHeaders[i].Type != MimeAny && ah.MHeaders[j].Type == MimeAny {
		return false, true
	}

	// '*' value has less priority than a specific type
	// If i contains '*' and j has specific type, then i less than j
	if ah.MHeaders[i].Subtype == MimeAny && ah.MHeaders[j].Subtype != MimeAny {
		return true, true
	}

	// If i contains a specific type and j contains '*' then i greater than j
	if ah.MHeaders[i].Subtype != MimeAny && ah.MHeaders[j].Subtype == MimeAny {
		return false, true
	}

	return false, false
}

// Swap function for sort.Interface interface.
func (ah *AcceptHeader) Swap(i, j int) {
	ah.MHeaders[i], ah.MHeaders[j] = ah.MHeaders[j], ah.MHeaders[i]
}

// Add mime header to accept header.
// MimeHeader will be validated and added ONLY if valid.
// AcceptHeader will be sorted.
// For performance reasons better to use Set, instead of Add.
func (ah *AcceptHeader) Add(mh MimeHeader) {
	if !mh.Valid() {
		return
	}

	ah.MHeaders = append(ah.MHeaders, mh)

	ah.sort()
}

// Set all valid headers to AcceprHeader (override old ones).
// Sorting will be applied.
func (ah *AcceptHeader) Set(mhs []MimeHeader) {
	mheaders := make([]MimeHeader, 0, len(mhs))

	for _, mh := range mhs {
		if mh.Valid() {
			mheaders = append(mheaders, mh)
		}
	}

	ah.MHeaders = mheaders

	ah.sort()
}

// Negotiate return appropriate type fot current accept list from supported (common) mime types.
// First parameter returns matched value from accept header.
// Second parameter returns matched common type.
// Third parameter returns matched common type or default type applied.
func (ah AcceptHeader) Negotiate(ctypes []string, dtype string) (accept MimeHeader, mimeType string, matched bool) {
	if len(ctypes) == 0 || len(ah.MHeaders) == 0 {
		return MimeHeader{}, dtype, false
	}

	var parsedCType MimeType

	mhid := -1

	for _, ctype := range ctypes {
		for hid, header := range ah.MHeaders {
			mtype, err := ParseMediaType(ctype)
			if err != nil {
				continue
			}

			if header.Match(mtype) && (mhid > hid || mhid < 0) {
				parsedCType = mtype
				mhid = hid
			}
		}
	}

	if mhid >= 0 {
		return ah.MHeaders[mhid], parsedCType.String(), true
	}

	return MimeHeader{}, dtype, false
}

// Match is the same function as AcceptHeader.Negotiate.
// It implements simplified interface to match only one type and return only matched or not information.
func (ah AcceptHeader) Match(mtype string) bool {
	_, _, matched := ah.Negotiate([]string{mtype}, "")

	return matched
}

func (ah *AcceptHeader) sort() {
	sort.Sort(sort.Reverse(ah))
}
