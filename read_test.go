package main

import "testing"

func TestReadInit(t *testing.T) {
	r := readCmd
	if r.UsageLine != "read [options] [filename]" {
		t.Error("Expected UsageLine to be simple")
	}
}
