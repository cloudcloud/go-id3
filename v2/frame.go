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
	DisplayContent() string

	GetExplain() string
	GetLength() string
	GetName() string

	Process(b []byte) []byte
}

// Frame defines a base structure shared across all Frame types. This frame
// format is "inherited" within specific Frame type for shared usage.
type Frame struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
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

// DisplayContent will provide a visual representation for pretty printing
// and basic display. This may not necessarily be the actual content.
func (t *Frame) DisplayContent() string {
	return fmt.Sprintf("%s", t.Data)
}

// GetExplain will describe the current field based on the name.
func (t *Frame) GetExplain() string {
	return "{}"
}

// GetLength will return a string of the Length for the frame.
func (t *Frame) GetLength() string {
	return string(t.Size)
}

// GetName will retrieve the current Frame name.
func (t *Frame) GetName() string {
	return t.Name
}

// GetUtf is a shared function to help with the parsing and processing of Utf
// strings. The spec defines the option use Utf16 instead of ISO formats so
// this function is used for that processing.
func GetUtf(b []byte) string {
	var e binary.ByteOrder

	if uint16(b[0])<<8|uint16(b[1]) == 65534 {
		e = binary.LittleEndian
	} else if uint16(b[0])<<8|uint16(b[1]) == 65279 {
		e = binary.BigEndian
	} else {
		return string(b)
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
