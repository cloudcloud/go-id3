package id3

import (
	"fmt"
	"os"
	"strconv"
)

const (
	header_length = 10
	version_begin = 6
	version_end   = 9
	count_limit   = 5
	full_limit    = 40
)

var (
	header []byte
)

type frame struct {
	name string
	data string
	size int
	init byte
}

type Id3V2 struct {
	Items []frame
}

func NewV2() *Id3V2 {
	i := new(Id3V2)

	return i
}

func (i *Id3V2) Parse(f string) {
	b := fileToBuffer(f, header_length, 0)
	if getString(b[0:tag_start]) != "ID3" {
		return
	}

	file, _ := os.Open(f)
	file.Seek(header_length, 0)

	// version information after the identifier
	major_version, _ := strconv.Atoi(fmt.Sprintf("%d", b[tag_start]))
	minor_version, _ := strconv.Atoi(fmt.Sprintf("%d", b[tag_start+1]))

	// we shouldn't need to use the flag at all

	// final four are the size
	size := int(b[version_begin]<<21 | b[version_begin+1]<<14 | b[version_end-1]<<7 | b[version_end])
	_, _, _ = major_version, minor_version, size
	// the annoying thing is, I don't see any of these
	//  really being an issue or actually being used
	//  at all

	// in testing, some frames have broken names
	//  this loop is to attempt to skip and find further
	//  valid names
	count := 0
	for c := 0; c < full_limit; c++ {
		fr := getFrame(file)

		if fr.IsValid() {
			i.Items = append(i.Items, *fr)
		} else {
			count += 1
			if count > count_limit {
				break
			}
		}
	}
}

func (f *frame) Name() string {
	return f.name
}

func (f *frame) Content() string {
	return f.data
}

func (f *frame) Length() string {
	return strconv.Itoa(f.size)
}

func (f *frame) Explain() string {
	switch f.name {
	case "COMM":
		return "Comments"
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

func (f *frame) IsValid() bool {
	if len(f.name) != 4 {
		if len(f.data) > 0 && f.size > 0 {
			debug(fmt.Sprintf("Name [%s] invalid length, but have data [%s]", f.name, f.data))
		} else {
			debug(fmt.Sprintf("Name [%s] invalid, no data", f.name))
		}

		return false
	}

	if f.size < 1 || len(f.data) < 1 {
		debug(fmt.Sprintf("No data [%s] for name [%s]", f.data, f.name))

		return false
	}

	return true
}

func getFrame(f *os.File) *frame {
	fr := new(frame)

	tag := make([]byte, 4)
	size := make([]byte, 4)
	padding := make([]byte, 2)

	f.Read(tag)
	fr.init = tag[0]
	fr.name = getString(tag)

	f.Read(size)
	fr.size = int(size[0]<<21 | size[1]<<14 | size[2]<<7 | size[3])

	frame_content := make([]byte, fr.size)

	f.Read(padding)
	f.Read(frame_content)

	if len(frame_content) > 1 && frame_content[0] == 0 {
		fr.data = getString(frame_content[1:])
		fr.size = fr.size - 1
	} else if len(frame_content) == 1 && frame_content[0] == 0 {
		fr.data = ""
		fr.size = 0
	} else {
		fr.data = getString(frame_content)
	}

	return fr
}
