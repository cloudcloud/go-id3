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
)

// Parse will trawl a file handle for frames
func (f *V2) Parse(h frames.FrameFile) error {
	f.file = h
	offset := v2HeaderOffset
	f.Frames = []frames.IFrame{}

	// have a stream, will read bytes from it
	buf := make([]byte, v2HeaderLength)
	f.file.Seek(v2HeaderStart, v2HeaderStart) // go back to start
	f.file.Read(buf)

	if string(buf[:offset]) != v2HeaderInit {
		return fmt.Errorf("Not a valid ID3 Version 2 tag found")
	}

	f.Major = frames.GetDirectInt(buf[offset]) // first index for major version

	offset++ // 5th index contains version min identifier
	f.Min = frames.GetDirectInt(buf[offset])

	offset++ // 6th index contains bits for specific tag modifiers
	f.Flag = buf[offset]

	f.Unsynchronised = frames.GetBoolBit(f.Flag, v2UnsyncBit)
	f.Extended = frames.GetBoolBit(f.Flag, v2ExtendedBit)
	f.Experimental = frames.GetBoolBit(f.Flag, v2ExperimentalBit)
	f.Footer = frames.GetBoolBit(f.Flag, v2FooterBit)

	offset++ // final 4 bytes contain size information about the entirety of the tag
	f.Size = frames.GetSize(buf[offset:], bitwiseSeventhShifter)

	if f.Extended {
		extended := f.nextBytes(v2HeaderLength)

		f.ExtendedSize = frames.GetSize(extended[:4], 8)
		f.ExtendedFlag = extended[4:6]
		f.ExtendedPadding = frames.GetSize(extended[6:], 8)

		f.Crc = frames.GetBoolBit(extended[4], 7)
		if f.Crc {
			f.CrcContent = f.nextBytes(4)
		}
	}

	// wait for a panic
	defer f.catcher(os.Stderr)

	// trawl frames time
	for {
		var frameName, tmpDetail []byte
		var resp func() frames.IFrame
		var frame frames.IFrame
		var ok bool
		var tmpSize int

		switch f.Major {
		case frames.Version3: // process version 3 frame
			frameName = f.nextBytes(v2NewByteLen)
			if len(frameName) != v2NewByteLen {
				break
			}

			resp, ok = frames.Version23Frames[frames.GetStr(frameName)]
			tmpDetail = f.nextBytes(v2HeaderLength - v2NewByteLen)
			if !ok {
				continue
			}

			tmpSize = frames.GetSize(tmpDetail[:v2NewByteLen], bitwiseEighthShifter)
		case frames.Version4: // process version 4 frame
			frameName = f.nextBytes(v2NewByteLen)
			if len(frameName) != v2NewByteLen {
				break
			}

			resp, ok = frames.Version24Frames[frames.GetStr(frameName)]
			tmpDetail = f.nextBytes(v2HeaderLength - v2NewByteLen)
			if !ok {
				continue
			}

			tmpSize = frames.GetSize(tmpDetail[:v2NewByteLen], bitwiseSeventhShifter)
		case frames.Version2: // process version 2 frame
			frameName = f.nextBytes(v2OrigByteLen)
			if len(frameName) != v2OrigByteLen {
				break
			}

			resp, ok = frames.Version22Frames[frames.GetStr(frameName)]
			if !ok {
				continue
			}

			tmpDetail = f.nextBytes(v2OrigByteLen)
			tmpSize = frames.GetSize(tmpDetail, bitwiseSeventhShifter)
		default:
			return fmt.Errorf("Frame version not supported v2.%d.%d", f.Major, f.Min)
		}

		// if no frame, nothing to soak up before breaking
		if tmpSize == 0 {
			break
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
	b := ""
	if a := f.GetFrame("TPE1"); a != nil {
		b = a.(*frames.TEXT).Cleaned
	} else if a := f.GetFrame("TPE2"); a != nil {
		b = a.(*frames.TEXT).Cleaned
	} else if a := f.GetFrame("TPE3"); a != nil {
		b = a.(*frames.TEXT).Cleaned
	} else if a := f.GetFrame("TPE4"); a != nil {
		b = a.(*frames.TEXT).Cleaned
	}

	return b
}

// GetAlbum will determine and give the ideal album
func (f *V2) GetAlbum() string {
	b := ""
	if a := f.GetFrame("TALB"); a != nil {
		b = a.(*frames.TEXT).Cleaned
	} else if a := f.GetFrame("TOAL"); a != nil {
		b = a.(*frames.TEXT).Cleaned
	}

	return b
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
	f.file.Read(t)

	return t
}
