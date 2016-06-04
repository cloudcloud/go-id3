package frames

import "testing"

func TestOwneGeneral(t *testing.T) {
	x := NewFrame("OWNE", "Ownership", Version3).(*OWNE)

	expected := "OWNE"
	found := x.GetName()
	if found != expected {
		t.Fatalf("Expected [%s], Got [%s]", expected, found)
	}

	expected = "Ownership"
	found = x.GetExplain()
	if found != expected {
		t.Fatalf("Expected [%s], Got [%s]", expected, found)
	}

	x.Name = "BOB"
	expected = "BOB"
	found = x.GetName()
	if found != expected {
		t.Fatalf("Expected [%s], Got [%s]", expected, found)
	}
}

func TestOwneParseV3(t *testing.T) {
	x := NewFrame("OWNE", "", Version3).(*OWNE)
	b := []byte("\x00aud666.66\x0020160529Bob")
	x.ProcessData(len(b), b)

	expected := "Ownership\n\tCurrency: aud\n\tPayed: 666.66\n\tDate: 20160529\n\tSeller: Bob\n"
	found := x.DisplayContent()
	if found != expected {
		t.Fatalf("Got [%s], Expected [%s]", found, expected)
	}
}

func TestOwneParseUtf16(t *testing.T) {
	x := NewFrame("OWNE", "", Version4).(*OWNE)
	b := []byte("\x01aud666.66\x0020160529\xfe\xff B o b")
	x.ProcessData(len(b), b)

	expected := "Ownership\n\tCurrency: aud\n\tPayed: 666.66\n\tDate: 20160529\n\tSeller: ⁂⁯⁢\n"
	found := x.DisplayContent()
	if found != expected {
		t.Fatalf("Got [%s], Expected [%s]", found, expected)
	}
}
