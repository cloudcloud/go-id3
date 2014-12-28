package id3

import (
	"bytes"
	"os"
	"strconv"
	"strings"
	"text/template"
	"unicode/utf8"
)

var (
	content []byte
)

type Id3 struct {
	Filename string `json:"filename"`
	Id3V1    *Id3V1 `json:"id3v1"`
	Id3V2    *Id3V2 `json:"id3v2"`
}

func New(f string) *Id3 {
	i := new(Id3)
	i.Filename = f

	return i
}

func (i *Id3) Process() *Id3 {
	i.Id3V1 = NewV1()
	i.Id3V1.Parse(i.Filename)

	i.Id3V2 = NewV2()
	i.Id3V2.Parse(i.Filename)

	return i
}

func (i *Id3) PrettyPrint() {
	w := os.Stdout
	text := `All ID3v1 information:
Artist: {{.Id3V1.Artist}}
Title: {{.Id3V1.Title}}
Album: {{.Id3V1.Album}}
Year: {{.Id3V1.Year}}
Comment: {{.Id3V1.Comment}}
Track: {{.Id3V1.Track}}
Genre ID: {{.Id3V1.Genre}}

All ID3v2 information:{{range .Id3V2.Items}}
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
