package frames

import "testing"

func TestRva2BasicOutput(t *testing.T) {
	x := NewFrame("RVA2", "Relative volume adjustment (2)", Version4).(*RVA2)

	expected := "RVA2"
	found := x.GetName()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}

	expected = "Relative volume adjustment (2)"
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

func TestRva2Parse(t *testing.T) {
	x := NewFrame("RVA2", "", Version4).(*RVA2)
	b := []byte("Bob\x00\x01\x02\x0a\x08\x35")

	x.ProcessData(len(b), b)
	expected := "Relative Volume Adjustment (Bob)\n\tChannel (Master volume), Adjusted (1.019531db), Peak (53)\n"
	found := x.DisplayContent()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}

func TestRva2Fixup(t *testing.T) {
	x := NewFrame("RVA2", "", Version4).(*RVA2)
	b := []byte("Jim\x00\x02\x04\x01\x16\x23\x23" +
		"\x29\x06\x00\x08\x12")

	x.ProcessData(len(b), b)
	expected := "Relative Volume Adjustment (Jim)\n\tChannel (Front right), Adjusted (2.001953db), Peak (8995)\n" +
		"\tChannel (Other), Adjusted (3.000000db), Peak (18)\n"
	found := x.DisplayContent()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}
