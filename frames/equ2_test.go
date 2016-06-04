package frames

import "testing"

func TestEqu2BasicOutput(t *testing.T) {
	x := NewFrame("EQU2", "Equalisation (2)", Version4).(*EQU2)

	expected := "EQU2"
	found := x.GetName()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}

	expected = "Equalisation (2)"
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

func TestEqu2Parse(t *testing.T) {
	x := NewFrame("EQU2", "", Version4).(*EQU2)
	b := []byte("\x00Bob\x00" + "\xa3\x01\x02\x00")

	x.ProcessData(len(b), b)
	expected := "Equalisation 2 (Interpolation: Band, Identification: Bob)\n" +
		"\tFrequency: 20864.500000hz, Volume: 1.000000db\n"
	found := x.DisplayContent()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}

func TestEqu2Fail(t *testing.T) {
	x := NewFrame("EQU2", "", Version4).(*EQU2)
	b := []byte("\x01Jim\x00" + "\x00\x20\x06\x23" +
		"\x00\x42\x42\x00")

	x.ProcessData(len(b), b)
	expected := "Equalisation 2 (Interpolation: Linear, Identification: Jim)\n" +
		"\tFrequency: 16.000000hz, Volume: 3.068359db\n" +
		"\tFrequency: 33.000000hz, Volume: 33.000000db\n"
	found := x.DisplayContent()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}
