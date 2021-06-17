# Mime header

## Motivation
This library created to help people to parse media type data, like headers, and store and match it.
The main features of the library are to:
* Save a type as a structure (it can help pass mime type as an argument to a function or use as a field of a structure)
* Match mime types by a wildcard. It can be helpful for Accept header negotiation

Library MUST be covered by tests to provide guarantees it works well.

## Examples
### Match a header with text

```go
package main

import (
	"fmt"

	"github.com/aohorodnyk/mimeheader"
)

// acceptHeader contains only 1 value - application/*;q=0.9
func parse(acceptHeader string) {
	mtype, err := mimeheader.ParseMediaType(acceptHeader)

	fmt.Println(err)                                 // nil
	fmt.Println(mtype.MatchText("application/json; param=1")) // true
	fmt.Println(mtype.MatchText("application/xml; param=1"))  // true
	fmt.Println(mtype.MatchText("*/plain; param=1"))          // true
	fmt.Println(mtype.MatchText("text/plain; param=1"))       // false
}
```
