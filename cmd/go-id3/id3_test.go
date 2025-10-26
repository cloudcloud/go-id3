package main

import (
	"bytes"
	"testing"
)

func TestBasicTmpl(t *testing.T) {
	defer func() {
		if x := recover(); x == nil {
			t.Errorf("Was expecting a panic in tmpl()")
		}
	}()

	var b bytes.Buffer
	tmpl(&b, "Hello, friend. {{.Fail}}", interface{}(""))
	t.Errorf("Panic wasn't triggered.")
}

func TestBasicName(t *testing.T) {
	found := readCmd.Name()
	expected := "read"

	if found != expected {
		t.Fatalf("Got [%s], Expected [%s]", found, expected)
	}
}

type tfile struct {
	buf *bytes.Buffer
}

func (t *tfile) Seek(o int64, w int) (int64, error) {
	return 0, nil
}

func (t *tfile) Close() error {
	return nil
}

func (t *tfile) Read(b []byte) (int, error) {
	_, _ = t.buf.Read(b)

	return len(b), nil
}

func (t *tfile) Write(b []byte) (int, error) {
	t.buf = bytes.NewBuffer(b)

	return t.buf.Len(), nil
}
