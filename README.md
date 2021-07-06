# Mime header
[![Go (tests and linters)](https://github.com/aohorodnyk/mimeheader/actions/workflows/go.yml/badge.svg)](https://github.com/aohorodnyk/mimeheader/actions/workflows/go.yml) ![GitHub](https://img.shields.io/github/license/aohorodnyk/mimeheader) ![GitHub Workflow Status](https://img.shields.io/github/workflow/status/aohorodnyk/mimeheader/Go) ![GitHub issues](https://img.shields.io/github/issues/aohorodnyk/mimeheader) ![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/aohorodnyk/mimeheader)

## RFC reference
Implementation of mime types and [Accept header support (RFC 2616 Sec 14.1)](https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.1).

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

### Accept header HTTP middleware
[OWASP](https://cheatsheetseries.owasp.org/cheatsheets/REST_Security_Cheat_Sheet.html#send-safe-response-content-types) suggests using this middleware in all applications.
```go
package main

import (
	"context"
	"log"
	"net/http"

	"github.com/aohorodnyk/mimeheader"
)

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/", acceptHeaderMiddleware([]string{"application/json", "text/html"})(handlerTestFunc))

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalln(err)
	}
}

func acceptHeaderMiddleware(acceptMimeTypes []string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(rw http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Accept")
			ah := mimeheader.ParseAcceptHeader(header)

			// We do not need default mime type.
			_, mtype, m := ah.Negotiate(acceptMimeTypes, "")
			if !m {
				// If not matched accept mim type, return 406.
				rw.WriteHeader(http.StatusNotAcceptable)

				return
			}

			// Add matched mime type to context.
			ctx := context.WithValue(r.Context(), "resp_content_type", mtype)
			// New requet from new context.
			rc := r.WithContext(ctx)

			// Call next middleware or handler.
			next(rw, rc)
		}
	}
}

func handlerTestFunc(rw http.ResponseWriter, r *http.Request) {
	mtype := r.Context().Value("resp_content_type").(string)
	rw.Write([]byte(mtype))
}
```

Requests example
```http request
GET http://localhost:8080/
Accept: text/*; q=0.9,application/json; q=1;

# HTTP/1.1 200 OK
# Date: Sat, 03 Jul 2021 19:14:58 GMT
# Content-Length: 16
# Content-Type: text/plain; charset=utf-8
# Connection: close

# application/json

###

GET http://localhost:8080/
Accept: text/*; q=1,application/json; q=1;

# HTTP/1.1 200 OK
# Date: Sat, 03 Jul 2021 19:15:51 GMT
# Content-Length: 16
# Content-Type: text/plain; charset=utf-8
# Connection: close

# application/json

###
GET http://localhost:8080/
Accept: text/html; q=1,application/*; q=1;

# HTTP/1.1 200 OK
# Date: Sat, 03 Jul 2021 19:16:15 GMT
# Content-Length: 9
# Content-Type: text/plain; charset=utf-8
# Connection: close

# text/html

###
GET http://localhost:8080/
Accept: text/*; q=1,application/*; q=0.9;

# HTTP/1.1 200 OK
# Date: Sat, 03 Jul 2021 19:16:48 GMT
# Content-Length: 9
# Content-Type: text/plain; charset=utf-8
# Connection: close

# text/html

###
GET http://localhost:8080/
Accept: text/plain; q=1,application/xml; q=1;

# HTTP/1.1 406 Not Acceptable
# Date: Sat, 03 Jul 2021 19:17:28 GMT
# Content-Length: 0
# Connection: close
```

## Current benchmark results
```
$ go test -bench=.
goos: darwin
goarch: arm64
pkg: github.com/aohorodnyk/mimeheader
BenchmarkParseAcceptHeaderLong-8                        	  595474	      1979 ns/op	    1672 B/op	      16 allocs/op
BenchmarkParseAcceptHeaderThreeWithWights-8             	  879034	      1364 ns/op	    1480 B/op	      14 allocs/op
BenchmarkParseAcceptHeaderOne-8                         	 4268950	       279.7 ns/op	     232 B/op	       7 allocs/op
BenchmarkParseAcceptHeaderAndCompareLong-8              	  295642	      4052 ns/op	    2344 B/op	      34 allocs/op
BenchmarkParseAcceptHeaderAndCompareThreeWithWights-8   	  408138	      2943 ns/op	    1992 B/op	      28 allocs/op
BenchmarkParseAcceptHeaderAndCompareOne-8               	 1399750	       856.0 ns/op	     411 B/op	      13 allocs/op
PASS
ok  	github.com/aohorodnyk/mimeheader	9.252s
```
