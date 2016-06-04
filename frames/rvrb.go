package frames

import (
	"fmt"
	"os"
)

// RVRB contains equalisation settings for the file
type RVRB struct {
	Frame

	ReverbLeft   int `json:"reverb_left"`
	ReverbRight  int `json:"reverb_right"`
	BouncesLeft  int `json:"bounces_left"`
	BouncesRight int `json:"bounces_right"`
	FeedbackLtol int `json:"feedback_ltol"`
	FeedbackLtor int `json:"feedback_ltor"`
	FeedbackRtor int `json:"feedback_rtor"`
	FeedbackRtol int `json:"feedback_rtol"`
	PremixLtor   int `json:"premix_ltor"`
	PremixRtol   int `json:"premix_rtol"`
}

// DisplayContent will comprehensively display known information
func (r *RVRB) DisplayContent() string {
	if r.ReverbLeft < 1 || r.ReverbRight < 1 {
		return "Reverb frame, no content\n"
	}

	return fmt.Sprintf("Reverb\n"+
		"\tReverb Left: %dms, Reverb Right: %dms\n"+
		"\tBounces Left: %d, Bounces Right: %d\n"+
		"\tFeedback Left to Left: %d%%, Feedback Left to Right: %d%%\n"+
		"\tFeedback Right to Right: %d%%, Feedback Right to Left: %d%%\n"+
		"\tPremix Left to Right: %d%%, Premix Right to Left: %d%%\n",
		r.ReverbLeft, r.ReverbRight,
		r.BouncesLeft, r.BouncesRight,
		r.FeedbackLtol, r.FeedbackLtor,
		r.FeedbackRtor, r.FeedbackRtol,
		r.PremixLtor, r.PremixRtol)
}

// ProcessData will parse bytes for details
func (r *RVRB) ProcessData(s int, d []byte) IFrame {
	r.Size = s
	r.Data = d

	if len(d) != 12 {
		fmt.Fprintf(os.Stderr, "Invalid RVRB frame content length")
	} else {
		r.ReverbLeft = GetSize(d[:2], 8)
		r.ReverbRight = GetSize(d[2:4], 8)
		r.BouncesLeft = GetSize([]byte{d[4]}, 8)
		r.BouncesRight = GetSize([]byte{d[5]}, 8)
		r.FeedbackLtol = GetBytePercent([]byte{d[6]}, 8)
		r.FeedbackLtor = GetBytePercent([]byte{d[7]}, 8)
		r.FeedbackRtor = GetBytePercent([]byte{d[8]}, 8)
		r.FeedbackRtol = GetBytePercent([]byte{d[9]}, 8)
		r.PremixLtor = GetBytePercent([]byte{d[10]}, 8)
		r.PremixRtol = GetBytePercent([]byte{d[11]}, 8)
	}

	return r
}
