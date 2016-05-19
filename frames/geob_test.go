package frames

import "testing"

func TestGeobBasicOutput(t *testing.T) {
	g := NewFrame("GEOB", "General encapsulated object", Version3).(*GEOB)
	if g.GetName() != "GEOB" {
		t.Error("Invalid name from GEOB frame")
	}

	if g.GetExplain() != "General encapsulated object" {
		t.Error("Invalid GEOB init")
	}

	b := []byte("\x00image/jpeg\x00bob.jpg\x00Bob\x00\x01\x03\x02")
	g.ProcessData(len(b), b)
	expected := "Mime Type:   image/jpeg\nFilename:    bob.jpg\nDescription: Bob"
	if g.DisplayContent() != expected {
		t.Errorf("Invalid DisplayContent() for GEOB, [%s] vs. [%s]", expected, g.DisplayContent())
	}
}

func TestGeobUtf16(t *testing.T) {
	g := NewFrame("GEOB", "", Version3).(*GEOB)
	b := []byte("\x01image/jpg\x00" +
		"\xfe\xff\x00b\x00o\x00b\x00.\x00j\x00p\x00g\x00\x00" +
		"\xfe\xff\x00B\x00o\x00b\x00\x00" +
		"\x01\x03\x02")
	g.ProcessData(len(b), b)

	expected := "Mime Type:   image/jpg\nFilename:    bob.jpg\nDescription: Bob"
	found := g.DisplayContent()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}
