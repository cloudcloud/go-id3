package frames

import (
	"bytes"
	"fmt"
)

// EQU2 provides equalisation details for the file
type EQU2 struct {
	Frame

	Interpolation  string   `json:"interpolation"`
	Identification string   `json:"identification"`
	Points         []*point `json:"points"`
}

type point struct {
	Frequency float64 `json:"frequency"`
	Volume    float64 `json:"volume"`
}

// DisplayContent provides a clean display of key information
func (e *EQU2) DisplayContent() string {
	str := fmt.Sprintf("Equalisation 2 (Interpolation: %s, Identification: %s)\n", e.Interpolation, e.Identification)
	for _, v := range e.Points {
		str = fmt.Sprintf("%s\tFrequency: %fhz, Volume: %fdb\n", str, v.Frequency, v.Volume)
	}

	return str
}

// ProcessData will take bytes and mush into something useful
func (e *EQU2) ProcessData(s int, d []byte) IFrame {
	e.Size = s
	e.Data = d
	e.Points = []*point{}

	erp := GetSize([]byte{d[0]}, 0)
	e.Interpolation = "Linear"
	if erp == 0 {
		e.Interpolation = "Band"
	}
	d = d[1:]

	idx := bytes.IndexByte(d, '\x00')
	e.Identification = GetStr(d[:idx])
	d = d[idx+1:]

	for len(d) >= 4 {
		p := &point{}

		p.Frequency = float64(GetSize(d[:2], 8)) / 2
		p.Volume = float64(GetSize(d[2:4], 8)) / 512

		e.Points = append(e.Points, p)
		if len(d) > 4 {
			d = d[4:]
		} else {
			d = []byte{}
		}
	}

	return e
}
