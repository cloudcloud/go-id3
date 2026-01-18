package main

import (
	"bytes"
)

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
