package frames

import "testing"

func TestUsltBasicOutput(t *testing.T) {
	x := NewFrame("USLT", "Unsynchronised test/lyrics transcription", Version3).(*USLT)

	expected := "USLT"
	found := x.GetName()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}

	expected = "Unsynchronised test/lyrics transcription"
	found = x.GetExplain()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}

	x.Name = "BOB"
	expected = "BOB"
	found = x.GetName()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}

func TestUsltParse(t *testing.T) {
	x := NewFrame("USLT", "", Version3).(*USLT)
	b := []byte("\x00engBob\x00Lyrics and stuff")

	x.ProcessData(len(b), b)
	expected := "Unsynchronised Text (eng)\n\t(Bob): Lyrics and stuff\n"
	found := x.DisplayContent()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}

func TestUsltParseUtf16(t *testing.T) {
	x := NewFrame("USLT", "", Version3).(*USLT)
	b := []byte("\x01eng\xfe\xff\x00\x42\x00\x4f\x00\x42\x00\x00\xfe\xff\x00\x42")

	x.ProcessData(len(b), b)
	expected := "Unsynchronised Text (eng)\n\t(BOB): B\n"
	found := x.DisplayContent()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}
