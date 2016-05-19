package frames

import "testing"

func TestRvadBasic(t *testing.T) {
	x := NewFrame("RVAD", "Relative volume adjustment", Version3).(*RVAD)

	expected := "RVAD"
	found := x.GetName()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}

	expected = "Relative volume adjustment"
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

func TestRvadDeprecated(t *testing.T) {
	x := NewFrame("RVAD", "", Version4).(*RVAD)

	expected := "RVAD (deprecated)"
	found := x.GetName()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}

func TestRvadProcess(t *testing.T) {
	x := NewFrame("RVAD", "", Version3).(*RVAD)
	b := []byte("\x60\x00")

	x.ProcessData(len(b), b)

	expected := "Relative Volume Adjustment\n" +
		"Right\n\tIncrement: false\n\tRelative Volume: 0.000000db\n\tPeak: 0.000000db\n" +
		"Left\n\tIncrement: false\n\tRelative Volume: 0.000000db\n\tPeak: 0.000000db\n" +
		"Right Back\n\tIncrement: false\n\tRelative Volume: 0.000000db\n\tPeak: 0.000000db\n" +
		"Left Back\n\tIncrement: false\n\tRelative Volume: 0.000000db\n\tPeak: 0.000000db\n" +
		"Center\n\tIncrement: false\n\tRelative Volume: 0.000000db\n\tPeak: 0.000000db\n" +
		"Bass\n\tIncrement: false\n\tRelative Volume: 0.000000db\n\tPeak: 0.000000db\n"
	found := x.DisplayContent()
	if found != expected {
		t.Fatalf("Got [%s], Expected [%s]", found, expected)
	}
}

func TestRvadFullProcess(t *testing.T) {
	x := NewFrame("RVAD", "", Version3).(*RVAD)
	b := []byte("\x60\x10\x00\x20\x01\x33\x03\xa0\x03\xa0\xb1\x00\xb1\x00")

	x.ProcessData(len(b), b)
	expected := "Relative Volume Adjustment\n" +
		"Right\n\tIncrement: false\n\tRelative Volume: 0.062500db\n\tPeak: 1.812500db\n" +
		"Left\n\tIncrement: true\n\tRelative Volume: 0.599609db\n\tPeak: 1.812500db\n" +
		"Right Back\n\tIncrement: true\n\tRelative Volume: 88.500000db\n\tPeak: 0.000000db\n" +
		"Left Back\n\tIncrement: false\n\tRelative Volume: 88.500000db\n\tPeak: 0.000000db\n" +
		"Center\n\tIncrement: false\n\tRelative Volume: 0.000000db\n\tPeak: 0.000000db\n" +
		"Bass\n\tIncrement: false\n\tRelative Volume: 0.000000db\n\tPeak: 0.000000db\n"
	found := x.DisplayContent()
	if found != expected {
		t.Fatalf("Got [%s], Expected [%s]", found, expected)
	}
}
