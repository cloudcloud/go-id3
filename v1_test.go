package id3

import "testing"

func TestParseV1(t *testing.T) {
	b := &tfile{}
	b.Write([]byte("TAGBob is great                  " +
		"Bob                           " +
		"Bobbum                        " +
		"2016" +
		"This is just a comment here " +
		"01\x01"))
	v := &V1{Debug: false}

	err := v.Parse(b)
	if err != nil {
		t.Fatalf("Unable to Parse() in V1, [%s]", err)
	}
	if v.Album != "Bobbum" {
		t.Fatalf("Invalid Album found for V1, [%s]", v.Album)
	}
}

func TestParseV1Fail(t *testing.T) {
	b := &tfile{}
	b.Write([]byte("NOPE"))
	v := &V1{Debug: false}

	err := v.Parse(b)
	if err == nil {
		t.Fatalf("Incorrectly Parse() the V1 instead of Fail")
	}
}
