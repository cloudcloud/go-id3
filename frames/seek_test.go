package frames

import "testing"

func TestSeekBasicOutput(t *testing.T) {
	x := NewFrame("SEEK", "Seek point", Version4).(*SEEK)

	expected := "SEEK"
	found := x.GetName()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}

	expected = "Seek point"
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

func TestSeekParse(t *testing.T) {
	x := NewFrame("SEEK", "", Version4).(*SEEK)
	b := []byte("\x00\x00\x00\x20")

	x.ProcessData(len(b), b)
	expected := "Seek point: 32\n"
	found := x.DisplayContent()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}

func TestSeekFail(t *testing.T) {
	x := NewFrame("SEEK", "", Version4).(*SEEK)
	b := []byte("\x00\x00")

	x.ProcessData(len(b), b)
	expected := "Seek point: 0\n"
	found := x.DisplayContent()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}
