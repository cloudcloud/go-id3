package file

import (
	"fmt"
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

	Debug  bool `json:"-"`
	file   *os.File
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
func (f *V2) Parse(h *os.File) {
	f.file = h
	offset := v2HeaderOffset

	// have a stream, will read bytes from it
	buf := make([]byte, v2HeaderLength)
	f.file.Seek(v2HeaderStart, v2HeaderStart) // go back to start
	f.file.Read(buf)

	if string(buf[:offset]) != v2HeaderInit {
		panic("Not a valid ID3 Version 2 tag found")
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
		fmt.Println("Unimplemented: handle extended tag header")

		// soaking up the bytes for the moment
		extended := f.nextBytes(v2HeaderLength)
		_ = extended
	}

	defer func() {
		if r := recover(); r != nil {
			// silently continue through... perhaps debug
			fmt.Printf("Caught a panic [%s]\n\n", r)
		}
	}()

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
			if !ok {
				continue
			}

			tmpDetail = f.nextBytes(v2HeaderLength - v2NewByteLen)
			tmpSize = frames.GetSize(tmpDetail[:v2NewByteLen], bitwiseEighthShifter)
		case frames.Version4: // process version 4 frame
			frameName = f.nextBytes(v2NewByteLen)
			if len(frameName) != v2NewByteLen {
				break
			}

			resp, ok = frames.Version24Frames[frames.GetStr(frameName)]
			if !ok {
				continue
			}

			tmpDetail = f.nextBytes(v2HeaderLength - v2NewByteLen)
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
			tmpSize = frames.GetSize(tmpDetail, bitwiseEighthShifter)
		default:
			panic(fmt.Sprintf("Frame version not supported v2.%d.%d", f.Major, f.Min))
		}

		// if no frame, nothing to soak up before breaking
		if tmpSize == 0 {
			break
		}

		// soak up in case of bad frame loading but still frame bytes
		tmpFrame := f.nextBytes(tmpSize)
		if resp == nil {
			fmt.Println("Invalid frame loaded")
			break
		}

		frame = resp()
		if f.Debug {
			fmt.Printf("Pushing in [%s]\n", frame.GetName())
		}
		f.Frames = append(f.Frames, frame.ProcessData(tmpSize, tmpFrame))
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
