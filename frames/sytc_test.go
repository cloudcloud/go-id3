package frames

import "testing"

func TestSytcBasic(t *testing.T) {
	x := NewFrame("SYTC", "Synchronised Tempo Codes", Version3).(*SYTC)

	expected := "SYTC"
	found := x.GetName()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}

	expected = "Synchronised Tempo Codes"
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

func TestSytcProcess(t *testing.T) {
	x := NewFrame("SYTC", "", Version3).(*SYTC)
	b := []byte("\x01\xff\x23\x00\x01\x13" +
		"\x00\x00\x02\x01\xff\x04\x00\x02\xa1")

	x.ProcessData(len(b), b)

	expected := "Synchronised Tempo\n" +
		"\tBPM [290] Time Code [275mpeg]\n" +
		"\tBPM [0] Time Code [513mpeg]\n" +
		"\tBPM [259] Time Code [673mpeg]\n"
	found := x.DisplayContent()
	if found != expected {
		t.Fatalf("Got [%s], Expected [%s]", found, expected)
	}
}
