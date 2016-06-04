package frames

import "testing"

func TestRbufBasic(t *testing.T) {
	x := NewFrame("RBUF", "Recommended buffer", Version3).(*RBUF)

	expected := "RBUF"
	found := x.GetName()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}

	expected = "Recommended buffer"
	found = x.GetExplain()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}

	x.Name = "BOB"
	expected = "BOB"
	found = x.GetName()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}

func TestRbufProcess(t *testing.T) {
	x := NewFrame("RBUF", "", Version3).(*RBUF)
	b := []byte("\x00\x01\x34\x00\x00\x00\x01\x34")

	x.ProcessData(len(b), b)

	expected := "Recommended buffer\n\tSize: 308\n\tInfo: false\n\tOffset: 308\n"
	found := x.DisplayContent()
	if found != expected {
		t.Fatalf("Got [%s], Expected [%s]", found, expected)
	}
}

func TestRbufVersion2(t *testing.T) {
	x := NewFrame("BUF", "Recommended buffer size", Version2).(*RBUF)
	b := []byte("\x00\x00\x10\x00\x00\x00\x10\x00")

	expected := "BUF"
	found := x.GetName()
	if found != expected {
		t.Fatalf("Got [%s], Expected [%s]", found, expected)
	}

	x.ProcessData(len(b), b)
}
