package frames

import "testing"

func TestApicBasicOutput(t *testing.T) {
	x := NewFrame("APIC", "Attached picture", Version3)
	if x.GetName() != "APIC" {
		t.Error("Invalid name from APIC frame")
	}
}
