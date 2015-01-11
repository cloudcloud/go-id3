// Package v2 provides specific functionality for working with the ID3 Version 2
// file format and frames.
package v2

import (
	"encoding/binary"
	"fmt"
	"unicode/utf16"
	"unicode/utf8"
)

const (
	tagLength = 4
)

// IFrame is a shared interface for use with defining types of Frame formats
// for processing within the ID3 tag.
type IFrame interface {
	Process(b []byte) []byte
}

// Frame defines a base structure shared across all Frame types. This frame
// format is "inherited" within specific Frame type for shared usage.
type Frame struct {
	Name string      `json:"name"`
	Data interface{} `json:"name"`
	Size int         `json:"size"`
}

func process(o IFrame, b []byte) {
	o.Process(b)
}

// NewFrame will provision an instance of a the base Frame.
func NewFrame() *Frame {
	f := new(Frame)

	return f
}

// Process completes the processing from source of the current frame. The
// specific implementation is overridden within each specific frame
// implementation.
func (t *Frame) Process(b []byte) []byte {
	fmt.Println("Frame unimplemented Process()")

	return []byte{}
}

// GetUtf is a shared function to help with the parsing and processing of Utf
// strings. The spec defines the option use Utf16 instead of ISO formats so
// this function is used for that processing.
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
