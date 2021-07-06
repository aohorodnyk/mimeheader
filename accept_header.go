package mimeheader

import "sort"

type MimeHeader struct {
	MimeType
	Quality float32
}

type AcceptHeader struct {
	mheaders []MimeHeader
}

func NewAcceptHeaderPlain(mheaders []MimeHeader) AcceptHeader {
	return AcceptHeader{mheaders: mheaders}
}

func NewAcceptHeader(mheaders []MimeHeader) AcceptHeader {
	ah := AcceptHeader{mheaders: mheaders}
	ah.sort()

	return ah
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

	less, done := ah.lessWildcard(i, j)
	if done {
		return less
	}

	return ah.lessParams(i, j)
}

func (ah AcceptHeader) lessParams(i, j int) bool {
	li := len(ah.mheaders[i].Params)
	_, ok := ah.mheaders[i].Params["q"]

	if ok {
		li--
	}

	lj := len(ah.mheaders[j].Params)
	_, ok = ah.mheaders[j].Params["q"]

	if ok {
		lj--
	}

	return li < lj
}

func (ah AcceptHeader) lessWildcard(i, j int) (less, done bool) {
	// '*' value has less priority than a specific type
	// If i contains '*' and j has specific type, then i less than j
	if ah.mheaders[i].Type == MimeAny && ah.mheaders[j].Type != MimeAny {
		return true, true
	}

	// If i contains a specific type and j contains '*' then i greater than j
	if ah.mheaders[i].Type != MimeAny && ah.mheaders[j].Type == MimeAny {
		return false, true
	}

	// '*' value has less priority than a specific type
	// If i contains '*' and j has specific type, then i less than j
	if ah.mheaders[i].Subtype == MimeAny && ah.mheaders[j].Subtype != MimeAny {
		return true, true
	}

	// If i contains a specific type and j contains '*' then i greater than j
	if ah.mheaders[i].Subtype != MimeAny && ah.mheaders[j].Subtype == MimeAny {
		return false, true
	}

	return false, false
}

// Swap function for sort.Interface interface.
func (ah *AcceptHeader) Swap(i, j int) {
	ah.mheaders[i], ah.mheaders[j] = ah.mheaders[j], ah.mheaders[i]
}

// Add mime header to accept header.
// MimeHeader will be validated and added ONLY if valid.
// AcceptHeader will be sorted.
// For performance reasons better to use Set, instead of Add.
func (ah *AcceptHeader) Add(mh MimeHeader) {
	if !mh.Valid() {
		return
	}

	ah.mheaders = append(ah.mheaders, mh)

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

	ah.mheaders = mheaders

	ah.sort()
}

// Negotiate return appropriate type fot current accept list from supported (common) mime types.
// First parameter returns matched value from accept header.
// Second parameter returns matched common type.
// Third parameter returns matched common type or default type applied.
func (ah AcceptHeader) Negotiate(ctypes []string, dtype string) (accept MimeHeader, mimeType string, matched bool) {
	if len(ctypes) == 0 || len(ah.mheaders) == 0 {
		return MimeHeader{}, dtype, false
	}

	var parsedCType MimeType

	mhid := -1

	for _, ctype := range ctypes {
		for hid, header := range ah.mheaders {
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
		return ah.mheaders[mhid], parsedCType.String(), true
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
