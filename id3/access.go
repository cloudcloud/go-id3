package id3

import (
	"fmt"
	"strconv"
)

// Getters for information trawled via the ID3 processing.
//
// Descriptions, references, and associations are all taken from
// http://id3.org/id3v2.4.0-frames and related docs (versioned).

// GetArtist to retrieve the Artist of the song.
func (i *ID3) GetArtist() string {
	a := i.V2.GetFrameValue("TPE1")
	// surely there's an elegant way to do this
	if len(a) < 1 {
		a = i.V2.GetFrameValue("TPE2")
	}
	if len(a) < 1 {
		a = i.V2.GetFrameValue("TPE3")
	}
	if len(a) < 1 {
		a = i.V2.GetFrameValue("TPE4")
	}
	if len(a) < 1 {
		a = i.V2.GetFrameValue("TEXT")
	}
	if len(a) < 1 {
		a = i.V2.GetFrameValue("TOPE")
	}

	return a
}

// GetAlbum to retrive the song Album.
func (i *ID3) GetAlbum() string {
	a := i.V2.GetFrameValue("TALB")
	if len(a) < 1 {
		a = i.V2.GetFrameValue("TOAL")
	}

	return a
}

// GetTitle to retrieve the song Title.
func (i *ID3) GetTitle() string {
	tmpTitle := i.V2.GetFrameValue("TIT3")
	s := i.V2.GetFrameValue("TIT2")

	if len(s) > 0 && len(tmpTitle) > 0 {
		return fmt.Sprintf("%s (%s)", s, tmpTitle)
	}

	return s
}

// GetBPM provides the BPM of the track.
func (i *ID3) GetBPM() string {
	return i.V2.GetFrameValue("TBPM")
}

// GetComposer will provide the composer.
func (i *ID3) GetComposer() string {
	return i.V2.GetFrameValue("TCOM")
}

// GetContentType provides the industry type of content, such as Remix.
func (i *ID3) GetContentType() string {
	return i.V2.GetFrameValue("TCON")
}

// GetCopyright provides specific copyright information about the track.
func (i *ID3) GetCopyright() string {
	return i.V2.GetFrameValue("TCOP")
}

// GetDate provides the date of the track.
func (i *ID3) GetDate() string {
	return i.V2.GetFrameValue("TDAT")
}

// GetPlaylistDelay defines a number of prefixed milliseconds that should be added.
func (i *ID3) GetPlaylistDelay() string {
	return i.V2.GetFrameValue("TDLY")
}

// GetEncodedBy provides information about the file encoder.
func (i *ID3) GetEncodedBy() string {
	return i.V2.GetFrameValue("TENC")
}

// GetFileType provides detail for the format of the file itself such as MP3.
func (i *ID3) GetFileType() string {
	return i.V2.GetFrameValue("TFLT")
}

// GetInitialKey provides the Key in which the track begins.
func (i *ID3) GetInitialKey() string {
	return i.V2.GetFrameValue("TKEY")
}

// GetLength will give the length of the song in Seconds.
func (i *ID3) GetLength() int {
	x, err := strconv.Atoi(i.V2.GetFrameValue("TLEN"))

	if err != nil {
		return 0
	}
	return x
}

// GetReleaseYear will provide the 4 digit year of release.
func (i *ID3) GetReleaseYear() int {
	x, err := strconv.Atoi(i.V2.GetFrameValue("TORY"))

	if err != nil {
		return 0
	}
	return x
}

// GetTime is the time of day the recording was made.
func (i *ID3) GetTime() string {
	return i.V2.GetFrameValue("TIME")
}

// GetGenre provides information about genre details.
func (i *ID3) GetGenre() string {
	return i.V2.GetFrameValue("TIT1")
}

// GetLanguage will give the language of the track.
func (i *ID3) GetLanguage() string {
	return i.V2.GetFrameValue("TLAN")
}

// GetMediaType provides the original Medium for the track.
func (i *ID3) GetMediaType() string {
	return i.V2.GetFrameValue("TMED")
}

// GetOriginalFilename returns the filename for the original file.
func (i *ID3) GetOriginalFilename() string {
	return i.V2.GetFrameValue("TOFN")
}

// GetOriginalLyricist provides the original lyric writer.
func (i *ID3) GetOriginalLyricist() string {
	return i.V2.GetFrameValue("TOLY")
}

// GetOwner provides the owner or licensee of the track.
func (i *ID3) GetOwner() string {
	return i.V2.GetFrameValue("TOWN")
}

// GetSetPart provides details on multiple discs or mediums. i.e. A double disc release.
func (i *ID3) GetSetPart() int {
	// A forward slash is used for PartNumber/Total
	a := i.V2.GetFrameValue("TPOS")
	if len(a) < 1 {
		return 0
	}

	if a[1] == '/' {
		x, err := strconv.Atoi(string(a[0]))

		if err == nil {
			return x
		}
	}

	x, err := strconv.Atoi(a)
	if err != nil {
		return 0
	}

	return x
}

// GetPublisher provides the publisher detail.
func (i *ID3) GetPublisher() string {
	return i.V2.GetFrameValue("TPUB")
}

// GetTrackNumber provides the number of track in the current set.
func (i *ID3) GetTrackNumber() int {
	// Possible / to give CurrentTrack/Total
	a := i.V2.GetFrameValue("TRCK")
	if len(a) < 1 {
		return 0
	} else if len(a) == 1 {
		x, err := strconv.Atoi(a)

		if err == nil {
			return x
		}

		// single char non-int
		return 0
	}

	if a[1] == '/' {
		x, err := strconv.Atoi(string(a[0]))

		if err == nil {
			return x
		}
	}

	x, err := strconv.Atoi(a)
	if err == nil {
		return x
	}
	return 0
}

// GetComment will grab any comment field in the track.
func (i *ID3) GetComment() string {
	return i.V2.GetFrameValue("COMM")
}
