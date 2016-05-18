package frames

// RBUF provides the recommended buffer size
type RBUF struct {
	Frame

	BufferSize   int    `json:"buffer_size"`
	EmbeddedInfo bool   `json:"embedded_info"`
	Offset       []byte `json:"offset"`
}

// Init will provide the initial values
func (r *RBUF) Init(n, d string, v int) {
	r.Name = n
	r.Description = d
	r.Version = v
}

// DisplayContent will comprehensively display known information
func (r *RBUF) DisplayContent() string {
	return ""
}

// GetExplain will provide output formatting briefly
func (r *RBUF) GetExplain() string {
	return ""
}

// GetLength will provide the length of frame
func (r *RBUF) GetLength() string {
	return ""
}

// GetName will provide the Name of EQUA
func (r *RBUF) GetName() string {
	return r.Name
}

// ProcessData will parse bytes for details
func (r *RBUF) ProcessData(s int, d []byte) IFrame {
	r.Size = s
	r.Data = d

	r.BufferSize = GetInt(d[:3])
	r.EmbeddedInfo = GetBoolBit(d[4], 1)
	r.Offset = d[5:]

	return r
}
