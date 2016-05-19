package frames

import "testing"

func TestPcntBasic(t *testing.T) {
	x := NewFrame("PCNT", "Play counter", Version3).(*PCNT)

	expected := "PCNT"
	found := x.GetName()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}

	expected = "Play counter"
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

func TestPcntProcess(t *testing.T) {
	x := NewFrame("PCNT", "", Version3).(*PCNT)
	b := []byte("\x00\x00\x00\x05")

	x.ProcessData(len(b), b)

	expected := "Count: 5\n"
	found := x.DisplayContent()
	if found != expected {
		t.Fatalf("Got [%s], Expected [%s]", found, expected)
	}

	b = []byte("\x00\x00\x05\x20")
	x.ProcessData(len(b), b)

	expected = "Count: 1312\n"
	found = x.DisplayContent()
	if found != expected {
		t.Fatalf("Got [%s], Expected [%s]", found, expected)
	}
}
