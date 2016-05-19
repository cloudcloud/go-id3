package frames

import "testing"

func TestTxxxBasicOutput(t *testing.T) {
	x := NewFrame("TXXX", "User text", Version3).(*TXXX)

	expected := "TXXX"
	found := x.GetName()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}

	expected = "User text"
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

func TestTxxxParse(t *testing.T) {
	x := NewFrame("TXXX", "Greeting", Version3).(*TXXX)
	b := []byte("\x00Hello\x00Bob")

	x.ProcessData(len(b), b)
	expected := "User text (Hello):(Bob)\n"
	found := x.DisplayContent()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}

func TestTxxxParseUtf16(t *testing.T) {
	x := NewFrame("TXXX", "", Version3).(*TXXX)
	b := []byte("\x01\xfe\xffHello \x00\x00\xfe\xffBob ")

	x.ProcessData(len(b), b)
	expected := "User text (䡥汬漠):(䉯戠)\n"
	found := x.DisplayContent()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}
