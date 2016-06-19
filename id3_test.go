package id3

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestBaseFile(t *testing.T) {
	if os.Getenv("INT") != "1" {
		t.SkipNow()
	}

	filename := os.Getenv("FILENAME")
	f := &File{Filename: filename, Debug: true}
	handle, err := os.Open(f.Filename)
	if err != nil {
		t.Fatalf("Unable to open file, %#v", err)
	}
	defer func() {
		if a := recover(); a != nil {
			t.Fatal(a)
		}
	}()

	f.Process(handle)
}

func TestNoFile(t *testing.T) {
	b := &tfile{}
	b.Write([]byte("ID3\x03\x00\x00\x00\x00\x00\x17" +
		"TPE2\x00\x00\x00\x0d\x00\x00\x00Cult of Luna" +
		"TAGBob is great                  " +
		"Bob                           " +
		"Bobbum                        " +
		"2016" +
		"This is just a comment here " +
		"01\x01"))

	f := &File{Debug: false}
	f.Process(b)

	expected := "Artist: Cult of Luna\nAlbum:  Bobbum\n"

	var o bytes.Buffer
	f.PrettyPrint(&o, "text")

	if expected != o.String() {
		t.Fatalf("Got [%s] but expected [%s]", o.String(), expected)
	}
}

func TestNoV2(t *testing.T) {
	b := &tfile{}
	b.Write([]byte("TAGBob is great                  " +
		"Bob                           " +
		"Bobbum                        " +
		"2016" +
		"This is just a comment here " +
		"01\x01"))

	f := &File{Debug: false}
	f.Process(b)

	expected := "Artist: Bob\nAlbum:  Bobbum\n"

	var o bytes.Buffer
	f.PrettyPrint(&o, "text")

	if expected != o.String() {
		t.Fatalf("Got [%s] but expected [%s]", o.String(), expected)
	}
}

func TestNoV2Json(t *testing.T) {
	b := &tfile{}
	b.Write([]byte("TAGBob is great                  " +
		"Bob                           " +
		"Bobbum                        " +
		"2016" +
		"This is just a comment here " +
		"01\x01"))

	f := &File{Debug: false}
	f.Process(b)

	expected := `{"filename":"","id3v1":{"artist":"Bob","title":"Bob is great","album":"Bobbum","year":2016,` +
		`"comment":"This is just a comment here","track":1,"genre":0},"id3v2":{"frames":[],"major_version":0,` +
		`"min_version":0,"flag":0,"tag_size":0,"unsynchronised":false,"extended":false,"experimental":false,` +
		`"footer":false,"extended_size":0,"extended_flag":null,"extended_padding":0,"crc":false,"crc_content":null}}`

	var o bytes.Buffer
	f.PrettyPrint(&o, "json")

	found := strings.TrimRight(o.String(), "\n")
	if expected != found {
		t.Fatalf("Got [%s] but expected [%s]", found, expected)
	}
}

func TestNoV2Yaml(t *testing.T) {
	b := &tfile{}
	b.Write([]byte("TAGBob is great                  " +
		"Bob                           " +
		"Bobbum                        " +
		"2016" +
		"This is just a comment here " +
		"01\x01"))

	f := &File{Debug: false}
	f.Process(b)

	expected := `filename: ""
v1:
  artist: Bob
  title: Bob is great
  album: Bobbum
  year: 2016
  comment: This is just a comment here
  track: 1
  genre: 0
  debug: false
v2:
  frames: []
  major: 0
  min: 0
  flag: 0
  size: 0
  unsynchronised: false
  extended: false
  experimental: false
  footer: false
  extended_size: 0
  extended_flag: []
  extended_padding: 0
  crc: false
  crc_content: []
  debug: false
debug: false`

	var o bytes.Buffer
	f.PrettyPrint(&o, "yaml")

	found := strings.TrimRight(o.String(), "\n")
	if expected != found {
		t.Fatalf("Got [%s] but expected [%s]", found, expected)
	}
}

type tfile struct {
	seekVal int
	buf     *bytes.Buffer
}

func (t *tfile) Seek(o int64, w int) (int64, error) {
	t.seekVal = w

	return 0, nil
}

func (t *tfile) Close() error {
	return nil
}

func (t *tfile) Read(b []byte) (int, error) {
	if t.seekVal == v1TagLocation {
		length := t.buf.Len()
		fmt.Println(length, len(b))

		if length > v1TagSize {
			r := bytes.NewReader(t.buf.Bytes())
			r.ReadAt(b, int64(length-v1TagSize))
		} else {
			t.buf.Read(b)
		}

		t.seekVal = 0
	} else {
		t.buf.Read(b)
	}

	return len(b), nil
}

func (t *tfile) Write(b []byte) (int, error) {
	t.buf = bytes.NewBuffer(b)

	return t.buf.Len(), nil
}
