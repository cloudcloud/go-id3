package frames

import "testing"

func TestEquaBasicOutput(t *testing.T) {
	e := NewFrame("EQUA", "Equalization", Version3)
	if e.GetName() != "EQUA" {
		t.Error("Invalid name from EQUA frame")
	}

	e = NewFrame("EQUA", "Equalization", Version4)
	if e.GetName() == "EQUA" {
		t.Error("Version 4 EQUA should be deprecated")
	}
}
