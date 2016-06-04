package frames

import "testing"

func TestLinkGeneral(t *testing.T) {
	x := NewFrame("LINK", "Linked information", Version3).(*LINK)

	expected := "LINK"
	found := x.GetName()
	if found != expected {
		t.Fatalf("Expected [%s], Got [%s]", expected, found)
	}

	expected = "Linked information"
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

func TestLinkParseV3(t *testing.T) {
	x := NewFrame("LINK", "Linked information", Version3).(*LINK)
	b := []byte("xyzhttp://example.com\x00007")
	x.ProcessData(len(b), b)

	expected := "Linked information\n\tIdentifier: xyz\n\tURL: http://example.com\n"
	found := x.DisplayContent()
	if found != expected {
		t.Fatalf("Got [%s], Expected [%s]", found, expected)
	}
}
