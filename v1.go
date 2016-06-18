package id3

import (
	"fmt"

	"github.com/cloudcloud/go-id3/frames"
)

// V1 is a structure for Version 1 ID3 content
type V1 struct {
	Artist  string `json:"artist"`
	Title   string `json:"title"`
	Album   string `json:"album"`
	Year    int    `json:"year"`
	Comment string `json:"comment"`
	Track   int    `json:"track"`
	Genre   int    `json:"genre"`

	Debug bool `json:"-"`
}

const (
	v1TagSize     = 128 // full number of bytes
	v1TagStart    = 3   // offset where tag content begins
	v1TitleEnd    = 33  // offset for track title
	v1ArtistEnd   = 63  // offset for artist name
	v1AlbumEnd    = 93  // offset for album name
	v1YearEnd     = 97  // offset where year completes
	v1CommentEnd  = 125 // offset where comment completes
	v1TagLocation = 2   // direction from which content is read
	v1StrLength   = 30  // base string length
	v1ComLength   = 28  // base comment length
)

// Parse completes the actual processing of the file
// and extracts the tag information.
func (i *V1) Parse(h frames.FrameFile) error {
	b := make([]byte, v1TagSize)
	h.Seek(-int64(v1TagSize), v1TagLocation)
	h.Read(b)

	if frames.GetStr(b[0:v1TagStart]) != "TAG" {
		return fmt.Errorf("No id3v1, %#v", b[0:v1TagStart])
	}
	b = b[v1TagStart:]

	i.Title = frames.GetStr(b[:v1StrLength])
	b = b[v1StrLength:]
	i.Artist = frames.GetStr(b[:v1StrLength])
	b = b[v1StrLength:]
	i.Album = frames.GetStr(b[:v1StrLength])
	b = b[v1StrLength:]
	i.Year = frames.GetInt(b[:4])
	b = b[4:]
	i.Comment = frames.GetStr(b[:v1ComLength])
	b = b[v1ComLength:]

	i.Track = frames.GetInt(b[0:2])
	i.Genre = frames.GetInt([]byte{b[2]})

	return nil
}
