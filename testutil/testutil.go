package testutil

import (
	"testing"
)

// Errorf output log with formatted
func Errorf(tb testing.TB, want, got int) {
	tb.Errorf("want = %d, got = %d", want, got)
}

// ErrorfHelper output log with formatted
func ErrorfHelper(tb testing.TB, want, got int) {
	tb.Helper()
	tb.Errorf("want = %d, got = %d", want, got)
}
