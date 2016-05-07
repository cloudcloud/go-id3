package frames

import "testing"

func TestTextBasicOutput(t *testing.T) {
	x := NewFrame("TEXT", "Lyricist/Text writer", Version3)
	if x.GetName() != "TEXT" {
		t.Error("Invalid name from TEXT frame")
	}
}
