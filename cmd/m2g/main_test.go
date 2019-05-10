package main

import (
	"testing"
)

func example(str string) (bool, error) {
	return false, nil
}
func TestExampleSuccess(t *testing.T) {
	result, err := example("hoge")

	if err != nil {
		t.Fatal("filed: errored.")
	}

	if !result {
		t.Fatal("filed: false")
	}
}
