package frames

import "testing"

func TestTextBasicOutput(t *testing.T) {
	x := NewFrame("TEXT", "Lyricist/Text writer", Version3).(*TEXT)

	expected := "TEXT"
	found := x.GetName()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}

	expected = "Lyricist/Text writer"
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

func TestTextParse(t *testing.T) {
	x := NewFrame("TEXT", "Greeting", Version3).(*TEXT)
	b := []byte("\x00Hello, Bob")

	x.ProcessData(len(b), b)
	expected := "[TEXT - 11] (Greeting) Hello, Bob\n"
	found := x.DisplayContent()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}

func TestTextParseUtf16(t *testing.T) {
	x := NewFrame("TEXT", "", Version3).(*TEXT)
	b := []byte("\x01\xfe\xffHello, Bob")

	x.ProcessData(len(b), b)
	expected := "[TEXT - 13] () 䡥汬漬⁂潢\n"
	found := x.DisplayContent()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}

func TestTextUtf16(t *testing.T) {
	x := NewFrame("TIME", "", Version4).(*TEXT)

	expected := "TIME (deprecated)"
	found := x.GetName()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}
