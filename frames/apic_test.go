package frames

import "testing"

func TestApicBasicOutput(t *testing.T) {
	x := NewFrame("APIC", "Attached picture", Version3).(*APIC)
	if x.GetName() != "APIC" {
		t.Error("Invalid name from APIC frame")
	}

	if x.GetExplain() != "Attached picture" {
		t.Error("Invalid APIC GetExplain() response")
	}

	x.Name = "BOB"
	if x.GetName() != "BOB" {
		t.Error("Invalid APIC Name setting")
	}
}

func TestApicProcess(t *testing.T) {
	x := NewFrame("APIC", "", Version3).(*APIC)
	b := []byte("\x00image/jpeg\x00\x03Something\x00\x01\x02\x00")

	x.ProcessData(len(b), b)
	expected := 3
	found := x.GetLength()
	if found != expected {
		t.Errorf("Got [%d], Expected [%d]", found, expected)
	}
}

func TestApicProcessUtf16(t *testing.T) {
	x := NewFrame("APIC", "", Version3).(*APIC)
	b := []byte("\x01image/png\x00\x02" +
		"\xfe\xff\x00B\x00o\x00b\x00 \x00i\x00s\x00 \x00G\x00r\x00e\x00a\x00t\x00\x00" +
		"\x01\x02\x03")

	x.ProcessData(len(b), b)
	expected := "Image (image/png, Other file icon, 3b) Bob is Great\n"
	found := x.DisplayContent()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}

func TestApicInvalidPicType(t *testing.T) {
	x := NewFrame("APIC", "", Version3).(*APIC)
	b := []byte("\x00image/png\x00\x22Hello\x00\x01\x02\x03")

	x.ProcessData(len(b), b)
	expected := "Image (image/png, Other, 3b) Hello\n"
	found := x.DisplayContent()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}
