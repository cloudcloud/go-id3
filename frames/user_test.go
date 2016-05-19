package frames

import "testing"

func TestUserBasicOutput(t *testing.T) {
	x := NewFrame("USER", "Terms of Use", Version3).(*USER)

	expected := "USER"
	found := x.GetName()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}

	expected = "Terms of Use"
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

func TestUserParse(t *testing.T) {
	x := NewFrame("USER", "", Version3).(*USER)
	b := []byte("\x00engSome terms here")

	x.ProcessData(len(b), b)
	expected := "Terms of use (eng): Some terms here\n"
	found := x.DisplayContent()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}

func TestUserParseUtf16(t *testing.T) {
	x := NewFrame("USER", "", Version3).(*USER)
	b := []byte("\x01eng\xfe\xff\x00\x42\x00\x4f\x00\x42")

	x.ProcessData(len(b), b)
	expected := "Terms of use (eng): BOB\n"
	found := x.DisplayContent()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}
