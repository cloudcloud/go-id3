package frames

import "testing"

func TestSyltBasic(t *testing.T) {
	x := NewFrame("SYLT", "Synchronised text", Version3).(*SYLT)

	expected := "SYLT"
	found := x.GetName()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}

	expected = "Synchronised text"
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

func TestSyltProcess(t *testing.T) {
	x := NewFrame("SYLT", "", Version3).(*SYLT)
	b := []byte("\x00eng\x02\x01Lyrics\x00" +
		"Bob\x00\x00\x35" +
		"Down\x00\x01\x56")

	x.ProcessData(len(b), b)

	expected := "Synchronised (Lyrics). Language(eng) Format(ms) Content Type(Lyrics)\n" +
		"\tBob [53]\n\tDown [342]\n"
	found := x.DisplayContent()
	if found != expected {
		t.Fatalf("Got [%s], Expected [%s]", found, expected)
	}
}

func TestSyltUtf16(t *testing.T) {
	x := NewFrame("SYLT", "", Version4).(*SYLT)
	b := []byte("\x01eng\x01\x08\xfe\xffDerp\x00\x00" +
		"XYZ\x00\x03\x44" +
		"ABC\x00\x00\x20")

	x.ProcessData(len(b), b)

	expected := "Synchronised (䑥牰). Language(eng) Format(mpeg) Content Type(Other)\n" +
		"\tXYZ [836]\n\tABC [32]\n"
	found := x.DisplayContent()
	if found != expected {
		t.Fatalf("Got [%s], Expected [%s]", found, expected)
	}
}
