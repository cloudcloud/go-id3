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
	ID3V1    *ID3V1 `json:"id3v1"`
	ID3V2    *ID3V2 `json:"id3v2"`
}

// New will provision a new instance of the ID3 struct for
// interaction with both tag types.
func New(f string) *ID3 {
	dbg = false

	i := new(ID3)
	i.Filename = f

	return i
}

// Process will use the already provided filename and scour
// for both tag types.
func (i *ID3) Process() *ID3 {
	i.ID3V1 = NewV1()
	i.ID3V1.Parse(i.Filename)

	i.ID3V2 = NewV2()
	i.ID3V2.Parse(i.Filename)

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
Artist: {{.ID3V1.Artist}}
Title: {{.ID3V1.Title}}
Album: {{.ID3V1.Album}}
Year: {{.ID3V1.Year}}
Comment: {{.ID3V1.Comment}}
Track: {{.ID3V1.Track}}
Genre ID: {{.ID3V1.Genre}}

All ID3v2 information:{{range .ID3V2.Items}}
 {{.Name}} ({{.Explain | printf "%-20s"}})[{{.Length | printf "%-3s"}}] "{{.Content}}"{{end}}
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
