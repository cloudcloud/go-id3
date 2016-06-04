package frames

import "testing"

func TestIplsBasicOutput(t *testing.T) {
	i := NewFrame("IPLS", "Involved people list", Version3).(*IPLS)
	if i.GetName() != "IPLS" {
		t.Error("Invalid name from IPLS frame")
	}

	if i.GetExplain() != "Involved people list" {
		t.Error("Invalid IPLS init")
	}

	b := []byte("\x00Producer\x00Bob Down\x00Dancer\x00Bill")
	i.ProcessData(len(b), b)

	expected := 2
	found := len(i.People)
	if found != expected {
		t.Errorf("Got [%d], Expected [%d]", found, expected)
	}
}

func TestIplsUtf16(t *testing.T) {
	i := NewFrame("IPLS", "", Version3).(*IPLS)
	b := []byte("\x01\xfe\xff\x00P\x00r\x00o\x00d\x00u\x00c\x00e\x00r\x00\x00" +
		"\xff\xfe\x00B\x00o\x00b\x00\x00" +
		"\xff\xfe\x00D\x00a\x00n\x00c\x00e\x00r\x00\x00" +
		"\xff\xfe\x00B\x00i\x00l\x00l")
	i.ProcessData(len(b), b)

	expected := 2
	found := len(i.People)

	if found != expected {
		t.Errorf("Got [%d], Expected [%d]", found, expected)
	}
}

func TestIplsVersion4(t *testing.T) {
	i := NewFrame("IPLS", "", Version4).(*IPLS)
	if i.GetName() != "IPLS (deprecated)" {
		t.Error("IPLS is deprecated in Version4")
	}
}

func TestIplsSimpleOutput(t *testing.T) {
	x := NewFrame("IPLS", "", Version3).(*IPLS)
	b := []byte("\x00Producer\x00Bob")

	x.ProcessData(len(b), b)

	expected := "Involved People:\n\tProducer: Bob\n"
	found := x.DisplayContent()
	if found != expected {
		t.Fatalf("Got [%s], Expected [%s]", found, expected)
	}
}
