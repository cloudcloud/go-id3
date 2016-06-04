package frames

import "testing"

func TestPrivBasic(t *testing.T) {
	x := NewFrame("PRIV", "Private frame", Version3).(*PRIV)

	expected := "PRIV"
	found := x.GetName()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}

	expected = "Private frame"
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

func TestPrivProcess(t *testing.T) {
	x := NewFrame("PRIV", "", Version3).(*PRIV)
	b := []byte("Bob\x00\x35\x66\x66\x66\x35")

	x.ProcessData(len(b), b)

	expected := "Private owner: Bob\nData: 0x3566666635\n"
	found := x.DisplayContent()
	if found != expected {
		t.Fatalf("Got [%s], Expected [%s]", found, expected)
	}

	b = []byte("Derp\x00Herpa derp")
	x.ProcessData(len(b), b)

	expected = "Private owner: Derp\nData: 0x48657270612064657270\n"
	found = x.DisplayContent()
	if found != expected {
		t.Fatalf("Got [%s], Expected [%s]", found, expected)
	}
}
