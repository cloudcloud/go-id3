package frames

import (
	"fmt"
	"math"
)

// EQUA contains equalisation settings for the file
type EQUA struct {
	Frame

	Adjustment int `json:"adjustment"`
	Steps      []*step
}

type step struct {
	Increment  bool `json:"increment"`
	Frequency  int  `json:"frequency"`
	Adjustment int  `json:"adjustment"`
}

// DisplayContent will comprehensively display known information
func (e *EQUA) DisplayContent() string {
	return fmt.Sprintf("Adjustment: %d\nSteps: %d", e.Adjustment, len(e.Steps))
}

// GetName will provide the Name of EQUA
func (e *EQUA) GetName() string {
	if e.Version == Version4 && e.Name == "EQUA" {
		return "EQUA (Deprecated)"
	}

	return e.Name
}

// ProcessData will parse bytes for details
func (e *EQUA) ProcessData(s int, d []byte) IFrame {
	e.Size = s

	e.Adjustment = GetDirectInt(d[0])
	d = d[1:]
	e.Data = d

	extraBytes := math.Ceil(float64(e.Adjustment / 8))
	chunked := int(2 + extraBytes)

	for len(d) >= chunked {
		e.Steps = append(e.Steps, &step{
			Increment:  GetBoolBit(d[0], 8),
			Frequency:  GetBitInt(d[0], false, 7) + GetBitInt(d[1], false, 8),
			Adjustment: GetInt(d[2:int(extraBytes)]),
		})

		d = d[chunked:]
	}

	return e
}
