package frames

// WOAF provides standard URL frame handling functionality
type WOAF struct {
	Frame

	URL string `json:"url"`
}

// Init will provide the initial values
func (w *WOAF) Init(n, d string, v int) {
	w.Name = n
	w.Description = d
	w.Version = v
}

// DisplayContent will comprehensively display known information
func (w *WOAF) DisplayContent() string {
	return ""
}

// GetExplain will provide output formatting briefly
func (w *WOAF) GetExplain() string {
	return w.Description
}

// GetLength will provide the length
func (w *WOAF) GetLength() string {
	return ""
}

// GetName will provide the Name
func (w *WOAF) GetName() string {
	return w.Name
}

// ProcessData will handle the acquisition of all data
func (w *WOAF) ProcessData(s int, d []byte) IFrame {
	w.Size = s
	w.Data = d

	// text encoding is a single byte, 0 for latin, 1 for unicode
	w.URL = GetStr(d)
	w.Frame.Cleaned = w.URL

	return w
}
