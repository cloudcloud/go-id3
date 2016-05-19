package frames

import "fmt"

// MLLT defines the mpeg location lookup table
type MLLT struct {
	Frame

	FramesBetween       []byte `json:"frames_between"`
	BytesBetween        []byte `json:"bytes_between"`
	MillisecondsBetween []byte `json:"milliseconds_between"`
	BitsForBytes        byte   `json:"bits_for_bytes"`
	BitsForMilliseconds byte   `json:"bits_for_milliseconds"`
	Deviations          []byte `json:"deviations"`
}

// DisplayContent will comprehensively display known information
func (m *MLLT) DisplayContent() string {
	return fmt.Sprintf("MPEG Lookup\n\tFrames: %d\n\tBytes: %d\n\tMilliseconds: %d\n",
		m.FramesBetween,
		m.BytesBetween,
		m.MillisecondsBetween)
}

// ProcessData will handle the acquisition of all data
func (m *MLLT) ProcessData(s int, d []byte) IFrame {
	m.Size = s
	m.Data = d

	if m.Size > 10 {
		m.FramesBetween = d[:2]
		m.BytesBetween = d[2:5]
		m.MillisecondsBetween = d[5:8]
		m.BitsForBytes = d[8]
		m.BitsForMilliseconds = d[9]
		m.Deviations = d[10:]
	}

	return m
}
