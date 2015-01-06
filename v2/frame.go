package v2

import "fmt"

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
	fmt.Println("In *Frame")

	return b
}

func (f *Frame) Explain() string {
	switch f.Name {
	case "GEOB":
		return "Encapsulated Object"
	case "PRIV":
		return "Private Frame"
	case "TALB":
		return "Album Name"
	case "TBPM":
		return "Beats Per Minute"
	case "TCOM":
		return "Composer"
	case "TCON":
		return "Content Type"
	case "TCOP":
		return "Copyright"
	case "TDAT":
		return "Date"
	case "TENC":
		return "Encoded by"
	case "TEXT":
		return "Lyricist"
	case "TFLT":
		return "File Type"
	case "TIME":
		return "Time"
	case "TIT2":
		return "Title"
	case "TLAN":
		return "Language"
	case "TLEN":
		return "Length"
	case "TOAL":
		return "Original Album"
	case "TOLY":
		return "Original Lyricist"
	case "TOPE":
		return "Original Artist"
	case "TORY":
		return "Original Release Year"
	case "TPE1":
		return "Lead Performer"
	case "TPUB":
		return "Publisher"
	case "TRCK":
		return "Track Number"
	case "TSSE":
		return "Encoding Settings"
	case "TXXX":
		return "User Text"
	case "TYER":
		return "Year"
	case "WXXX":
		return "Provided URL"
	default:
		return "Unknown"
	}
}

func (f *Frame) IsValid() bool {
	return false
}
