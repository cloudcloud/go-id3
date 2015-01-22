package id3

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/cloudcloud/go-id3/v2"
)

const (
	headerLength  = 10
	tagNameLength = 4
	versionIndex  = 3
	flagIndex     = 5
	sizeBegin     = 6
	sizeLength    = 4
	initOffset    = 0

	giveTitle  = 0
	giveObject = 1
)

var (
	err error
)

// ID3V2 provides structure for processing and working with
// full ID3 V2 tags and their frames.
type ID3V2 struct {
	Items []v2.IFrame `json:"items"`

	Major int  `json:"major_version"`
	Min   int  `json:"min_version"`
	Flag  rune `json:"flag"`
	Size  int  `json:"size"`

	Unsynchronised bool `json:"unsynchronised"`
	Extended       bool `json:"extended"`
	Experimental   bool `json:"experimental"`
	Footer         bool `json:"footer"`
}

// NewV2 will provision an instance of ID3V2.
func NewV2() *ID3V2 {
	i := new(ID3V2)

	i.Unsynchronised = false
	i.Extended = false
	i.Experimental = false
	i.Footer = false

	return i
}

// Parse completes the full processing of the file provided
// by the string argument.
func (i *ID3V2) Parse(f string) {
	b := fileToBuffer(f, headerLength, initOffset)
	if getString(b[initOffset:versionIndex]) != "ID3" {
		return
	}

	file, _ := os.Open(f)
	file.Seek(headerLength, initOffset)

	i.Major, _ = strconv.Atoi(fmt.Sprintf("%d", b[versionIndex]))
	i.Min, _ = strconv.Atoi(fmt.Sprintf("%d", b[versionIndex+1]))
	i.Flag = rune(b[flagIndex])

	if i.Major > 2 {
		if i.Flag&128 == 128 {
			i.Unsynchronised = true
		}
		if i.Flag&64 == 64 {
			i.Extended = true
		}
		if i.Flag&32 == 32 {
			i.Experimental = true
		}
		if i.Major == 4 && i.Flag&16 == 16 {
			i.Footer = true
		}
	}

	i.Size = int(rune(b[sizeBegin])<<21 | rune(b[sizeBegin+1])<<14 | rune(b[sizeBegin+2])<<7 | rune(b[sizeBegin+3]))
	b = fileToBuffer(f, i.Size, headerLength)

	for {
		if b, err = i.chompFrame(b); err != nil {
			break
		}
	}
}

func (i *ID3V2) chompFrame(b []byte) ([]byte, error) {
	title := b[initOffset:tagNameLength]
	b = b[versionIndex+1:]
	t, err := switchTitle(title)

	if err != nil {
		return b, err
	}

	b = t.Process(b)
	i.Items = append(i.Items, t)

	return b, nil
}

func switchTitle(b []byte) (v2.IFrame, error) {
	switch string(b) {
	case "APIC":
		a := v2.NewAPIC(string(b))
		return a, nil

	case "COMM":
		a := v2.NewCOMM(string(b))
		return a, nil

	case "TALB", "TCOM", "TCON", "TCOP", "TENC", "TIT2", "TOPE", "TPE1", "TPE2", "TRCK", "TYER":
		a := v2.NewTEXT(string(b))
		return a, nil

	case "WXXX":
		a := v2.NewWXXX(string(b))
		return a, nil

	default:
		fmt.Println(string(b), b)
	}

	return nil, errors.New("invalid title")
}
