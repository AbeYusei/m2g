package testutil

import (
	"testing"
)

// Errorf output log with formatted
func Errorf(tb testing.TB, want, got interface{}) {
	tb.Errorf("want = %v, got = %v", want, got)
}

// ErrorfHelper output log with formatted
func ErrorfHelper(tb testing.TB, want, got interface{}) {
	tb.Helper()
	tb.Errorf("want = %v, got = %v", want, got)
}
