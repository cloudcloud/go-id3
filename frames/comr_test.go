package frames

import "testing"

func TestComrBasicOutput(t *testing.T) {
	x := NewFrame("COMR", "Commercial", Version3).(*COMR)
	if x.GetName() != "COMR" {
		t.Error("Invalid name from COMR frame")
	}

	if x.GetExplain() != "Commercial" {
		t.Error("Invalid COMR GetExplain() response")
	}

	x.Name = "BOB"
	if x.GetName() != "BOB" {
		t.Error("Invalid COMR Name setting")
	}
}

func TestComrProcess(t *testing.T) {
	x := NewFrame("COMR", "", Version3).(*COMR)
	b := []byte("\x00aud888.88\x0020200101http://example.com\x00\x00Bob\x00Thing\x00image/jpeg\x00\x01\x02\x02\x01")

	x.ProcessData(len(b), b)

	expected := `Price:           aud888.8
Valid Until:     20200101
Contact URL:     http://example.com
Seller Name:     Bob
Commercial Name: Thing
Mime Type:       image/jpeg`
	found := x.DisplayContent()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}

func TestComrProcessUtf16(t *testing.T) {
	x := NewFrame("COMR", "", Version3).(*COMR)
	b := []byte("\x01aud888.88\x0020200101http://example.com\x00\x00\xfe\xff\x00B\x00o\x00b\x00\x00" +
		"\xfe\xff\x00T\x00h\x00i\x00n\x00g\x00\x00image/jpeg\x00\x01\x02\x02\x01")

	x.ProcessData(len(b), b)

	expected := `Price:           aud888.8
Valid Until:     20200101
Contact URL:     http://example.com
Seller Name:     Bob
Commercial Name: Thing
Mime Type:       image/jpeg`
	found := x.DisplayContent()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}
