package frames

import (
	"bytes"
	"fmt"
)

// APIC will house any specific APIC frame data
type APIC struct {
	Frame

	MimeType    string `json:"mime_type"`
	PictureType string `json:"picture_type"`
	Image       []byte `json:"image" yaml:"image"`
	Title       string `json:"title"`
}

var picList = map[int]string{
	0:  "Other",
	1:  "32x32 pixels 'file icon' (PNG only)",
	2:  "Other file icon",
	3:  "Cover (front)",
	4:  "Cover (back)",
	5:  "Leaflet page",
	6:  "Media (e.g. lable side of CD)",
	7:  "Lead artist/lead performer/soloist",
	8:  "Artist/performer",
	9:  "Conductor",
	10: "Band/Orchestra",
	11: "Composer",
	12: "Lyricist/text writer",
	13: "Recording Location",
	14: "During recording",
	15: "During performance",
	16: "Movie/video screen capture",
	17: "A bright coloured fish",
	18: "Illustration",
	19: "Band/artist logotype",
	20: "Publisher/Studio logotype",
}

// DisplayContent will comprehensively display known information for APIC
func (a *APIC) DisplayContent() string {
	return fmt.Sprintf("Image (%s, %s, %db) %s\n", a.MimeType, a.PictureType, len(a.Image), a.Title)
}

// ProcessData grabs the meta and binary detail for the image
func (a *APIC) ProcessData(s int, d []byte) IFrame {
	a.Size = s
	a.Data = d

	// encoding is first, 1 is unicode
	if d[0] == '\x01' {
		a.Utf16 = true
	}
	d = d[1:]

	// mime type next, null term
	idx := bytes.IndexByte(d, '\x00')
	a.MimeType = GetStr(d[:idx])

	// picture type
	pic := GetSize([]byte{d[idx+1]}, 8)
	a.PictureType = picList[pic]
	if len(a.PictureType) < 2 {
		a.PictureType = picList[0]
	}
	d = d[idx+2:]

	// image description, null term
	if !a.Utf16 {
		idx = bytes.IndexByte(d, '\x00')
		a.Title = GetStr(d[:idx])

		// image is next now
		a.Image = d[idx+1:]
	} else {
		idx = bytes.Index(d, []byte{'\x00', '\x00'})
		a.Title = GetUnicodeStr(d[:idx])

		a.Image = d[idx+2:]
	}
	a.Size = len(a.Image)

	return a
}
