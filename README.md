# Mime header
[![Go (tests and linters)](https://github.com/aohorodnyk/mimeheader/actions/workflows/go.yml/badge.svg)](https://github.com/aohorodnyk/mimeheader/actions/workflows/go.yml)
## Motivation
This library created to help people to parse media type data, like headers, and store and match it.
The main features of the library are to:
* Save a type as a structure (it can help pass mime type as an argument to a function or use as a field of a structure)
* Match mime types by a wildcard. It can be helpful for Accept header negotiation
* MUST be covered by tests to provide guarantees it works well
* Zero dependencies

## Examples
### Match a header with text

```go
package main

import (
	"fmt"

	"github.com/aohorodnyk/mimeheader"
)

// Content-Type - application/json;q=0.9;p=2
func parse(contentType string) {
	mtype, err := mimeheader.ParseMediaType(contentType)

	fmt.Println(err)                                 // nil
	fmt.Println(mtype.MatchText("application/json; param=1")) // true
	fmt.Println(mtype.MatchText("application/xml; param=1"))  // false
	fmt.Println(mtype.MatchText("*/plain; param=1"))          // false
	fmt.Println(mtype.MatchText("text/plain; param=1"))       // false
}
```

### Match accept header

```go
package main

import (
	"fmt"

	"github.com/aohorodnyk/mimeheader"
)

// Accept - application/json;q=1.0,*/*;q=1.0; param=wild,image/png;q=1.0;param=test
func parse(acceptHeader string) {
	ah := mimeheader.ParseAcceptHeader(acceptHeader)

	fmt.Println(ah.Negotiate([]string{"application/json;param=1", "image/png"}, "text/javascript")) // image/png, image/png, true
}
```
