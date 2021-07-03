package mimeheader_test

import (
	"testing"

	"github.com/aohorodnyk/mimeheader"
)

func BenchmarkParseAcceptHeaderLong(b *testing.B) {
	benchmarkParseAcceptHeader(b, "*/*; q=0.9; s=1, image/*; q=0.9; s=4, application/json; q=0.9; b=3;, text/plain")
}

func BenchmarkParseAcceptHeaderThreeWithWights(b *testing.B) {
	benchmarkParseAcceptHeader(b, "*/*; q=0.9;, image/*; q=0.9, application/json; q=0.9")
}

func BenchmarkParseAcceptHeaderOne(b *testing.B) {
	benchmarkParseAcceptHeader(b, "*/*")
}

func BenchmarkParseAcceptHeaderAndCompareLong(b *testing.B) {
	header := "*/*; q=0.9; s=1, image/*; q=0.9; s=4, application/json; q=0.9; b=3;, text/plain"
	negotiate := []string{"application/json", "application/xml"}

	benchmarkParseAcceptHeaderAndCompare(b, header, negotiate)
}

func BenchmarkParseAcceptHeaderAndCompareThreeWithWights(b *testing.B) {
	header := "*/*; q=0.9;, image/*; q=0.9, application/json; q=0.9"
	negotiate := []string{"application/json", "application/xml"}

	benchmarkParseAcceptHeaderAndCompare(b, header, negotiate)
}

func BenchmarkParseAcceptHeaderAndCompareOne(b *testing.B) {
	header := "*/*"
	negotiate := []string{"application/json", "application/xml"}

	benchmarkParseAcceptHeaderAndCompare(b, header, negotiate)
}

func benchmarkParseAcceptHeaderAndCompare(b *testing.B, header string, negotiate []string) {
	b.Helper()

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		ah := mimeheader.ParseAcceptHeader(header)

		ah.Negotiate(negotiate, "")
	}
}

func benchmarkParseAcceptHeader(b *testing.B, header string) {
	b.Helper()

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		mimeheader.ParseAcceptHeader(header)
	}
}
