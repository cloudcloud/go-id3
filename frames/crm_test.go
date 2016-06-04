package frames

import "testing"

func TestBaseCrm(t *testing.T) {
	x := NewFrame("CRM", "Encrypted meta frame", Version2).(*CRM)

	found := x.GetName()
	expected := "CRM"
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}

	found = x.GetExplain()
	expected = "Encrypted meta frame"
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}

	x.Name = "BOB"
	found = x.GetName()
	expected = "BOB"
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}

func TestCrmParse(t *testing.T) {
	x := NewFrame("CRM", "", Version2).(*CRM)
	b := []byte("\x00Bob\x00Home\x00\x34\x52\x42")

	x.ProcessData(len(b), b)

	found := x.DisplayContent()
	expected := "Encryption Meta\n\tOwner: Bob\n\tExplanation: Home\n"
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}

func TestCrmParseUtf16(t *testing.T) {
	x := NewFrame("CRM", "", Version2).(*CRM)
	b := []byte("\x01\xfe\xff\x00B\x00o\x00b\x00\x00\xfe\xff\x00H\x00o\x00m\x00e\x00\x00" +
		"\x42\x42\x42\x42")

	x.ProcessData(len(b), b)

	found := x.DisplayContent()
	expected := "Encryption Meta\n\tOwner: Bob\n\tExplanation: Home\n"
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}
