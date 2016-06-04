package frames

import "testing"

func TestPossBasic(t *testing.T) {
	x := NewFrame("POSS", "Position synchronisation", Version3).(*POSS)

	expected := "POSS"
	found := x.GetName()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}

	expected = "Position synchronisation"
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

func TestPossProcess(t *testing.T) {
	x := NewFrame("POSS", "", Version3).(*POSS)
	b := []byte("\x01\x22\x01\x21\x03")

	x.ProcessData(len(b), b)

	expected := "Position synchronisation\n\tFormat: MPEG\n\tPositions: 0x22012103\n"
	found := x.DisplayContent()
	if found != expected {
		t.Fatalf("Got [%s], Expected [%s]", found, expected)
	}

	b = []byte("\x02\x22\x01\x21\x03")
	x.ProcessData(len(b), b)

	expected = "Position synchronisation\n\tFormat: Milliseconds\n\tPositions: 0x22012103\n"
	found = x.DisplayContent()
	if found != expected {
		t.Fatalf("Got [%s], Expected [%s]", found, expected)
	}
}
