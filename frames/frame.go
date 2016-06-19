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

var (
	// Version22Frames defines the version 2.2 frame mapping
	// http://id3.org/id3v2-00
	Version22Frames = map[string]func() IFrame{
		"BUF": Gen("BUF", "Recommended buffer size", Version2),
		"CNT": Gen("CNT", "Play counter", Version2),
		"COM": Gen("COM", "Comments", Version2),
		"CRA": Gen("CRA", "Audio encryption", Version2),
		"CRM": Gen("CRM", "Encrypted meta frame", Version2),
		"ETC": Gen("ETC", "Event timing codes", Version2),
		"EQU": Gen("EQU", "Equalization", Version2),
		"GEO": Gen("GEO", "General encapsulated object", Version2),
		"IPL": Gen("IPL", "Involved people list", Version2),
		"LNK": Gen("LNK", "Linked information", Version2),
		"MCI": Gen("MCI", "Music CD identifier", Version2),
		"MLL": Gen("MLL", "MPEG location lookup table", Version2),
		"PIC": Gen("PIC", "Attached picture", Version2),
		"POP": Gen("POP", "Popularimeter", Version2),
		"REV": Gen("REV", "Reverb", Version2),
		"RVA": Gen("RVA", "Relative volume adjustment", Version2),
		"SLT": Gen("SLT", "Synchronized lyric/text", Version2),
		"STC": Gen("STC", "Synched tempo codes", Version2),
		"TAL": Gen("TAL", "Album/Movie/Show title", Version2),
		"TBP": Gen("TBP", "BPM (Beats Per Minute)", Version2),
		"TCM": Gen("TCM", "Composer", Version2),
		"TCO": Gen("TCO", "Content type", Version2),
		"TCR": Gen("TCR", "Copyright message", Version2),
		"TDA": Gen("TDA", "Date", Version2),
		"TDY": Gen("TDY", "Playlist delay", Version2),
		"TEN": Gen("TEN", "Encoded by", Version2),
		"TFT": Gen("TFT", "File type", Version2),
		"TIM": Gen("TIM", "Time", Version2),
		"TKE": Gen("TKE", "Initial key", Version2),
		"TLA": Gen("TLA", "Language(s)", Version2),
		"TLE": Gen("TLE", "Length", Version2),
		"TMT": Gen("TMT", "Media type", Version2),
		"TOA": Gen("TOA", "Original artist(s)/performer(s)", Version2),
		"TOF": Gen("TOF", "Original filename", Version2),
		"TOL": Gen("TOL", "Original Lyricist(s)/text writer(s)", Version2),
		"TOR": Gen("TOR", "Original release year", Version2),
		"TOT": Gen("TOT", "Original album/Movie/Show title", Version2),
		"TP1": Gen("TP1", "Lead artist(s)/Lead performer(s)/Soloist(s)/Performing group", Version2),
		"TP2": Gen("TP2", "Band/Orchestra/Accompaniment", Version2),
		"TP3": Gen("TP3", "Conductor/Performer refinement", Version2),
		"TP4": Gen("TP4", "Interpreted, remixed, or otherwise modified by", Version2),
		"TPA": Gen("TPA", "Part of a set", Version2),
		"TPB": Gen("TPB", "Publisher", Version2),
		"TRC": Gen("TRC", "ISRC (International Standard Recording Code)", Version2),
		"TRD": Gen("TRD", "Recording dates", Version2),
		"TRK": Gen("TRK", "Track number/Position in set", Version2),
		"TSI": Gen("TSI", "Size", Version2),
		"TSS": Gen("TSS", "Software/hardware and settings used for encoding", Version2),
		"TT1": Gen("TT1", "Content group description", Version2),
		"TT2": Gen("TT2", "Title/Songname/Content description", Version2),
		"TT3": Gen("TT3", "Subtitle/Description refinement", Version2),
		"TXT": Gen("TXT", "Lyricist/text writer", Version2),
		"TXX": Gen("TXX", "User defined text information frame", Version2),
		"TYE": Gen("TYE", "Year", Version2),
		"UFI": Gen("UFI", "Unique file identifier", Version2),
		"ULT": Gen("ULT", "Unsynchronized lyric/text transcription", Version2),
		"WAF": Gen("WAF", "Official audio file webpage", Version2),
		"WAR": Gen("WAR", "Official artist/performer webpage", Version2),
		"WAS": Gen("WAS", "Official audio source webpage", Version2),
		"WCM": Gen("WCM", "Commercial information", Version2),
		"WCP": Gen("WCP", "Copyright/Legal information", Version2),
		"WPB": Gen("WPB", "Publishers official webpage", Version2),
		"WXX": Gen("WXX", "User defined URL link frame", Version2),
	}

	// Version23Frames defines the version 2.3 frame mapping
	// http://id3.org/id3v2.3.0
	Version23Frames = map[string]func() IFrame{
		"AENC": Gen("AENC", "Audio encryption", Version3),
		"APIC": Gen("APIC", "Attached picture", Version3),
		"COMM": Gen("COMM", "User comment", Version3),
		"COMR": Gen("COMR", "Commercial frame", Version3),
		"ENCR": Gen("ENCR", "Encryption method registration", Version3),
		"EQUA": Gen("EQUA", "Equalization", Version3),
		"ETCO": Gen("ETCO", "Event timing codes", Version3),
		"GEOB": Gen("GEOB", "General encapsulated object", Version3),
		"GRID": Gen("GRID", "Group identification registration", Version3),
		"IPLS": Gen("IPLS", "Involved people list", Version3),
		"LINK": Gen("LINK", "Linked information", Version3),
		"MCDI": Gen("MCDI", "Music CD identifier", Version3),
		"MLLT": Gen("MLLT", "MPEG location lookup table", Version3),
		"OWNE": Gen("OWNE", "Ownership frame", Version3),
		"PRIV": Gen("PRIV", "Private frame", Version3),
		"PCNT": Gen("PCNT", "Play counter", Version3),
		"POPM": Gen("POPM", "Popularimeter", Version3),
		"POSS": Gen("POSS", "Position synchronisation frame", Version3),
		"RBUF": Gen("RBUF", "Recommended buffer size", Version3),
		"RVAD": Gen("RVAD", "Relative volume adjustment", Version3),
		"RVRB": Gen("RVRB", "Reverb", Version3),
		"SYLT": Gen("SYLT", "Synchronized lyrics/text", Version3),
		"SYTC": Gen("SYTC", "Synchronized tempo codes", Version3),
		"TALB": Gen("TALB", "Album/Show/Movie title", Version3),
		"TBPM": Gen("TBPM", "BPM (beats per minute)", Version3),
		"TCOM": Gen("TCOM", "Composer", Version3),
		"TCON": Gen("TCON", "Content type", Version3),
		"TCOP": Gen("TCOP", "Copyright message", Version3),
		"TDAT": Gen("TDAT", "Date", Version3),
		"TDLY": Gen("TDLY", "Playlist delay", Version3),
		"TENC": Gen("TENC", "Encoded by", Version3),
		"TEXT": Gen("TEXT", "Lyricist/Text writer", Version3),
		"TFLT": Gen("TFLT", "File type", Version3),
		"TIME": Gen("TIME", "Time", Version3),
		"TIT1": Gen("TIT1", "Content group description", Version3),
		"TIT2": Gen("TIT2", "Title/songname/content description", Version3),
		"TIT3": Gen("TIT3", "Subtitle/Description refinement", Version3),
		"TKEY": Gen("TKEY", "Initial key", Version3),
		"TLAN": Gen("TLAN", "Language(s)", Version3),
		"TLEN": Gen("TLEN", "Length", Version3),
		"TMED": Gen("TMED", "Media type", Version3),
		"TOAL": Gen("TOAL", "Original album/movie/show title", Version3),
		"TOFN": Gen("TOFN", "Original filename", Version3),
		"TOLY": Gen("TOLY", "Original lyricist(s)/text writer(s)", Version3),
		"TOPE": Gen("TOPE", "Original artist(s)/performer(s)", Version3),
		"TORY": Gen("TORY", "Original release year", Version3),
		"TOWN": Gen("TOWN", "File owner/licensee", Version3),
		"TPE1": Gen("TPE1", "Lead performer(s)/Soloist(s)", Version3),
		"TPE2": Gen("TPE2", "Band/orchestra/accompaniment", Version3),
		"TPE3": Gen("TPE3", "Conductor/performer refinement", Version3),
		"TPE4": Gen("TPE4", "Interpreted, remixed, or otherwise modified by", Version3),
		"TPOS": Gen("TPOS", "Part of a set", Version3),
		"TPUB": Gen("TPUB", "Publisher", Version3),
		"TRCK": Gen("TRCK", "Track number/Position in set", Version3),
		"TRDA": Gen("TRDA", "Recording dates", Version3),
		"TRSN": Gen("TRSN", "Internet radio station name", Version3),
		"TRSO": Gen("TRSO", "Internet radio station owner", Version3),
		"TSIZ": Gen("TSIZ", "Size", Version3),
		"TSRC": Gen("TSRC", "ISRC (international standard recording code)", Version3),
		"TSSE": Gen("TSSE", "Software/Hardware and settings used for encoding", Version3),
		"TYER": Gen("TYER", "Year", Version3),
		"TXXX": Gen("TXXX", "User defined text information frame", Version3),
		"UFID": Gen("UFID", "Unique File Identifier", Version3),
		"USER": Gen("USER", "Terms of use", Version3),
		"USLT": Gen("USLT", "Unsynchronised lyrics/text transcription", Version3),
		"WCOM": Gen("WCOM", "Commerical information webpage", Version3),
		"WCOP": Gen("WCOP", "Copyright/legal information webpage", Version3),
		"WOAF": Gen("WOAF", "Official audio file webpage", Version3),
		"WOAR": Gen("WOAR", "Official artist/performer webpage", Version3),
		"WOAS": Gen("WOAS", "Official audio source webpage", Version3),
		"WORS": Gen("WORS", "Official internet radio station homepage", Version3),
		"WPAY": Gen("WPAY", "Payment webpage", Version3),
		"WPUB": Gen("WPUB", "Publishers official webpage", Version3),
		"WXXX": Gen("WXXX", "User defined webpage", Version3),
	}

	// Version24Frames defines the version 2.4 frame mapping
	// http://id3.org/id3v2.4.0-frames
	Version24Frames = map[string]func() IFrame{
		"AENC": Gen("AENC", "Audio encryption", Version4),
		"APIC": Gen("APIC", "Attached picture", Version4),
		"ASPI": Gen("ASPI", "Audio seek point index", Version4),
		"COMM": Gen("COMM", "Comments", Version4),
		"COMR": Gen("COMR", "Commercial frame", Version4),
		"ENCR": Gen("ENCR", "Encryption method registration", Version4),
		"EQUA": Gen("EQUA", "Equalization", Version4),
		"EQU2": Gen("EQU2", "Equalisation (2)", Version4),
		"ETCO": Gen("ETCO", "Event timing codes", Version4),
		"GEOB": Gen("GEOB", "General encapsulated object", Version4),
		"GRID": Gen("GRID", "Group registration identifier", Version4),
		"IPLS": Gen("IPLS", "Involved people list", Version4),
		"LINK": Gen("LINK", "Linked information", Version4),
		"MCDI": Gen("MCDI", "Music CD identifier", Version4),
		"MLLT": Gen("MLLT", "MPEG location lookup table", Version4),
		"OWNE": Gen("OWNE", "Ownership frame", Version4),
		"PRIV": Gen("PRIV", "Private frame", Version4),
		"PCNT": Gen("PCNT", "Play counter", Version4),
		"POPM": Gen("POPM", "Popularimeter", Version4),
		"POSS": Gen("POSS", "Position synchronisation frame", Version4),
		"RBUF": Gen("RBUF", "Recommended buffer size", Version4),
		"RVAD": Gen("RVAD", "Relative volume adjustment", Version4),
		"RVA2": Gen("RVA2", "Relative volume adjustment (2)", Version4),
		"RVRB": Gen("RVRB", "Reverb", Version4),
		"SEEK": Gen("SEEK", "Seek frame", Version4),
		"SIGN": Gen("SIGN", "Signature frame", Version4),
		"SYLT": Gen("SYLT", "Synchronised lyric/text", Version4),
		"SYTC": Gen("SYTC", "Synchronised tempo codes", Version4),
		"TALB": Gen("TALB", "Album/Movie/Show title", Version4),
		"TBPM": Gen("TBPM", "BPM (beats per minute)", Version4),
		"TCOM": Gen("TCOM", "Composer", Version4),
		"TCON": Gen("TCON", "Content type", Version4),
		"TCOP": Gen("TCOP", "Copyright message", Version4),
		"TDAT": Gen("TDAT", "Date", Version4),
		"TDEN": Gen("TDEN", "Encoding time", Version4),
		"TDLY": Gen("TDLY", "Playlist delay", Version4),
		"TDOR": Gen("TDOR", "Original release time", Version4),
		"TDRC": Gen("TDRC", "Recording time", Version4),
		"TDRL": Gen("TDRL", "Release time", Version4),
		"TDTG": Gen("TDTG", "Tagging time", Version4),
		"TENC": Gen("TENC", "Encoded by", Version4),
		"TEXT": Gen("TEXT", "Lyricist/Text writer", Version4),
		"TFLT": Gen("TFLT", "File type", Version4),
		"TIME": Gen("TIME", "Time", Version4),
		"TIPL": Gen("TIPL", "Involved people list", Version4),
		"TIT1": Gen("TIT1", "Content group description", Version4),
		"TIT2": Gen("TIT2", "Title/songname/content description", Version4),
		"TIT3": Gen("TIT3", "Subtitle/Description refinement", Version4),
		"TKEY": Gen("TKEY", "Initial key", Version4),
		"TLAN": Gen("TLAN", "Language(s)", Version4),
		"TLEN": Gen("TLEN", "Length", Version4),
		"TMCL": Gen("TMCL", "Musician credits list", Version4),
		"TMED": Gen("TMED", "Media type", Version4),
		"TMOO": Gen("TMOO", "Mood", Version4),
		"TOAL": Gen("TOAL", "Original album/movie/show title", Version4),
		"TOFN": Gen("TOFN", "Original filename", Version4),
		"TOLY": Gen("TOLY", "Original lyricist(s)/text writer(s)", Version4),
		"TOPE": Gen("TOPE", "Original artist(s)/performer(s)", Version4),
		"TORY": Gen("TORY", "Original release year", Version4),
		"TOWN": Gen("TOWN", "File owner/licensee", Version4),
		"TPE1": Gen("TPE1", "Lead performer(s)/Soloist(s)", Version4),
		"TPE2": Gen("TPE2", "Band/orchestra/accompaniment", Version4),
		"TPE3": Gen("TPE3", "Conductor/performer refinement", Version4),
		"TPE4": Gen("TPE4", "Interpreted, remixed, or otherwise modified by", Version4),
		"TPOS": Gen("TPOS", "Part of a set", Version4),
		"TPRO": Gen("TPRO", "Produced notice", Version4),
		"TPUB": Gen("TPUB", "Publisher", Version4),
		"TRCK": Gen("TRCK", "Track number/Position in set", Version4),
		"TRDA": Gen("TRDA", "Recording dates", Version4),
		"TRSN": Gen("TRSN", "Internet radio station name", Version4),
		"TRSO": Gen("TRSO", "Internet radio station owner", Version4),
		"TSIZ": Gen("TSIZ", "Size", Version4),
		"TSOA": Gen("TSOA", "Album sort order", Version4),
		"TSOP": Gen("TSOP", "Performer sort order", Version4),
		"TSOT": Gen("TSOT", "Title sort order", Version4),
		"TSRC": Gen("TSRC", "ISRC (international standard recording code)", Version4),
		"TSSE": Gen("TSSE", "Software/Hardware and settings used for encoding", Version4),
		"TSST": Gen("TSST", "Set subtitle", Version4),
		"TYER": Gen("TYER", "Year", Version4),
		"TXXX": Gen("TXXX", "User defined text information frame", Version4),
		"UFID": Gen("TYER", "Year", Version4),
		"USER": Gen("USER", "Terms of use", Version4),
		"USLT": Gen("USLT", "Unsynchronised lyric/text transcription", Version4),
		"WCOM": Gen("WCOM", "Commercial information", Version4),
		"WCOP": Gen("WCOP", "Copyright/Legal information", Version4),
		"WOAF": Gen("WOAF", "Official audio file webpage", Version4),
		"WOAR": Gen("WOAR", "Official artist/performer webpage", Version4),
		"WOAS": Gen("WOAS", "Official audio source webpage", Version4),
		"WORS": Gen("WORS", "Official Internet radio station homepage", Version4),
		"WPAY": Gen("WPAY", "Payment", Version4),
		"WPUB": Gen("WPUB", "Publishers official webpage", Version4),
		"WXXX": Gen("WXXX", "User defined URL link frame", Version4),
	}
)

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

	return strings.TrimSpace(str)
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
func GetBitInt(b byte, ltr bool, l int) int {
	r := 0
	if !ltr {
		for a := 0; a < l; a++ {
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
		s = s << sig
		s |= int(b)
	}

	return s
}

// GetBytePercent will fetch an int from a byte as a percentage
func GetBytePercent(b []byte, sig uint) int {
	div := float64(len(b) * 256 / (9 - int(sig)))
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
	var a IFrame

	switch n {
	case "CRA", "AENC":
		a = new(AENC)
	case "PIC", "APIC":
		a = new(APIC)
	case "ASPI":
		a = new(ASPI)
	case "COM", "COMM":
		a = new(COMM)
	case "COMR":
		a = new(COMR)
	case "CRM":
		a = new(CRM)
	case "ENCR":
		a = new(ENCR)
	case "EQU", "EQUA":
		a = new(EQUA)
	case "EQU2":
		a = new(EQU2)
	case "ETC", "ETCO":
		a = new(ETCO)
	case "GEO", "GEOB":
		a = new(GEOB)
	case "GRID":
		a = new(GRID)
	case "IPL", "IPLS":
		a = new(IPLS)
	case "LNK", "LINK":
		a = new(LINK)
	case "MCI", "MCDI":
		a = new(MCDI)
	case "MLL", "MLLT":
		a = new(MLLT)
	case "OWNE":
		a = new(OWNE)
	case "CNT", "PCNT":
		a = new(PCNT)
	case "POP", "POPM":
		a = new(POPM)
	case "POSS":
		a = new(POSS)
	case "PRIV":
		a = new(PRIV)
	case "BUF", "RBUF":
		a = new(RBUF)
	case "RVA", "RVAD":
		a = new(RVAD)
	case "RVA2":
		a = new(RVA2)
	case "REV", "RVRB":
		a = new(RVRB)
	case "SEEK":
		a = new(SEEK)
	case "SIGN":
		a = new(SIGN)
	case "SLT", "SYLT":
		a = new(SYLT)
	case "STC", "SYTC":
		a = new(SYTC)
	case "TRC", "TBP", "TSS", "TRK", "TP3", "TP2", "TP1", "TOR", "TMT", "TLE", "TCM", "TDY", "TOT", "TCR", "TT1",
		"TT2", "TP4", "TCO", "TOF", "TPB", "TXT", "TOL", "TIM", "TFT", "TDA", "TYE", "TPA", "TEN", "TKE", "TRD", "TSI",
		"TOA", "TT3", "TLA", "TAL", "TALB", "TBPM", "TCOM", "TCON", "TCOP", "TDAT", "TDEN", "TDLY", "TDOR", "TDRC",
		"TDRL", "TDTG", "TENC", "TEXT", "TFLT", "TIME", "TIPL", "TIT1", "TIT2", "TIT3", "TKEY", "TLAN", "TLEN", "TMCL",
		"TMED", "TMOO", "TOAL", "TOFN", "TOLY", "TOPE", "TORY", "TOWN", "TPE1", "TPE2", "TPE3", "TPE4", "TPOS", "TPRO",
		"TPUB", "TRCK", "TRDA", "TRSN", "TRSO", "TSIZ", "TSOA", "TSOP", "TSOT", "TSRC", "TSSE", "TSST", "TYER":
		a = new(TEXT)
	case "TXX", "TXXX":
		a = new(TXXX)
	case "UFI", "UFID":
		a = new(UFID)
	case "USER":
		a = new(USER)
	case "ULT", "USLT":
		a = new(USLT)
	case "WAR", "WAS", "WAF", "WCM", "WPB", "WCP", "WCOM", "WCOP", "WOAF", "WOAR", "WOAS", "WORS", "WPAY", "WPUB":
		a = new(WOAF)
	case "WXX", "WXXX":
		a = new(WXXX)
	default:
		fmt.Println(n)
		return a
	}

	// if an interface could have vars...
	a.Init(n, d, s)

	return a
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
