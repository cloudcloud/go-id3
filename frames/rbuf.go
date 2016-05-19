package frames

import "fmt"

// RBUF provides the recommended buffer size
type RBUF struct {
	Frame

	BufferSize   int  `json:"buffer_size"`
	EmbeddedInfo bool `json:"embedded_info"`
	Offset       int  `json:"offset"`
}

// DisplayContent will comprehensively display known information
func (r *RBUF) DisplayContent() string {
	return fmt.Sprintf("Recommended buffer\n\tSize: %d\n\tInfo: %v\n\tOffset: %d\n",
		r.BufferSize,
		r.EmbeddedInfo,
		r.Offset)
}

// ProcessData will parse bytes for details
func (r *RBUF) ProcessData(s int, d []byte) IFrame {
	r.Size = s
	r.Data = d

	r.BufferSize = GetSize(d[:3], 8)
	r.EmbeddedInfo = GetBoolBit(d[4], 1)
	r.Offset = GetSize(d[5:], 8)

	return r
}
