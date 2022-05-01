package main

import "testing"

func TestGetFile(t *testing.T) {
	total, err := getFileContent("./.envrc")
	if err != nil || total[0] == "" {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", err, 10)
	}
}
