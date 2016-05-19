package frames

import "testing"

func TestAencBasicOutput(t *testing.T) {
	a := NewFrame("AENC", "Audio encryption", Version3).(*AENC)
	if a.GetName() != "AENC" {
		t.Error("Invalid name from AENC frame")
	}

	if a.GetExplain() != "Audio encryption" {
		t.Errorf("Invalid AENC explanation, got %#v expected %#v", a.GetExplain(), "Audio encryption")
	}

	a.Name = "BOB"
	if a.GetName() != "BOB" {
		t.Error("Invalid AENC Name setting")
	}

	b := []byte("Owner Bob\x000411\x01\x02\x00")
	a.ProcessData(len(b), b)
	if a.GetLength() != len(b) {
		t.Error("Invalid AENC ProcessData() result [Size]")
	}
	if a.Contact != "Owner Bob" {
		t.Error("Invalid AENC ProcessData() result [Contact]")
	}

	out := "Contact: Owner Bob\nPreviewStart: []byte{0x30, 0x34}\n" +
		"PreviewLength: []byte{0x31, 0x31}\nEncryption: []byte{0x1, 0x2, 0x0}\n"
	if a.DisplayContent() != out {
		t.Error("Invalid AENC DisplayContent() output")
	}

	b = []byte("\x00\x00")
	a.ProcessData(len(b), b)
	if a.Size != len(b) {
		t.Error("Invalid AENC ProcessData() result")
	}
}
