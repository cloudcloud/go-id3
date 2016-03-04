// Package id3 provides standard avenues for processing of
// id3 tags within mp3 files.
package id3

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/template"
	"unicode/utf8"
)

var (
	content []byte
	dbg     bool
)

// ID3 provides a simple object for both ID3 tag types.
type ID3 struct {
	Filename string `json:"filename"`
	V1       *V1    `json:"id3v1"`
	V2       *V2    `json:"id3v2"`
}

// New will provision a new instance of the ID3 struct for
// interaction with both tag types.
func New(f string) (*ID3, error) {
	dbg = false

	i := new(ID3)
	i.Filename = f

	if _, err := os.Stat(i.Filename); err != nil {
		return nil, err
	}

	return i, nil
}

// Process will use the already provided filename and scour
// for both tag types.
func (i *ID3) Process() *ID3 {
	i.V1 = NewV1()
	i.V1.Parse(i.Filename)

	i.V2 = NewV2()
	i.V2.Parse(i.Filename)

	return i
}

// SetDebug will force the debug mode to be enabled.
func (i *ID3) SetDebug() {
	dbg = true
}

// PrettyPrint will format and dump the full output of the
// found tags to STDOUT.
func (i *ID3) PrettyPrint() {
	w := os.Stdout
	text := `All ID3v1 information:
Artist: {{.V1.Artist}}
Title: {{.V1.Title}}
Album: {{.V1.Album}}
Year: {{.V1.Year}}
Comment: {{.V1.Comment}}
Track: {{.V1.Track}}
Genre ID: {{.V1.Genre}}

All ID3v2 information:{{range .V2.Items}}
 {{.GetName}} {{.GetExplain | printf "%-40s"}} {{.GetLength | printf "%-7s"}} "{{.DisplayContent}}"{{end}}
`

	t := template.New("top")
	template.Must(t.Parse(text))
	if err := t.Execute(w, i); err != nil {
		panic(err)
	}
}

func fileToBuffer(f string, size int, begin int64) []byte {
	buffer := make([]byte, size)

	file, err := os.Open(f)
	defer file.Close()

	if err != nil {
		return buffer
	}

	if begin < 0 {
		file.Seek(begin, 2)
	} else if begin > 0 {
		file.Seek(begin, 0)
	}

	file.Read(buffer)
	return buffer
}

func getString(buf []byte) string {
	p := bytes.IndexByte(buf, 0)

	if p == -1 {
		p = len(buf)
	}

	return strings.TrimSpace(cleanUTF8(string(buf[0:p])))
}

func debug(m string) {
	if dbg {
		fmt.Println(m)
	}
}

func getInt(buf []byte) int {
	p := bytes.IndexByte(buf, 0)

	if p == -1 {
		p = len(buf)
	}

	response, _ := strconv.Atoi(cleanUTF8(string(buf[0:p])))
	return response
}

func cleanUTF8(s string) string {
	if !utf8.ValidString(s) {
		v := make([]rune, 0, len(s))
		for i, r := range s {
			if r == utf8.RuneError {
				_, size := utf8.DecodeRuneInString(s[i:])
				if size == 1 {
					continue
				}
			}
			v = append(v, r)
		}
		s = string(v)
	}

	return s
}
