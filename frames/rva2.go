package frames

import (
	"bytes"
	"fmt"
	"math"
)

// RVA2 provides relative volume adjustment for the file
type RVA2 struct {
	Frame

	Identification string     `json:"identification"`
	Channels       []*channel `json:"channels"`
}

type channel struct {
	Type       string  `json:"type"`
	Adjustment float64 `json:"adjustment"`
	Volume     int     `json:"volume"`
}

var channelTypes = map[int]string{
	0: "Other",
	1: "Master volume",
	2: "Front right",
	3: "Front left",
	4: "Back right",
	5: "Back left",
	6: "Front centre",
	7: "Back centre",
	8: "Subwoofer",
}

// DisplayContent provides a clean display of key information
func (r *RVA2) DisplayContent() string {
	str := fmt.Sprintf("Relative Volume Adjustment (%s)\n", r.Identification)
	for _, v := range r.Channels {
		str = fmt.Sprintf("%s\tChannel (%s), Adjusted (%fdb), Peak (%d)\n", str, v.Type, v.Adjustment, v.Volume)
	}

	return str
}

// ProcessData will take bytes and mush into something useful
func (r *RVA2) ProcessData(s int, d []byte) IFrame {
	r.Size = s
	r.Data = d
	r.Channels = []*channel{}

	idx := bytes.IndexByte(d, '\x00')
	r.Identification = GetStr(d[:idx])
	d = d[idx+1:]

	for len(d) > 0 {
		c := &channel{}
		check := GetSize([]byte{d[0]}, 8)
		c.Type = channelTypes[check]
		if len(c.Type) < 2 {
			c.Type = channelTypes[0]
		}
		d = d[1:]

		c.Adjustment = float64(GetSize(d[:2], 8)) / float64(512)
		d = d[2:]

		bits := GetSize([]byte{d[0]}, 8)
		d = d[1:]

		bytes := int(math.Ceil(float64(bits / 8)))
		if len(d) >= bytes {
			c.Volume = GetSize(d[:bytes], 8)
		}

		if len(d) > bytes {
			d = d[bytes:]
		} else {
			d = []byte{}
		}

		r.Channels = append(r.Channels, c)
	}

	return r
}
