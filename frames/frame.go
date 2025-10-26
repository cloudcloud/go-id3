// Package frames will provide abstractions for processing indivudal frames within tags
package frames

import (
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode/utf16"
	"unicode/utf8"
)

// IFrame is a shared interface for use with defining types of Frame formats
// for processing within the ID3 tag.
type IFrame interface {
	DisplayContent() string

	GetExplain() string
	GetLength() int
	GetName() string

	Init(n, d string, s int)
	ProcessData(int, []byte) IFrame
}

// FrameFile provides an interface for overloading of os.File
type FrameFile interface {
	Close() error
	Read([]byte) (int, error)
	Seek(int64, int) (int64, error)
	Write([]byte) (int, error)
}

const (
	// Version2 denotes the selection of v2.2.*
	Version2 = 2
	// Version3 denotes the selection of v2.3.*
	Version3 = 3
	// Version4 denotes the selection of v2.4.*
	Version4 = 4
	// LengthUnicode defines the byte size for each character
	LengthUnicode = 2
	// LengthStandard defines the normal byte size for a character
	LengthStandard = 1
)

// Frame defines a base structure shared across all Frame types. This frame
// format is "inherited" within specific Frame type for shared usage.
type Frame struct {
	Name        string `json:"name"`
	Version     int    `json:"-" yaml:"-"`
	Description string `json:"description"`
	Data        []byte `json:"-" yaml:"-"`
	Cleaned     string `json:"cleaned"`
	Size        int    `json:"size"`

	Flags        int  `json:"flags"`
	TagPreserve  bool `json:"tag_preserve"`
	FilePreserve bool `json:"file_preserve"`
	ReadOnly     bool `json:"read_only"`
	Compression  bool `json:"compression"`
	Encryption   bool `json:"encryption"`
	Grouping     bool `json:"grouping"`

	Utf16 bool `json:"utf16"`
}

// GetStr will convert the byte slice into a String
func GetStr(b []byte) string {
	str := string(b)
	if !utf8.ValidString(str) {
		r := make([]rune, 0, len(str))
		for i, v := range str {
			if v == utf8.RuneError {
				_, size := utf8.DecodeRuneInString(str[i:])
				if size == 1 {
					continue
				}
			}
			r = append(r, v)
		}
		str = string(r)
	}

	return strings.Trim(str, " \t\n\r\x00")
}

// GetUnicodeStr will read a unicode byte slice to a string
func GetUnicodeStr(d []byte) string {
	byteOrder := []byte{d[0], d[1]}
	var b binary.ByteOrder
	d = d[2:]

	resp := ""
	if len(d) > 1 {
		if byteOrder[0] == '\xFF' && byteOrder[1] == '\xFE' {
			b = binary.LittleEndian
		} else if byteOrder[0] == '\xFE' && byteOrder[1] == '\xFF' {
			b = binary.BigEndian
		}

		str := make([]uint16, 0, len(d)/2)
		for i := 0; i < len(d); i += 2 {
			str = append(str, b.Uint16(d[i:i+2]))
		}

		resp = string(utf16.Decode(str))
	}

	return resp
}

// GetInt will make use of GetStr and convert the input to an Integer
func GetInt(b []byte) int {
	s, _ := strconv.Atoi(GetStr(b)) // convert internally to string and then to integer

	return s
}

// GetBitInt will generate an int from a range of bits in a byte
func GetBitInt(b byte, ltr bool, l uint) int {
	var a uint

	r := 0
	if !ltr {
		for a = 0; a < l; a++ {
			r += int(b << uint(a*8))
		}
	}
	return r
}

// GetDirectInt uses fmt to convert from []byte to int rather than using strconv
func GetDirectInt(b byte) int {
	i, _ := strconv.Atoi(fmt.Sprintf("%d", b)) // print as an integer before converting to one

	return i
}

// GetBoolBit will retrieve the bit from a byte at the offset as a boolean
func GetBoolBit(b byte, i uint) bool {
	n := byte(1 << i)   // shift the positive the appropriate number of times
	return (b & n) == n // check if the byte has the positive at the desired location
}

// GetSize is used for grabbing the int value of the byte size pieces
func GetSize(b []byte, sig uint) int {
	s := 0
	for _, b := range b {
		s <<= sig
		s |= int(b)
	}

	return s
}

// GetBytePercent will fetch an int from a byte as a percentage
func GetBytePercent(b []byte, sig uint) int {
	var length uint
	if len(b) > 0 {
		length = uint(len(b))
	}

	div := float64(length * 256 / (9 - sig))
	i := float64(GetSize(b, sig))
	x := math.Ceil((i / div) * 100)

	return int(x)
}

// Gen provides a wrapper for generating Frames
func Gen(n, d string, s int) func() IFrame {
	return func() IFrame {
		x := NewFrame(n, d, s)
		return x
	}
}

// NewFrame provides a fresh instance of a frame
func NewFrame(n, d string, s int) IFrame {
	a, ok := frameInst[n]
	if !ok {
		fmt.Printf("Unknown frame [%s]\n", n)
		return nil
	}

	x := a()
	// if an interface could have vars...
	x.Init(n, d, s)

	return x
}

// Init will provide the initial values
func (f *Frame) Init(n, d string, v int) {
	f.Name = n
	f.Description = d
	f.Version = v
}

// GetName simply returns the stored Name for the frame
func (f *Frame) GetName() string {
	return f.Name
}

// GetLength will provide the length of the frame content
func (f *Frame) GetLength() int {
	return f.Size
}

// GetExplain will provide the English description for a Frames' purpose
func (f *Frame) GetExplain() string {
	return f.Description
}
