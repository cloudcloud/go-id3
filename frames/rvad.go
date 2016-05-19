package frames

import (
	"fmt"
	"math"
)

// RVAD is the frame for relative volume adjustment
type RVAD struct {
	Frame

	IncrementRight     bool    `json:"increment_right"`
	IncrementLeft      bool    `json:"increment_left"`
	IncrementRightBack bool    `json:"increment_right_back"`
	IncrementLeftBack  bool    `json:"increment_left_back"`
	IncrementCenter    bool    `json:"increment_center"`
	IncrementBass      bool    `json:"increment_bass"`
	Bytes              int     `json:"bits"`
	RelativeRight      float64 `json:"relative_right"`
	RelativeLeft       float64 `json:"relative_left"`
	PeakRight          float64 `json:"peak_right"`
	PeakLeft           float64 `json:"peak_left"`
	RelativeRightBack  float64 `json:"relative_right_back"`
	RelativeLeftBack   float64 `json:"relative_left_back"`
	PeakRightBack      float64 `json:"peak_right_back"`
	PeakLeftBack       float64 `json:"peak_left_back"`
	RelativeCenter     float64 `json:"relative_center"`
	PeakCenter         float64 `json:"peak_center"`
	RelativeBass       float64 `json:"relative_bass"`
	PeakBass           float64 `json:"peak_bass"`
}

// DisplayContent will comprehensively display known information
func (r *RVAD) DisplayContent() string {
	str := "Relative Volume Adjustment\n"

	str = fmt.Sprintf("%sRight\n\tIncrement: %t\n\tRelative Volume: %fdb\n\tPeak: %fdb\n",
		str, r.IncrementRight, r.RelativeRight, r.PeakRight)
	str = fmt.Sprintf("%sLeft\n\tIncrement: %t\n\tRelative Volume: %fdb\n\tPeak: %fdb\n",
		str, r.IncrementLeft, r.RelativeLeft, r.PeakLeft)
	str = fmt.Sprintf("%sRight Back\n\tIncrement: %t\n\tRelative Volume: %fdb\n\tPeak: %fdb\n",
		str, r.IncrementRightBack, r.RelativeRightBack, r.PeakRightBack)
	str = fmt.Sprintf("%sLeft Back\n\tIncrement: %t\n\tRelative Volume: %fdb\n\tPeak: %fdb\n",
		str, r.IncrementLeftBack, r.RelativeLeftBack, r.PeakLeftBack)
	str = fmt.Sprintf("%sCenter\n\tIncrement: %t\n\tRelative Volume: %fdb\n\tPeak: %fdb\n",
		str, r.IncrementCenter, r.RelativeCenter, r.PeakCenter)
	str = fmt.Sprintf("%sBass\n\tIncrement: %t\n\tRelative Volume: %fdb\n\tPeak: %fdb\n",
		str, r.IncrementBass, r.RelativeBass, r.PeakBass)

	return str
}

// GetName includes deprecation notice for v2.4.*
func (r *RVAD) GetName() string {
	if r.Version == Version4 {
		return fmt.Sprintf("%s (deprecated)", r.Name)
	}

	return r.Name
}

// ProcessData will handle the acquisition of all data
func (r *RVAD) ProcessData(s int, d []byte) IFrame {
	r.Size = s
	r.Data = d

	if len(d) < 3 {
		return r
	}

	// let's not just blindly check all these lengths...
	defer func() IFrame {
		if x := recover(); x != nil {
			// a panic happened!
		}
		return r
	}()

	r.IncrementRight = GetBoolBit(d[0], 7)
	r.IncrementLeft = GetBoolBit(d[0], 6)
	r.IncrementRightBack = GetBoolBit(d[0], 5)
	r.IncrementLeftBack = GetBoolBit(d[0], 4)
	r.IncrementCenter = GetBoolBit(d[0], 3)
	r.IncrementBass = GetBoolBit(d[0], 2)
	d = d[1:]

	r.Bytes = int(math.Ceil(float64(GetSize([]byte{d[0]}, 8)) / 8))
	d = d[1:]

	r.RelativeRight = float64(GetSize(d[:r.Bytes], 8)) / float64(512)
	d = d[r.Bytes:]

	r.RelativeLeft = float64(GetSize(d[:r.Bytes], 8)) / float64(512)
	d = d[r.Bytes:]

	r.PeakRight = float64(GetSize(d[:r.Bytes], 8)) / float64(512)
	d = d[r.Bytes:]

	r.PeakLeft = float64(GetSize(d[:r.Bytes], 8)) / float64(512)
	d = d[r.Bytes:]

	r.RelativeRightBack = float64(GetSize(d[:r.Bytes], 8)) / float64(512)
	d = d[r.Bytes:]

	r.RelativeLeftBack = float64(GetSize(d[:r.Bytes], 8)) / float64(512)
	d = d[r.Bytes:]

	r.PeakRightBack = float64(GetSize(d[:r.Bytes], 8)) / float64(512)
	d = d[r.Bytes:]

	r.PeakLeftBack = float64(GetSize(d[:r.Bytes], 8)) / float64(512)
	d = d[r.Bytes:]

	r.RelativeCenter = float64(GetSize(d[:r.Bytes], 8)) / float64(512)
	d = d[r.Bytes:]

	r.PeakCenter = float64(GetSize(d[:r.Bytes], 8)) / float64(512)
	d = d[r.Bytes:]

	r.RelativeBass = float64(GetSize(d[:r.Bytes], 8)) / float64(512)
	d = d[r.Bytes:]

	r.PeakBass = float64(GetSize(d[:r.Bytes], 8)) / float64(512)

	return r
}
