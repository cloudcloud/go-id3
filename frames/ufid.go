package frames

import "bytes"

// UFID provides descriptor of uniqueness for the file
type UFID struct {
	Frame

	Owner      string `json:"owner"`
	Identifier []byte `json:"identifier"`
}

// Init will provide the initial values
func (u *UFID) Init(n, d string, v int) {
	u.Name = n
	u.Description = d
	u.Version = v
}

// DisplayContent will comprehensively display known information
func (u *UFID) DisplayContent() string {
	return ""
}

// GetExplain will provide output formatting briefly
func (u *UFID) GetExplain() string {
	return u.Description
}

// GetLength will provide the length
func (u *UFID) GetLength() string {
	return ""
}

// GetName will provide the Name
func (u *UFID) GetName() string {
	return u.Name
}

// ProcessData will handle the acquisition of all data
func (u *UFID) ProcessData(s int, d []byte) IFrame {
	u.Size = s
	u.Data = d

	if len(d) > 2 {
		idx := bytes.IndexByte(d, '\x00')
		u.Owner = GetStr(d[:idx])
		u.Identifier = d[idx+1:]
	}

	return u
}
