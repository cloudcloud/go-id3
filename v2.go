package id3

import (
	"fmt"
	"io"
	"os"

	"github.com/cloudcloud/go-id3/frames"
)

// V2 is a structure for housing all ID3 V2 information
// Frames are listed individually, but overall meta is local
type V2 struct {
	Frames         []frames.IFrame `json:"frames"`
	Major          int             `json:"major_version"`
	Min            int             `json:"min_version"`
	Flag           byte            `json:"flag"`
	Size           int             `json:"tag_size"`
	Unsynchronised bool            `json:"unsynchronised"`
	Extended       bool            `json:"extended"`
	Experimental   bool            `json:"experimental"`
	Footer         bool            `json:"footer"`

	ExtendedSize    int    `json:"extended_size" yaml:"extended_size"`
	ExtendedFlag    []byte `json:"extended_flag" yaml:"extended_flag"`
	ExtendedPadding int    `json:"extended_padding" yaml:"extended_padding"`
	Crc             bool   `json:"crc"`
	CrcContent      []byte `json:"crc_content" yaml:"crc_content"`

	Debug  bool `json:"-"`
	file   frames.FrameFile
	offset int
}

const (
	v2HeaderLength        = 10    // bytes in a header
	v2HeaderInit          = "ID3" // everything starts with something
	v2HeaderStart         = 0     // where it will begin
	v2HeaderVersionLength = 1     // a version item length
	v2UnsyncBit           = 7     // offset for bit of unsynchronised flag
	v2ExtendedBit         = 6     // offset for bit of extended flag
	v2ExperimentalBit     = 5     // offset for bit of experimental flag
	v2FooterBit           = 4     // offset for bit of footer flag
	v2HeaderOffset        = 3     // initial offset for the tag header
	v2OrigByteLen         = 3     // original lengths in the v2 spec for frame header items
	v2NewByteLen          = 4     // subsequent versioned lengths
	bitwiseSeventhShifter = 7     // significant bits for an int
	bitwiseEighthShifter  = 8     // significant bits for an int

	v2OffsetMajor = 3
	v2OffsetMinor = 4
	v2OffsetFlag  = 5
	v2OffsetSize  = 6
)

func getBuffer(f frames.FrameFile) ([]byte, error) {
	buf := make([]byte, v2HeaderLength)
	_, _ = f.Seek(v2HeaderStart, v2HeaderStart)
	_, _ = f.Read(buf)

	if string(buf[:v2HeaderOffset]) != v2HeaderInit {
		return nil, fmt.Errorf("not a valid ID3 Version 2 tag found")
	}

	return buf, nil
}

func (f *V2) primeHeaderFromFile(h frames.FrameFile) error {
	f.file = h
	f.Frames = []frames.IFrame{}

	buf, err := getBuffer(h)
	if err != nil {
		return err
	}

	f.Major = frames.GetDirectInt(buf[v2OffsetMajor])
	f.Min = frames.GetDirectInt(buf[v2OffsetMinor])
	f.Flag = buf[v2OffsetFlag]

	f.Unsynchronised = frames.GetBoolBit(f.Flag, v2UnsyncBit)
	f.Extended = frames.GetBoolBit(f.Flag, v2ExtendedBit)
	f.Experimental = frames.GetBoolBit(f.Flag, v2ExperimentalBit)
	f.Footer = frames.GetBoolBit(f.Flag, v2FooterBit)

	f.Size = frames.GetSize(buf[v2OffsetSize:], bitwiseSeventhShifter)

	return nil
}

func (f *V2) primeExtended() {
	if !f.Extended {
		return
	}

	extended := f.nextBytes(v2HeaderLength)

	f.ExtendedSize = frames.GetSize(extended[:4], 8)
	f.ExtendedFlag = extended[4:6]
	f.ExtendedPadding = frames.GetSize(extended[6:], 8)

	f.Crc = frames.GetBoolBit(extended[4], 7)
	if f.Crc {
		f.CrcContent = f.nextBytes(4)
	}
}

// Parse will trawl a file handle for frames
func (f *V2) Parse(h frames.FrameFile) error {
	if err := f.primeHeaderFromFile(h); err != nil {
		return err
	}

	f.primeExtended()

	// wait for a panic
	defer f.catcher(os.Stderr)

	// trawl frames time
	for {
		var resp func() frames.IFrame
		var frame frames.IFrame
		tmpSize := 0

		switch f.Major {
		case frames.Version3:
			tmpSize, resp = f.prepV2Frame(v2NewByteLen, bitwiseEighthShifter, frames.Version23Frames, true)
		case frames.Version4:
			tmpSize, resp = f.prepV2Frame(v2NewByteLen, bitwiseSeventhShifter, frames.Version24Frames, true)
		case frames.Version2:
			tmpSize, resp = f.prepV2Frame(v2OrigByteLen, bitwiseSeventhShifter, frames.Version22Frames, false)
		default:
			return fmt.Errorf("frame version not supported v2.%d.%d", f.Major, f.Min)
		}

		// lack of frame or invalid frame
		if tmpSize == 0 {
			break
		} else if tmpSize == -1 {
			continue
		}

		tmpFrame := f.nextBytes(tmpSize)
		frame = resp()
		if f.Debug {
			fmt.Printf("Pushing in [%s]\n", frame.GetName())
		}
		f.Frames = append(f.Frames, frame.ProcessData(tmpSize, tmpFrame))
	}

	return nil
}

func (f *V2) prepV2Frame(l int, s uint, fr map[string]func() frames.IFrame, before bool) (int, func() frames.IFrame) {
	frameName := f.nextBytes(l)
	if f.Debug {
		fmt.Printf("Potential name [%s]\n", frameName)
	}

	if len(frameName) != l {
		return 0, nil
	}

	resp, ok := fr[frames.GetStr(frameName)]
	detail := []byte{}
	// soak up the bytes before continuing
	if before {
		detail = f.nextBytes(v2HeaderLength - l)
	}

	if !ok {
		return -1, nil
	}

	if !before {
		detail = f.nextBytes(l)
	}

	size := 0
	if before {
		size = frames.GetSize(detail[:l], s)
	} else {
		size = frames.GetSize(detail, s)
	}

	return size, resp
}

// GetFrame will provide a specific Frame if it exists
func (f *V2) GetFrame(n string) frames.IFrame {
	for _, v := range f.Frames {
		if v.GetName() == n {
			return v
		}
	}

	return nil
}

// GetArtist will retrieve the ideal artist for use
func (f *V2) GetArtist() string {
	return f.findTextIdx([]string{"TPE1", "TPE2", "TPE3", "TPE4"})
}

// GetAlbum will determine and give the ideal album
func (f *V2) GetAlbum() string {
	return f.findTextIdx([]string{"TALB", "TOAL"})
}

// GetTitle will determine the ideal title for the song.
func (f *V2) GetTitle() string {
	return f.findTextIdx([]string{"TIT2", "TIT3", "TIT1", "TT2", "TT3", "TT1"})
}

func (f *V2) findTextIdx(i []string) string {
	for _, x := range i {
		if a := f.GetFrame(x); a != nil {
			return a.(*frames.TEXT).Cleaned
		}
	}

	return ""
}

func (f *V2) catcher(o io.Writer) {
	if r := recover(); r != nil {
		fmt.Fprintf(o, "Stumbled upon a panic(), %s.\n", r)
	}
}

func (f *V2) nextBytes(length int) []byte {
	if length < 1 {
		panic("WTF. Please give a positive length for grabbing bytes.")
	}

	f.offset += length
	if f.Debug {
		fmt.Printf("%d vs. %d [total: %d]\n", length, f.offset, f.Size)
	}

	if f.offset > f.Size {
		// want to panic to stop, but hopefully not
		return []byte{}
	}

	t := make([]byte, length)
	_, _ = f.file.Read(t)

	return t
}
