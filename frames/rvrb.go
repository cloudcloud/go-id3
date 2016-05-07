package frames

import "fmt"

// RVRB contains equalisation settings for the file
type RVRB struct {
	Frame

	ReverbLeft   []byte `json:"reverb_left"`
	ReverbRight  []byte `json:"reverb_right"`
	BouncesLeft  byte   `json:"bounces_left"`
	BouncesRight byte   `json:"bounces_right"`
	FeedbackLtol byte   `json:"feedback_ltol"`
	FeedbackLtor byte   `json:"feedback_ltor"`
	FeedbackRtor byte   `json:"feedback_rtor"`
	FeedbackRtol byte   `json:"feedback_rtol"`
	PremixLtor   byte   `json:"premix_ltor"`
	PremixRtol   byte   `json:"premix_rtol"`
}

// Init will provide the initial values
func (r *RVRB) Init(n, d string, v int) {
	r.Name = n
	r.Description = d
	r.Version = v
}

// DisplayContent will comprehensively display known information
func (r *RVRB) DisplayContent() string {
	return ""
}

// GetExplain will provide output formatting briefly
func (r *RVRB) GetExplain() string {
	return ""
}

// GetLength will provide the length of frame
func (r *RVRB) GetLength() string {
	return ""
}

// GetName will provide the Name of EQUA
func (r *RVRB) GetName() string {
	return r.Name
}

// ProcessData will parse bytes for details
func (r *RVRB) ProcessData(s int, d []byte) IFrame {
	r.Size = s
	r.Data = d

	if len(d) != 12 {
		fmt.Println("Invalid RVRB frame content length")
	} else {
		r.ReverbLeft = d[:2]
		r.ReverbRight = d[2:4]
		r.BouncesLeft = d[4]
		r.BouncesRight = d[5]
		r.FeedbackLtol = d[6]
		r.FeedbackLtor = d[7]
		r.FeedbackRtor = d[8]
		r.FeedbackRtol = d[9]
		r.PremixLtor = d[10]
		r.PremixRtol = d[11]
	}

	return r
}
