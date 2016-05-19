package frames

import "testing"

func TestPopmBasic(t *testing.T) {
	x := NewFrame("POPM", "Popularimeter", Version3).(*POPM)

	expected := "POPM"
	found := x.GetName()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}

	expected = "Popularimeter"
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

func TestPopmProcess(t *testing.T) {
	x := NewFrame("POPM", "", Version3).(*POPM)
	b := []byte("bob@example.com\x00\xc7\x01\x00\x00\x23")

	x.ProcessData(len(b), b)

	expected := "Popularimeter\n\tEmail: bob@example.com\n\tPopularity: 199\n\tCount: 16777251\n"
	found := x.DisplayContent()
	if found != expected {
		t.Fatalf("Got [%s], Expected [%s]", found, expected)
	}
}
