package frames

import (
	"fmt"
	"os"
)

// ASPI defines structure for a seek point index within the audio
type ASPI struct {
	Frame

	Start        int    `json:"start"`
	Length       int    `json:"length"`
	Number       int    `json:"number"`
	Bits         int    `json:"bits"`
	FractionData []byte `json:"fraction_data"`
}

// DisplayContent will comprehensively display known information
func (a *ASPI) DisplayContent() string {
	return fmt.Sprintf("Seek Points (%d) [%d:%d]\n", a.Number, a.Start, a.Length)
}

// ProcessData will handle the acquisition of all data
func (a *ASPI) ProcessData(s int, d []byte) IFrame {
	a.Size = s
	a.Data = d

	if len(a.Data) < 12 {
		return a
	}
	a.Start = GetSize(d[:4], 8)
	d = d[4:]

	a.Length = GetSize(d[:4], 8)
	d = d[4:]

	a.Number = GetSize(d[:2], 8)
	d = d[2:]

	a.Bits = GetSize([]byte{d[0]}, 8)
	d = d[1:]

	expect := a.Number * (a.Bits / 8)
	found := len(d) / (a.Bits / 8)
	if expect != a.Number || found != a.Number {
		// well, this is awkward...
		fmt.Fprintf(os.Stderr, "ASPI Frame is configured incorrectly, expected [%d], got [%d]\n", expect, found)

		a.Start = 0
		a.Length = 0
		a.Number = 0

		return a
	}
	a.FractionData = d

	return a
}
