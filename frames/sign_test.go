package frames

import "testing"

func TestSignBasicOutput(t *testing.T) {
	x := NewFrame("SIGN", "Signature", Version4).(*SIGN)

	expected := "SIGN"
	found := x.GetName()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}

	expected = "Signature"
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

func TestSignParse(t *testing.T) {
	x := NewFrame("SIGN", "", Version4).(*SIGN)
	b := []byte("\x00\x00\x00\x20")

	x.ProcessData(len(b), b)
	expected := "Signature (0): 000020\n"
	found := x.DisplayContent()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}

func TestSignFail(t *testing.T) {
	x := NewFrame("SIGN", "", Version4).(*SIGN)
	b := []byte("\x00\x00")

	x.ProcessData(len(b), b)
	expected := "Signature (0): 00\n"
	found := x.DisplayContent()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}
