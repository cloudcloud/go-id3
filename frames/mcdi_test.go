package frames

import (
	"fmt"
	"testing"
)

func TestMcdiGeneral(t *testing.T) {
	x := NewFrame("MCDI", "Music CD Identifier", Version3).(*MCDI)

	expected := "MCDI"
	found := x.GetName()
	if found != expected {
		t.Fatalf("Expected [%s], Got [%s]", expected, found)
	}

	expected = "Music CD Identifier"
	found = x.GetExplain()
	if found != expected {
		t.Fatalf("Expected [%s], Got [%s]", expected, found)
	}

	x.Name = "BOB"
	expected = "BOB"
	found = x.GetName()
	if found != expected {
		t.Fatalf("Expected [%s], Got [%s]", expected, found)
	}
}

func TestMcdiParseV3(t *testing.T) {
	x := NewFrame("MCDI", "", Version3).(*MCDI)
	b := []byte("\x35\x35\x35\x35" +
		"\x66\x66\x66\x67\x76\x64\x46\x99" +
		"\xa1\x33\x88\x98\x44\x56\x23\x43")
	x.ProcessData(len(b), b)

	expected := fmt.Sprintf("MCD ID (%x)\n\tTrack 0: %x\n\tTrack 1: %x\n",
		"\x35\x35\x35\x35",
		"\x66\x66\x66\x67\x76\x64\x46\x99",
		"\xa1\x33\x88\x98\x44\x56\x23\x43")
	found := x.DisplayContent()
	if found != expected {
		t.Fatalf("Got [%s], Expected [%s]", found, expected)
	}
}
