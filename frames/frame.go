// Package frames will provide abstractions for processing indivudal frames within tags
package frames

import (
	"bytes"
	"encoding/binary"
	"fmt"
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
	GetLength() string
	GetName() string

	Init(n, d string, s int)
	ProcessData(int, []byte) IFrame
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
	Version     int    `json:"-"`
	Description string `json:"description"`
	Data        []byte `json:"-"`
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

var (
	// Version22Frames defines the version 2.2 frame mapping
	// http://id3.org/id3v2-00
	Version22Frames = map[string]func() IFrame{
		"BUF": nil,
		"CNT": nil,
		"COM": nil,
		"CRA": nil,
		"CRM": nil,
		"ETC": nil,
		"EQU": nil,
		"GEO": nil,
		"IPL": nil,
		"LNK": nil,
		"MCI": nil,
		"MLL": nil,
		"PIC": nil,
		"POP": nil,
		"REV": nil,
		"RVA": nil,
		"SLT": nil,
		"STC": nil,
		"TAL": nil,
		"TBP": nil,
		"TCM": nil,
		"TCO": nil,
		"TCR": nil,
		"TDA": nil,
		"TDY": nil,
		"TEN": nil,
		"TFT": nil,
		"TIM": nil,
		"TKE": nil,
		"TLA": nil,
		"TLE": nil,
		"TMT": nil,
		"TOA": nil,
		"TOF": nil,
		"TOL": nil,
		"TOR": nil,
		"TOT": nil,
		"TP1": nil,
		"TP2": nil,
		"TP3": nil,
		"TP4": nil,
		"TPA": nil,
		"TPB": nil,
		"TRC": nil,
		"TRD": nil,
		"TRK": nil,
		"TSI": nil,
		"TSS": nil,
		"TT1": nil,
		"TT2": nil,
		"TT3": nil,
		"TXT": nil,
		"TXX": nil,
		"TYE": nil,
		"UFI": nil,
		"ULT": nil,
		"WAF": nil,
		"WAR": nil,
		"WAS": nil,
		"WCM": nil,
		"WCP": nil,
		"WPB": nil,
		"WXX": nil,
	}

	// Version23Frames defines the version 2.3 frame mapping
	// http://id3.org/id3v2.3.0
	Version23Frames = map[string]func() IFrame{
		"AENC": func() IFrame { return NewFrame("AENC", "Audio encryption", Version3) },
		"APIC": func() IFrame { return NewFrame("APIC", "Attached picture", Version3) },
		"COMM": func() IFrame { return NewFrame("COMM", "User comment", Version3) },
		"COMR": func() IFrame { return NewFrame("COMR", "Commercial frame", Version3) },
		"ENCR": func() IFrame { return NewFrame("ENCR", "Encryption method registration", Version3) },
		"EQUA": func() IFrame { return NewFrame("EQUA", "Equalization", Version3) },
		"ETCO": func() IFrame { return NewFrame("ETCO", "Event timing codes", Version3) },
		"GEOB": func() IFrame { return NewFrame("GEOB", "General encapsulated object", Version3) },
		"GRID": func() IFrame { return NewFrame("GRID", "Group identification registration", Version3) },
		"IPLS": func() IFrame { return NewFrame("IPLS", "Involved people list", Version3) },
		"LINK": func() IFrame { return NewFrame("LINK", "Linked information", Version3) },
		"MCDI": func() IFrame { return NewFrame("MCDI", "Music CD identifier", Version3) },
		"MLLT": func() IFrame { return NewFrame("MLLT", "MPEG location lookup table", Version3) },
		"OWNE": func() IFrame { return NewFrame("OWNE", "Ownership frame", Version3) },
		"PRIV": func() IFrame { return NewFrame("PRIV", "Private frame", Version3) },
		"PCNT": func() IFrame { return NewFrame("PCNT", "Play counter", Version3) },
		"POPM": func() IFrame { return NewFrame("POPM", "Popularimeter", Version3) },
		"POSS": func() IFrame { return NewFrame("POSS", "Position synchronisation frame", Version3) },
		"RBUF": func() IFrame { return NewFrame("RBUF", "Recommended buffer size", Version3) },
		"RVAD": func() IFrame { return NewFrame("RVAD", "Relative volume adjustment", Version3) },
		"RVRB": func() IFrame { return NewFrame("RVRB", "Reverb", Version3) },
		"SYLT": func() IFrame { return NewFrame("SYLT", "Synchronized lyrics/text", Version3) },
		"SYTC": func() IFrame { return NewFrame("SYTC", "Synchronized tempo codes", Version3) },
		"TALB": func() IFrame { return NewFrame("TALB", "Album/Show/Movie title", Version3) },
		"TBPM": func() IFrame { return NewFrame("TBPM", "BPM (beats per minute)", Version3) },
		"TCOM": func() IFrame { return NewFrame("TCOM", "Composer", Version3) },
		"TCON": func() IFrame { return NewFrame("TCON", "Content type", Version3) },
		"TCOP": func() IFrame { return NewFrame("TCOP", "Copyright message", Version3) },
		"TDAT": func() IFrame { return NewFrame("TDAT", "Date", Version3) },
		"TDLY": func() IFrame { return NewFrame("TDLY", "Playlist delay", Version3) },
		"TENC": func() IFrame { return NewFrame("TENC", "Encoded by", Version3) },
		"TEXT": func() IFrame { return NewFrame("TEXT", "Lyricist/Text writer", Version3) },
		"TFLT": func() IFrame { return NewFrame("TFLT", "File type", Version3) },
		"TIME": func() IFrame { return NewFrame("TIME", "Time", Version3) },
		"TIT1": func() IFrame { return NewFrame("TIT1", "Content group description", Version3) },
		"TIT2": func() IFrame { return NewFrame("TIT2", "Title/songname/content description", Version3) },
		"TIT3": func() IFrame { return NewFrame("TIT3", "Subtitle/Description refinement", Version3) },
		"TKEY": func() IFrame { return NewFrame("TKEY", "Initial key", Version3) },
		"TLAN": func() IFrame { return NewFrame("TLAN", "Language(s)", Version3) },
		"TLEN": func() IFrame { return NewFrame("TLEN", "Length", Version3) },
		"TMED": func() IFrame { return NewFrame("TMED", "Media type", Version3) },
		"TOAL": func() IFrame { return NewFrame("TOAL", "Original album/movie/show title", Version3) },
		"TOFN": func() IFrame { return NewFrame("TOFN", "Original filename", Version3) },
		"TOLY": func() IFrame { return NewFrame("TOLY", "Original lyricist(s)/text writer(s)", Version3) },
		"TOPE": func() IFrame { return NewFrame("TOPE", "Original artist(s)/performer(s)", Version3) },
		"TORY": func() IFrame { return NewFrame("TORY", "Original release year", Version3) },
		"TOWN": func() IFrame { return NewFrame("TOWN", "File owner/licensee", Version3) },
		"TPE1": func() IFrame { return NewFrame("TPE1", "Lead performer(s)/Soloist(s)", Version3) },
		"TPE2": func() IFrame { return NewFrame("TPE2", "Band/orchestra/accompaniment", Version3) },
		"TPE3": func() IFrame { return NewFrame("TPE3", "Conductor/performer refinement", Version3) },
		"TPE4": func() IFrame { return NewFrame("TPE4", "Interpreted, remixed, or otherwise modified by", Version3) },
		"TPOS": func() IFrame { return NewFrame("TPOS", "Part of a set", Version3) },
		"TPUB": func() IFrame { return NewFrame("TPUB", "Publisher", Version3) },
		"TRCK": func() IFrame { return NewFrame("TRCK", "Track number/Position in set", Version3) },
		"TRDA": func() IFrame { return NewFrame("TRDA", "Recording dates", Version3) },
		"TRSN": func() IFrame { return NewFrame("TRSN", "Internet radio station name", Version3) },
		"TRSO": func() IFrame { return NewFrame("TRSO", "Internet radio station owner", Version3) },
		"TSIZ": func() IFrame { return NewFrame("TSIZ", "Size", Version3) },
		"TSRC": func() IFrame { return NewFrame("TSRC", "ISRC (international standard recording code)", Version3) },
		"TSSE": func() IFrame { return NewFrame("TSSE", "Software/Hardware and settings used for encoding", Version3) },
		"TYER": func() IFrame { return NewFrame("TYER", "Year", Version3) },
		"TXXX": func() IFrame { return NewFrame("TXXX", "User defined text information frame", Version3) },
		"UFID": func() IFrame { return NewFrame("UFID", "Unique File Identifier", Version3) },
		"USER": func() IFrame { return NewFrame("USER", "Terms of use", Version3) },
		"USLT": func() IFrame { return NewFrame("USLT", "Unsynchronised lyrics/text transcription", Version3) },
		"WCOM": func() IFrame { return NewFrame("WCOM", "Commerical information webpage", Version3) },
		"WCOP": func() IFrame { return NewFrame("WCOP", "Copyright/legal information webpage", Version3) },
		"WOAF": func() IFrame { return NewFrame("WOAF", "Official audio file webpage", Version3) },
		"WOAR": func() IFrame { return NewFrame("WOAR", "Official artist/performer webpage", Version3) },
		"WOAS": func() IFrame { return NewFrame("WOAS", "Official audio source webpage", Version3) },
		"WORS": func() IFrame { return NewFrame("WORS", "Official internet radio station homepage", Version3) },
		"WPAY": func() IFrame { return NewFrame("WPAY", "Payment webpage", Version3) },
		"WPUB": func() IFrame { return NewFrame("WPUB", "Publishers official webpage", Version3) },
		"WXXX": func() IFrame { return NewFrame("WXXX", "User defined webpage", Version3) },
	}

	// Version24Frames defines the version 2.4 frame mapping
	// http://id3.org/id3v2.4.0-frames
	Version24Frames = map[string]func() IFrame{
		"AENC": nil,
		"APIC": func() IFrame { return NewFrame("APIC", "Attached picture", Version4) },
		"COMM": nil,
		"COMR": nil,
		"ENCR": nil,
		"EQUA": func() IFrame { return NewFrame("EQUA", "Equalization", Version4) },
		"ETCO": nil,
		"GEOB": nil,
		"GRID": nil,
		"IPLS": nil,
		"LINK": nil,
		"MCDI": nil,
		"MLLT": nil,
		"OWNE": nil,
		"PRIV": nil,
		"PCNT": nil,
		"POPM": nil,
		"POSS": nil,
		"RBUF": nil,
		"RVAD": nil,
		"RVRB": nil,
		"SYLT": nil,
		"SYTC": nil,
		"TALB": nil,
		"TBPM": nil,
		"TCOM": nil,
		"TCON": nil,
		"TCOP": nil,
		"TDAT": nil,
		"TDLY": nil,
		"TENC": nil,
		"TEXT": func() IFrame { return NewFrame("TEXT", "Lyricist/Text writer", Version4) },
		"TFLT": nil,
		"TIME": nil,
		"TIT1": nil,
		"TIT2": func() IFrame { return NewFrame("TIT2", "Title/songname/content description", Version4) },
		"TIT3": nil,
		"TKEY": nil,
		"TLAN": nil,
		"TLEN": nil,
		"TMED": nil,
		"TOAL": nil,
		"TOFN": nil,
		"TOLY": nil,
		"TOPE": nil,
		"TORY": nil,
		"TOWN": nil,
		"TPE1": nil,
		"TPE2": nil,
		"TPE3": nil,
		"TPE4": nil,
		"TPOS": nil,
		"TPUB": nil,
		"TRCK": nil,
		"TRDA": nil,
		"TRSN": nil,
		"TRSO": nil,
		"TSIZ": nil,
		"TSRC": nil,
		"TSSE": nil,
		"TYER": nil,
		"TXXX": nil,
		"UFID": nil,
		"USER": nil,
		"USLT": nil,
		"WCOM": nil,
		"WCOP": nil,
		"WOAF": nil,
		"WOAR": nil,
		"WOAS": nil,
		"WORS": nil,
		"WPAY": nil,
		"WPUB": nil,
		"WXXX": nil,
	}
)

// GetStr will convert the byte slice into a String
func GetStr(b []byte) string {
	p := bytes.IndexByte(b, 0)
	if p == -1 {
		p = len(b)
	}

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

	fin := ""
	if len(str) >= p {
		fin = strings.TrimSpace(str[0:p])
	}

	return fin
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
			if (len(d)-i) < 2 || d[i] == '\x00' && d[i+1] == '\x00' {
				break
			}

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
		s = s << sig
		s |= int(b)
	}

	return s
}

// NewFrame provides a fresh instance of a frame
func NewFrame(n, d string, s int) IFrame {
	var a IFrame

	switch n {
	case "AENC":
		a = new(AENC)
	case "APIC":
		a = new(APIC)
	case "COMM":
		a = new(COMM)
	case "COMR":
		a = new(COMR)
	case "ENCR":
		a = new(ENCR)
	case "EQUA":
		a = new(EQUA)
	case "ETCO":
		a = new(ETCO)
	case "GEOB":
		a = new(GEOB)
	case "GRID":
		a = new(GRID)
	case "IPLS":
		a = new(IPLS)
	case "LINK":
		a = new(LINK)
	case "MCDI":
		a = new(MCDI)
	case "MLLT":
		a = new(MLLT)
	case "OWNE":
		a = new(OWNE)
	case "PCNT":
		a = new(PCNT)
	case "POPM":
		a = new(POPM)
	case "POSS":
		a = new(POSS)
	case "PRIV":
		a = new(PRIV)
	case "RBUF":
		a = new(RBUF)
	case "RVAD":
		a = new(RVAD)
	case "RVRB":
		a = new(RVRB)
	case "SYLT":
		a = new(SYLT)
	case "SYTC":
		a = new(SYTC)
	case "TALB", "TBPM", "TCOM", "TCON", "TCOP", "TDAT", "TDLY", "TENC", "TEXT", "TFLT", "TIME", "TIT1", "TIT2", "TIT3",
		"TKEY", "TLAN", "TLEN", "TMED", "TOAL", "TOFN", "TOLY", "TOPE", "TORY", "TOWN", "TPE1", "TPE2", "TPE3", "TPE4",
		"TPOS", "TPUB", "TRCK", "TRDA", "TRSN", "TRSO", "TSIZ", "TSRC", "TSSE", "TYER":
		a = new(TEXT)
	case "TXXX":
		a = new(TXXX)
	case "UFID":
		a = new(UFID)
	case "USER":
		a = new(USER)
	case "USLT":
		a = new(USLT)
	case "WCOM", "WCOP", "WOAF", "WOAR", "WOAS", "WORS", "WPAY", "WPUB":
		a = new(WOAF)
	case "WXXX":
		a = new(WXXX)
	default:
		fmt.Println(n)
		return a
	}

	// if an interface could have vars...
	a.Init(n, d, s)

	return a
}
