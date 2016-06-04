package frames

import "testing"

func TestCommBasicOutput(t *testing.T) {
	x := NewFrame("COMM", "Comment", Version3).(*COMM)
	if x.GetName() != "COMM" {
		t.Error("Invalid name from COMM frame")
	}

	if x.GetExplain() != "Comment" {
		t.Error("Invalid COMM GetExplain() response")
	}

	x.Name = "BOB"
	if x.GetName() != "BOB" {
		t.Error("Invalid COMM Name setting")
	}

	b := []byte("\x00engComment\x00This is a comment")
	x.ProcessData(len(b), b)
	if x.Language != "eng" {
		t.Fatalf("Invalid language in COMM, got %s but expected eng", x.Language)
	}
	if x.ContentDescription != "Comment" {
		t.Fatalf("Invalid COMM title found")
	}
	if x.Comment != "This is a comment" {
		t.Fatalf("Invalid COMM content parsing")
	}

	if x.DisplayContent() != "Title: Comment\nComment: This is a comment" {
		t.Fatalf("Invalid COMM DisplayContent")
	}

	b = []byte("\x01eng\xff\xfe\xe4\xba\x88\xe8\xa5\xb2\xe5\xbe\xa9\xe8\xae\x90\x00\x00" +
		"\xff\xfe\xe4\xba\x88\xe8\xa5\xb2\xe5\xbe\xa9\xe8\xae\x90")
	x.ProcessData(len(b), b)
	if x.Language != "eng" || x.Utf16 != true {
		t.Fatal("Invalid COMM basic Utf16 parsing")
	}
	if x.GetLength() != 34 {
		t.Fatal("Invalid COMM string Utf16 get")
	}
}
