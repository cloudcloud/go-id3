// Package file provides the interfacing methods to working with a file on the filesystem and pushing
// content into necessary processors for tag discovery and usage
package file

import (
	"encoding/json"
	"os"
)

// File provides the data container for an individual File
type File struct {
	Filename string `json:"filename"`
	V1       *V1    `json:"id3v1"`
	V2       *V2    `json:"id3v2"`
	Debug    bool   `json:"-"`

	fileHandle *os.File
}

const (
	readFromEnd   = 2
	readFromStart = 0
)

// Process will begin the opening and loading of File content
func (f *File) Process() *File {
	var err error
	f.fileHandle, err = os.Open(f.Filename)
	if err != nil {
		panic(err)
	}

	// run through v1
	f.V1 = &V1{}
	f.V1.Parse(f.fileHandle)

	// run through v2
	f.V2 = &V2{Debug: f.Debug}
	f.V2.Parse(f.fileHandle)

	return f
}

// CleanUp will run-through any post-processing requirements
func (f *File) CleanUp() {
	f.fileHandle.Close()
}

// PrettyPrint draws a nice representation of the file for the command line
func (f *File) PrettyPrint() string {
	a, _ := json.Marshal(f)
	return string(a)
}

// GetArtist will determine the ideal Artist string for use
func (f *File) GetArtist() string {
	return ""
}
