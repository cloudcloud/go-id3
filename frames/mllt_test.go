package frames

import "testing"

func TestMlltGeneral(t *testing.T) {
	x := NewFrame("MLLT", "MPEG Location Lookup Table", Version3).(*MLLT)

	expected := "MLLT"
	found := x.GetName()
	if found != expected {
		t.Fatalf("Expected [%s], Got [%s]", expected, found)
	}

	expected = "MPEG Location Lookup Table"
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

func TestMlltParseV3(t *testing.T) {
	x := NewFrame("MLLT", "", Version3).(*MLLT)
	b := []byte("\x00\x01" +
		"\x00\x03\x00" +
		"\x00\x30\x00" +
		"\x02" +
		"\x04" +
		"\x44\x55\x66\x77")
	x.ProcessData(len(b), b)

	expected := "MPEG Lookup\n\tFrames: [0 1]\n\tBytes: [0 3 0]\n\tMilliseconds: [0 48 0]\n"
	found := x.DisplayContent()
	if found != expected {
		t.Fatalf("Got [%s], Expected [%s]", found, expected)
	}
}
