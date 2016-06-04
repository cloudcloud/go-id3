package frames

import "testing"

func TestRvrbBasic(t *testing.T) {
	x := NewFrame("RVRB", "Reverb", Version3).(*RVRB)

	expected := "RVRB"
	found := x.GetName()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}

	expected = "Reverb"
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

func TestRvrbProcess(t *testing.T) {
	x := NewFrame("RVRB", "", Version3).(*RVRB)
	b := []byte("\x00\x20\x00\x20\x03\x03\x7f\x7f\x7f\x7f\x44\x44")

	x.ProcessData(len(b), b)

	expected := "Reverb\n\tReverb Left: 32ms, Reverb Right: 32ms\n\tBounces Left: 3, Bounces Right: 3\n" +
		"\tFeedback Left to Left: 50%, Feedback Left to Right: 50%\n" +
		"\tFeedback Right to Right: 50%, Feedback Right to Left: 50%\n" +
		"\tPremix Left to Right: 27%, Premix Right to Left: 27%\n"
	found := x.DisplayContent()
	if found != expected {
		t.Fatalf("Got [%s], Expected [%s]", found, expected)
	}
}

func TestRvrbFail(t *testing.T) {
	x := NewFrame("RVRB", "", Version3).(*RVRB)
	b := []byte("\x00\x20\x00\x20\x03\x03\x7f\x7f\x7f\x7f")

	x.ProcessData(len(b), b)

	expected := "Reverb frame, no content\n"
	found := x.DisplayContent()
	if found != expected {
		t.Fatalf("Got [%s], Expected [%s]", found, expected)
	}
}
