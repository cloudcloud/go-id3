package frames

import "fmt"

// WOAF provides standard URL frame handling functionality
type WOAF struct {
	Frame

	URL string `json:"url"`
}

// DisplayContent will comprehensively display known information
func (w *WOAF) DisplayContent() string {
	return fmt.Sprintf("(%s|%s): %s\n", w.Name, w.Description, w.URL)
}

// ProcessData will handle the acquisition of all data
func (w *WOAF) ProcessData(s int, d []byte) IFrame {
	w.Size = s
	w.Data = d
	w.URL = GetStr(d)

	return w
}
