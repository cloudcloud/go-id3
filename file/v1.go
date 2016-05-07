package file

import (
	"os"

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
	v1TagSize    = 128 // full number of bytes
	v1TagStart   = 3   // offset where tag content begins
	v1TitleEnd   = 33  // offset for track title
	v1ArtistEnd  = 63  // offset for artist name
	v1AlbumEnd   = 93  // offset for album name
	v1YearEnd    = 97  // offset where year completes
	v1CommentEnd = 125 // offset where comment completes
)

// Parse completes the actual processing of the file
// and extracts the tag information.
func (i *V1) Parse(h *os.File) {
	b := make([]byte, v1TagSize)
	h.Seek(int64(v1TagStart), -v1TagStart)
	h.Read(b)

	if frames.GetStr(b[0:v1TagStart]) != "TAG" {
		return
	}

	i.Title = frames.GetStr(b[v1TagStart:v1TitleEnd])
	i.Artist = frames.GetStr(b[v1TitleEnd:v1ArtistEnd])
	i.Album = frames.GetStr(b[v1ArtistEnd:v1AlbumEnd])
	i.Year = frames.GetInt(b[v1AlbumEnd:v1YearEnd])
	i.Comment = frames.GetStr(b[v1YearEnd:v1CommentEnd])

	i.Track = frames.GetInt([]byte{b[v1CommentEnd+1]})
	i.Genre = frames.GetInt([]byte{b[v1CommentEnd+2]})
}
