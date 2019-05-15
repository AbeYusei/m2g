package main

import (
	"errors"
	"testing"
)

func TestErrorHandlingProxy(t *testing.T) {
	errCode := e(errors.New("err"))

	if errCode != 1 {
		t.Fatal("error handler was returned unexpected code.")
	}
}
