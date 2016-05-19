package frames

import "testing"

func TestAspiBasic(t *testing.T) {
	x := NewFrame("ASPI", "Audio seek point", Version4).(*ASPI)

	expected := "ASPI"
	found := x.GetName()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}

	expected = "Audio seek point"
	found = x.GetExplain()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}

	expected = "BOB"
	x.Name = "BOB"
	found = x.GetName()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}

func TestAspiProcess(t *testing.T) {
	x := NewFrame("ASPI", "", Version4).(*ASPI)
	b := []byte("\x00\x00\x13\xab\x00\x01\x00\xac\x00\x02\x08" +
		"\x02\x02")

	x.ProcessData(len(b), b)

	expected := "Seek Points (2) [5035:65708]\n"
	found := x.DisplayContent()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}

func TestAspiFailure(t *testing.T) {
	x := NewFrame("ASPI", "", Version4).(*ASPI)
	b := []byte("\x01\x02\x03")

	x.ProcessData(len(b), b)

	expected := "Seek Points (0) [0:0]\n"
	found := x.DisplayContent()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}

func TestAspiAltFailure(t *testing.T) {
	x := NewFrame("ASPI", "", Version4).(*ASPI)
	b := []byte("\x00\x00\x00\x35\x00\x00\x35\x00\x00\x27\x08\x03")
	x.ProcessData(len(b), b)

	expected := "Seek Points (0) [0:0]\n"
	found := x.DisplayContent()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}
