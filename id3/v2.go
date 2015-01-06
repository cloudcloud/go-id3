package id3

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/cloudcloud/go-id3/v2"
)

const (
	header_length   = 10
	tag_name_length = 4
	version_index   = 3
	flag_index      = 5
	size_begin      = 6
	size_length     = 4
	init_offset     = 0

	give_title  = 0
	give_object = 1
)

var (
	err error
)

type Id3V2 struct {
	Items []v2.Frame `json:"items"`
	Major int        `json:"major_version"`
	Min   int        `json:"min_version"`
	Flag  rune       `json:"flag"`
	Size  int        `json:"size"`

	Unsynchronised bool `json:"unsynchronised"`
	Extended       bool `json:"extended"`
	Experimental   bool `json:"experimental"`
	Footer         bool `json:"footer"`
}

func NewV2() *Id3V2 {
	i := new(Id3V2)

	i.Unsynchronised = false
	i.Extended = false
	i.Experimental = false
	i.Footer = false

	return i
}

func (i *Id3V2) Parse(f string) {
	b := fileToBuffer(f, header_length, init_offset)
	if getString(b[init_offset:version_index]) != "ID3" {
		return
	}

	file, _ := os.Open(f)
	file.Seek(header_length, init_offset)

	i.Major, _ = strconv.Atoi(fmt.Sprintf("%d", b[version_index]))
	i.Min, _ = strconv.Atoi(fmt.Sprintf("%d", b[version_index+1]))
	i.Flag = rune(b[flag_index])

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

	i.Size = int(rune(b[size_begin])<<21 | rune(b[size_begin+1])<<14 | rune(b[size_begin+2])<<7 | rune(b[size_begin+3]))
	b = fileToBuffer(f, i.Size, header_length)

	for {
		if b, err = i.chompFrame(b); err != nil {
			break
		}
	}
}

func (i *Id3V2) chompFrame(b []byte) ([]byte, error) {
	title := b[init_offset:tag_name_length]
	b = b[version_index:]
	t, err := switchTitle(title)

	if err != nil {
		return b, err
	}

	t.Process(b)

	return b, nil
}

func switchTitle(b []byte) (v2.IFrame, error) {
	switch string(b) {
	case "COMM":
		return v2.NewCOMM(string(b)), nil

	case "TYER":
		return v2.NewTEXT(string(b)), nil
	}

	return nil, errors.New("Invalid title")
}
