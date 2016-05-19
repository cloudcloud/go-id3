package frames

import "testing"

func TestEncrBasicOutput(t *testing.T) {
	x := NewFrame("ENCR", "Encryption", Version3).(*ENCR)
	if x.GetName() != "ENCR" {
		t.Error("Invalid name from ENCR frame")
	}

	if x.GetExplain() != "Encryption" {
		t.Error("Invalid ENCR GetExplain() response")
	}

	x.Name = "BOB"
	if x.GetName() != "BOB" {
		t.Error("Invalid ENCR Name setting")
	}

	b := []byte("Bob\x00\x81\x01\x02\x03\x02")
	x.ProcessData(len(b), b)

	expected := "Owner: Bob\nMethod: 0"
	if x.DisplayContent() != expected {
		t.Errorf("Invalid DisplayContent() for ENCR, expected '%s' got '%#v'", expected, x.DisplayContent())
	}
}
