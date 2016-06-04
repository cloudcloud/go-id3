package frames

import "fmt"

// SYTC defines the synchronised tempo codes
type SYTC struct {
	Frame

	Format    string   `json:"format"`
	TempoData []*tempo `json:"tempo_data"`
}

type tempo struct {
	BeatsPerMinute int `json:"beats_per_minute"`
	TimeCode       int `json:"time_code"`
}

// DisplayContent will comprehensively display known information
func (z *SYTC) DisplayContent() string {
	out := "Synchronised Tempo\n"
	for _, v := range z.TempoData {
		out = fmt.Sprintf("%s\tBPM [%d] Time Code [%d%s]\n", out, v.BeatsPerMinute, v.TimeCode, z.Format)
	}

	return out
}

// ProcessData will handle the acquisition of all data
func (z *SYTC) ProcessData(s int, d []byte) IFrame {
	z.Size = s
	z.Data = d

	f := d[0]
	z.Format = "ms"
	if f == '\x01' {
		z.Format = "mpeg"
	}
	d = d[1:]

	for {
		if len(d) < 4 {
			break
		}

		x := &tempo{BeatsPerMinute: 0}
		b := d[0]
		if b == '\xff' {
			x.BeatsPerMinute = GetSize([]byte{d[0]}, 8)
			d = d[1:]
			b = d[0]
		}

		x.BeatsPerMinute = x.BeatsPerMinute + GetSize([]byte{b}, 8)
		d = d[1:]

		x.TimeCode = GetSize(d[:3], 8)
		z.TempoData = append(z.TempoData, x)
		d = d[3:]
	}

	return z
}
