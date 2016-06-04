package frames

import "testing"

func TestEquaBasicOutput(t *testing.T) {
	e := NewFrame("EQUA", "Equalization", Version3).(*EQUA)
	if e.GetName() != "EQUA" {
		t.Error("Invalid name from EQUA frame")
	}

	e = NewFrame("EQUA", "Equalization", Version4).(*EQUA)
	if e.GetName() == "EQUA" {
		t.Error("Version 4 EQUA should be deprecated")
	}

	b := []byte("\x10\x01\x34\x11\x22")
	e.ProcessData(len(b), b)
	if e.Adjustment != 16 {
		t.Fatalf("Invalid Adjustment for EQUA, got '%d'", e.Adjustment)
	}
	if len(e.Steps) != 1 {
		t.Fatalf("Expected 1 step for EQUA")
	}

	expected := "Adjustment: 16\nSteps: 1"
	if e.DisplayContent() != expected {
		t.Fatal("DisplayContent() incorrect for EQUA")
	}
}
