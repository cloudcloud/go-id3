package id3

import (
	"bytes"
	"testing"
)

func TestBaseV2(t *testing.T) {
	b := &tfile{}
	b.Write([]byte("Bob is "))
	v := &V2{Debug: false}

	err := v.Parse(b)
	if err == nil {
		t.Fatalf("Invalid Parse() success with V2")
	}
}

func TestParseV2(t *testing.T) {
	b := &tfile{}
	b.Write([]byte("ID3\x03\x00\x40\x00\x00\x00\x2b" +
		"\x00\x00\x00\x00\x00\x10\x00\x00\x00\x00" +
		"TPE1\x00\x00\x00\x0d\x00\x00\x00Cult of Luna" +
		"FAIL\x00\x00\x00\x00\x00\x00"))
	v := &V2{Debug: false}
	err := v.Parse(b)
	if err != nil {
		t.Fatalf("Unable to Parse() with V2 (v2.3.0)")
	}

	found := v.GetArtist()
	expected := "Cult of Luna"
	if found != expected {
		t.Fatalf("Found [%s], Expected [%s]", found, expected)
	}
}

func TestParseV2Crc(t *testing.T) {
	b := &tfile{}
	b.Write([]byte("ID3\x03\x00\x40\x00\x00\x00\x2b" +
		"\x00\x00\x00\x00\x80\x00\x00\x00\x00\x00" +
		"\x00\x02\x00\xaa" +
		"TPE1\x00\x00\x00\x0d\x00\x00\x00Cult of Luna" +
		"FAIL\x00\x00\x00\x00\x00\x00"))
	v := &V2{Debug: false}
	err := v.Parse(b)
	if err != nil {
		t.Fatalf("Unable to Parse() with V2 (v2.3.0)")
	}

	if !v.Crc {
		t.Fatal("Expected CRC processing")
	}

	calculated := v.CrcContent
	should := []byte("\x00\x02\x00\xaa")
	if !bytes.Equal(should, calculated) {
		t.Fatalf("Found %v, Expected %v", should, calculated)
	}
}

func TestParseV2Debug(t *testing.T) {
	b := &tfile{}
	v := &V2{Debug: true}
	b.Write([]byte("ID3\x04\x00\x00\x00\x00\x00\x35" +
		"TEXT\x00\x00\x00\x0a\x00\x00\x00CerealBoy" +
		"FAIL\x00\x00\x00\x00\x00\x00" +
		"TPE2\x00\x00\x00\x0d\x00\x00\x00Cult of Luna"))
	err := v.Parse(b)
	if err != nil {
		t.Fatalf("Unable to Parse() with V2 (v2.4.0)")
	}

	found := v.GetArtist()
	expected := "Cult of Luna"
	if found != expected {
		t.Fatalf("Found [%s], Expected [%s]", found, expected)
	}
}

func TestV2Pe3(t *testing.T) {
	b := &tfile{}
	v := &V2{Debug: true}
	b.Write([]byte("ID3\x04\x00\x00\x00\x00\x00\x46" +
		"TEXT\x00\x00\x00\x0a\x00\x00\x00CerealBoy" +
		"TPE3\x00\x00\x00\x0d\x00\x00\x00Cult of Luna" +
		"TALB\x00\x00\x00\x0b\x00\x00\x00The Beyond"))
	v.Parse(b)

	expected := "Cult of Luna"
	found := v.GetArtist()
	if expected != found {
		t.Fatalf("Got [%s], expected [%s]", found, expected)
	}

	expected = "The Beyond"
	found = v.GetAlbum()
	if expected != found {
		t.Fatalf("Got [%s], expected [%s]", found, expected)
	}
}

func TestV2Pe4(t *testing.T) {
	b := &tfile{}
	v := &V2{Debug: true}
	b.Write([]byte("ID3\x04\x00\x00\x00\x00\x00\x46" +
		"TEXT\x00\x00\x00\x0a\x00\x00\x00CerealBoy" +
		"TPE4\x00\x00\x00\x0d\x00\x00\x00Cult of Luna" +
		"TOAL\x00\x00\x00\x0b\x00\x00\x00The Beyond"))
	v.Parse(b)

	expected := "Cult of Luna"
	found := v.GetArtist()

	if expected != found {
		t.Fatalf("Got [%s], expected [%s]", found, expected)
	}

	expected = "The Beyond"
	found = v.GetAlbum()
	if expected != found {
		t.Fatalf("Got [%s], expected [%s]", found, expected)
	}
}

func TestPanicNextBytes(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Was expecting a panic() from zero length read")
		}
	}()

	v := &V2{}
	_ = v.nextBytes(0)
}

func TestParseV2Original(t *testing.T) {
	b := &tfile{}
	v := &V2{Debug: false}
	b.Write([]byte("ID3\x02\x00\x00\x00\x00\x00\x14" +
		"BUF\x00\x00\x08\x00\x00\x42\x00\x00\x00\x00\x05" +
		"BUD\x00\x00\x00"))
	v.Parse(b)

	expected := 1
	found := len(v.Frames)
	if found != expected {
		t.Fatalf("Got [%d], Expected [%d]", found, expected)
	}
}

func TestParseInvalidVersion(t *testing.T) {
	b := &tfile{}
	v := &V2{}
	b.Write([]byte("ID3\x00\x05\x00\x00\x00\x00\x10" +
		"FAIL\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"))
	v.Parse(b)

	expected := 0
	found := len(v.Frames)
	if found != expected {
		t.Fatalf("Got [%d], Expected [%d]", found, expected)
	}
}

func TestSimpleCatcher(t *testing.T) {
	v := &V2{}
	b := &tfile{}

	v.catcher(b)
	found := b.buf.String()
	expected := "<nil>"
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}

func TestCatcherElegance(t *testing.T) {
	v := &V2{}
	b := &tfile{}

	catchMe(v, b)

	found := b.buf.String()
	expected := "Stumbled upon a panic(), testing.\n"
	if found != expected {
		t.Errorf("Got [%s], Expected [%s]", found, expected)
	}
}

func catchMe(v *V2, b *tfile) {
	defer v.catcher(b)

	panic("testing")
}
