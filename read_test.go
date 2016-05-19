package main

import "testing"

func TestReadInit(t *testing.T) {
	r := readCmd
	if r.UsageLine != "read [filename]" {
		t.Error("Expected UsageLine to be simple")
	}

	if len(r.Short) < 3 {
		t.Error("Short should be of a decent length")
	}
}

func TestReadBaseProcess(t *testing.T) {
	b := &tfile{}
	readProcess([]string{}, b)

	found := b.buf.String()
	expected := "No filename provided"
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}
