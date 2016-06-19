// Package id3 provides the interfacing methods to working with a file on the filesystem and pushing
// content into necessary processors for tag discovery and usage
package id3

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/cloudcloud/go-id3/frames"
	"gopkg.in/yaml.v2"
)

// File provides the data container for an individual File
type File struct {
	Filename string `json:"filename"`
	V1       *V1    `json:"id3v1"`
	V2       *V2    `json:"id3v2"`
	Debug    bool   `json:"-"`

	fileHandle frames.FrameFile
}

// Version is an interface for individual version implementations
type Version interface {
	Parse(frames.FrameFile) error
}

const (
	readFromEnd   = 2
	readFromStart = 0
)

// Process will begin the opening and loading of File content
func (f *File) Process(h frames.FrameFile) *File {
	f.fileHandle = h

	// run through v1
	f.V1 = &V1{Debug: f.Debug}
	f.V1.Parse(f.fileHandle)

	// run through v2
	f.V2 = &V2{Debug: f.Debug}
	f.V2.Parse(f.fileHandle)

	return f
}

// PrettyPrint draws a nice representation of the file for the command line
func (f *File) PrettyPrint(o io.Writer, format string) {
	switch format {
	case "text":
		fmt.Fprintf(o, "Artist: %s\n", f.GetArtist())
		fmt.Fprintf(o, "Album:  %s\n", f.GetAlbum())

	case "yaml":
		out, _ := yaml.Marshal(f)
		fmt.Fprintf(o, string(out))

	case "raw":

	case "json":
		fallthrough
	default:
		e := json.NewEncoder(o)
		e.Encode(f)
	}
}

// GetArtist will determine the ideal Artist string for use
func (f *File) GetArtist() string {
	a := f.V2.GetArtist()
	if len(a) < 1 {
		a = f.V1.Artist
	}

	return a
}

// GetAlbum will determine the ideal Album string for use
func (f *File) GetAlbum() string {
	a := f.V2.GetAlbum()
	if len(a) < 1 {
		a = f.V1.Album
	}

	return a
}
