package frames

import (
	"fmt"
	"math"
)

// MCDI is the music cd identifier frame
type MCDI struct {
	Frame

	DiscHeader []byte   `json:"disc_header"`
	Tracks     [][]byte `json:"tracks"`
}

// DisplayContent will comprehensively display known information
func (m *MCDI) DisplayContent() string {
	str := fmt.Sprintf("MCD ID (%x)\n", m.DiscHeader)
	for k, v := range m.Tracks {
		str = fmt.Sprintf("%s\tTrack %d: %x\n", str, k, v)
	}

	return str
}

// ProcessData will handle the acquisition of all data
func (m *MCDI) ProcessData(s int, d []byte) IFrame {
	m.Size = s
	m.Data = d
	m.Tracks = make([][]byte, int(math.Ceil(float64(s/8))))

	if len(d) > 19 {
		m.DiscHeader = d[:4]
		d = d[4:]

		count := 0
		for i := 8; i <= len(d); {
			if len(d) == 8 {
				m.Tracks[count] = d
			}

			if len(d) <= 8 {
				break
			}

			m.Tracks[count] = d[:i]
			d = d[i:]
			count++
		}
	}

	return m
}
