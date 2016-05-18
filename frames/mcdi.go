package frames

// MCDI is the music cd identifier frame
type MCDI struct {
	Frame

	DiscHeader []byte   `json:"disc_header"`
	Tracks     [][]byte `json:"tracks"`
}

// Init will provide the initial values
func (m *MCDI) Init(n, d string, v int) {
	m.Name = n
	m.Description = d
	m.Version = v
}

// DisplayContent will comprehensively display known information
func (m *MCDI) DisplayContent() string {
	return ""
}

// GetExplain will provide output formatting briefly
func (m *MCDI) GetExplain() string {
	return m.Description
}

// GetLength will provide the length
func (m *MCDI) GetLength() string {
	return ""
}

// GetName will provide the Name
func (m *MCDI) GetName() string {
	return m.Name
}

// ProcessData will handle the acquisition of all data
func (m *MCDI) ProcessData(s int, d []byte) IFrame {
	m.Size = s
	m.Data = d

	// text encoding is a single byte, 0 for latin, 1 for unicode
	if len(d) > 19 {
		m.DiscHeader = d[:4]
		d = d[4:]

		for i := 0; i < len(d); i += 8 {
			if len(d) < 8 {
				break
			}

			m.Tracks = append(m.Tracks, d[:8])
			d = d[8:]
		}
	}

	return m
}
