package frames

import "testing"

func TestWxxxBasicOutput(t *testing.T) {
	x := NewFrame("WXXX", "User defined webpage", Version3).(*WXXX)

	expected := "WXXX"
	found := x.GetName()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}

	expected = "User defined webpage"
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

func TestWxxxParse(t *testing.T) {
	x := NewFrame("WXXX", "User defined webpage", Version3).(*WXXX)
	b := []byte("\x00Bob's home\x00http://example.com")

	x.ProcessData(len(b), b)
	expected := "User webpage: Bob's home [http://example.com]\n"
	found := x.DisplayContent()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}

func TestWxxxParseUtf16(t *testing.T) {
	x := NewFrame("WXXX", "User defined webpage", Version3).(*WXXX)
	b := []byte("\x01\xfe\xff\x00\x42\x00\x6f\x00\x62\x00\x27\x00\x73\x00\x20\x00\x68\x00\x6f\x00\x6d\x00\x65\x00\x00" +
		"http://example.com")

	x.ProcessData(len(b), b)
	expected := "User webpage: Bob's home [http://example.com]\n"
	found := x.DisplayContent()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}
