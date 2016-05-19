package frames

import "testing"

func TestWoafBasicOutput(t *testing.T) {
	x := NewFrame("WOAF", "Official audio webpage", Version3).(*WOAF)

	expected := "WOAF"
	found := x.GetName()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}

	expected = "Official audio webpage"
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

func TestWoafParse(t *testing.T) {
	x := NewFrame("WOAF", "Official webpage", Version3).(*WOAF)
	b := []byte("http://example.com")

	x.ProcessData(len(b), b)
	expected := "(WOAF|Official webpage): http://example.com\n"
	found := x.DisplayContent()
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}
