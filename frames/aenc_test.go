package frames

import "testing"

func TestAencBasicOutput(t *testing.T) {
	a := NewFrame("AENC", "Audio encrption", Version3)
	if a.GetName() != "AENC" {
		t.Error("Invalid name from AENC frame")
	}
}
