package frames

import "testing"

func TestEtcoBasicOutput(t *testing.T) {
	e := NewFrame("ETCO", "Event timing codes", Version3).(*ETCO)
	if e.GetName() != "ETCO" {
		t.Error("Invalid name from ETCO frame")
	}

	if e.GetExplain() != "Event timing codes" {
		t.Error("Invalid ETCO init")
	}

	b := []byte("\x01\x02\x01\x00\x03\xa4")
	e.ProcessData(len(b), b)
	if e.DisplayContent() != "Format: 1\nCode Count: 1" {
		t.Error("Invalid DisplayContent() for ETCO")
	}
}
