package frames

import (
	"bytes"
	"fmt"
)

// SYLT defines the Synchronised lyrics/text
type SYLT struct {
	Frame

	Language    string     `json:"language"`
	Format      string     `json:"format"`
	ContentType string     `json:"content_type"`
	Descriptor  string     `json:"lyrics"`
	Items       []*txtItem `json:"items"`
}

var (
	types = map[int]string{
		0: "Other",
		1: "Lyrics",
		2: "Transcription",
		3: "Movement",
		4: "Events",
		5: "Chord",
		6: "Trivia",
	}
)

type txtItem struct {
	Content   string `json:"content"`
	Timestamp int    `json:"timestamp"`
}

// DisplayContent will comprehensively display known information
func (y *SYLT) DisplayContent() string {
	out := fmt.Sprintf("Synchronised (%s). Language(%s) Format(%s) Content Type(%s)\n",
		y.Descriptor,
		y.Language,
		y.Format,
		y.ContentType)

	for _, v := range y.Items {
		out = fmt.Sprintf("%s\t%s [%d]\n", out, v.Content, v.Timestamp)
	}

	return out
}

// ProcessData will handle the acquisition of all data
func (y *SYLT) ProcessData(s int, d []byte) IFrame {
	y.Size = s
	y.Data = d
	y.Items = []*txtItem{}

	x := d[0]
	if x == '\x01' {
		y.Utf16 = true
	}
	d = d[1:]
	y.Language = GetStr(d[0:3])
	d = d[3:]

	form := GetSize([]byte{d[0]}, 1)
	y.Format = "ms"
	if form == 1 {
		y.Format = "mpeg"
	}

	ct := GetSize([]byte{d[1]}, 1)
	if ct > 6 || ct < 0 {
		ct = 0
	}
	y.ContentType = types[ct]
	d = d[2:]

	if !y.Utf16 {
		idx := bytes.IndexByte(d, '\x00')
		y.Descriptor = GetStr(d[:idx])
		d = d[idx+1:]
	} else {
		idx := bytes.Index(d, []byte{'\x00', '\x00'})
		y.Descriptor = GetUnicodeStr(d[:idx])
		d = d[idx+2:]
	}

	for {
		idx := bytes.IndexByte(d, '\x00')
		t := &txtItem{Content: GetStr(d[:idx])}
		d = d[idx+1:]

		t.Timestamp = GetSize(d[:2], 8)
		d = d[2:]

		y.Items = append(y.Items, t)

		if len(d) < 3 || bytes.IndexByte(d, '\x00') < 2 {
			break
		}
	}

	return y
}
