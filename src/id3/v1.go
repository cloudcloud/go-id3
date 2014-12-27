package id3

import (
	"fmt"
	"strconv"
)

const (
	tag_size  = 128
	tag_start = 3

	title_end   = 33
	artist_end  = 63
	album_end   = 93
	year_end    = 97
	comment_end = 125
)

type Id3V1 struct {
	Artist  string `json:"artist"`
	Title   string `json:"title"`
	Album   string `json:"album"`
	Year    int    `json:"year"`
	Comment string `json:"comment"`
	Track   int    `json:"track"`
	Genre   int    `json:"genre"`
}

func NewV1() *Id3V1 {
	i := new(Id3V1)

	return i
}

func (i *Id3V1) Parse(f string) {
	b := fileToBuffer(f, tag_size, -tag_size)

	if getString(b[0:tag_start]) != "TAG" {
		return
	}

	i.Title = getString(b[tag_start:title_end])
	i.Artist = getString(b[title_end:artist_end])
	i.Album = getString(b[artist_end:album_end])
	i.Year = getInt(b[album_end:year_end])
	i.Comment = getString(b[year_end:comment_end])

	i.Track, _ = strconv.Atoi(fmt.Sprintf("%d", b[comment_end+1]))
	i.Genre, _ = strconv.Atoi(fmt.Sprintf("%d", b[comment_end+2]))
}
