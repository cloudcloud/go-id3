package frames

import (
	"fmt"
	"testing"
)

func TestGridBasicOutput(t *testing.T) {
	x := NewFrame("GRID", "", Version3).(*GRID)
	if x.GetName() != "GRID" {
		t.Error("Invalid name from GRID frame")
	}

	x.Name = "BOB"
	if x.GetName() != "BOB" {
		t.Error("Invalid GRID Name setting")
	}

	b := []byte("Bob\x00\x13\x06\x06\x06")
	x.ProcessData(len(b), b)

	expected := fmt.Sprintf("Group Reg Identifier\n\tOwner: Bob\n\tSymbol: %b\n", '\x13')
	found := x.DisplayContent()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}
