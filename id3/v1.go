package id3

import (
	"fmt"
	"strconv"
)

const (
	tagSize  = 128
	tagStart = 3

	titleEnd   = 33
	artistEnd  = 63
	albumEnd   = 93
	yearEnd    = 97
	commentEnd = 125
)

// ID3V1 provides the full structure for processing
// and working with ID3 V1 tags.
type ID3V1 struct {
	Artist  string `json:"artist"`
	Title   string `json:"title"`
	Album   string `json:"album"`
	Year    int    `json:"year"`
	Comment string `json:"comment"`
	Track   int    `json:"track"`
	Genre   int    `json:"genre"`
}

// NewV1 will provision an instance of ID3V1.
func NewV1() *ID3V1 {
	i := new(ID3V1)

	return i
}

// Parse completes the actual processing of the file
// and extracts the tag information.
func (i *ID3V1) Parse(f string) {
	b := fileToBuffer(f, tagSize, -tagSize)

	if getString(b[0:tagStart]) != "TAG" {
		return
	}

	i.Title = getString(b[tagStart:titleEnd])
	i.Artist = getString(b[titleEnd:artistEnd])
	i.Album = getString(b[artistEnd:albumEnd])
	i.Year = getInt(b[albumEnd:yearEnd])
	i.Comment = getString(b[yearEnd:commentEnd])

	i.Track, _ = strconv.Atoi(fmt.Sprintf("%d", b[commentEnd+1]))
	i.Genre, _ = strconv.Atoi(fmt.Sprintf("%d", b[commentEnd+2]))
}
