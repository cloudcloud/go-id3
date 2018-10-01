package frames

import "testing"

func TestUfidBasicOutput(t *testing.T) {
	x := NewFrame("UFID", "Unique file identifier", Version3).(*UFID)

	expected := "UFID"
	found := x.GetName()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}

	expected = "Unique file identifier"
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

func TestUfidParse(t *testing.T) {
	x := NewFrame("UFID", "", Version3).(*UFID)
	b := []byte("Bob\x00xyz")

	x.ProcessData(len(b), b)
	expected := "Owner: (Bob) Identifier: (78797a)\n"
	found := x.DisplayContent()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}
