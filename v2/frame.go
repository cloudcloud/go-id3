package v2

import (
	"encoding/binary"
	"fmt"
	"unicode/utf16"
	"unicode/utf8"
)

const (
	tag_length = 4
)

type IFrame interface {
	Process(b []byte) []byte
}

type Frame struct {
	Name string      `json:"name"`
	Data interface{} `json:"name"`
	Size int         `json:"size"`
}

func process(o IFrame, b []byte) {
	o.Process(b)
}

func NewFrame() *Frame {
	f := new(Frame)

	return f
}

func (t *Frame) Process(b []byte) []byte {
	fmt.Println("Frame unimplemented Process()")

	return []byte{}
}

func GetUtf(b []byte) string {
	var e binary.ByteOrder = binary.BigEndian
	if uint16(b[1])<<8|uint16(b[0]) == 0xFFEF {
		e = binary.LittleEndian
	}

	utf := make([]uint16, (len(b)+(2-1))/2)
	for i := 0; i+(2-1) < len(b); i += 2 {
		utf[i/2] = e.Uint16(b[i:])
	}

	if len(b)/2 < len(utf) {
		utf[len(utf)-1] = utf8.RuneError
	}

	return string(utf16.Decode(utf))
}
