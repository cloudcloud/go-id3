package frames

import "bytes"

// OWNE is the ownership frame
type OWNE struct {
	Frame

	Encoding     byte   `json:"encoding"`
	Currency     string `json:"currency"`
	Payed        string `json:"payed"`
	PurchaseDate string `json:"purchase_date"`
	Seller       string `json:"seller"`
}

// Init will provide the initial values
func (o *OWNE) Init(n, d string, v int) {
	o.Name = n
	o.Description = d
	o.Version = v
}

// DisplayContent will comprehensively display known information
func (o *OWNE) DisplayContent() string {
	return ""
}

// GetExplain will provide output formatting briefly
func (o *OWNE) GetExplain() string {
	return ""
}

// GetLength will provide the length of frame
func (o *OWNE) GetLength() string {
	return ""
}

// GetName will provide the Name of EQUA
func (o *OWNE) GetName() string {
	return o.Name
}

// ProcessData will parse bytes for details
func (o *OWNE) ProcessData(s int, d []byte) IFrame {
	o.Size = s
	o.Data = d

	o.Encoding = d[0]
	o.Currency = GetStr(d[1:4])

	d = d[4:]
	idx := bytes.IndexByte(d, '\x00')
	o.Payed = GetStr(d[:idx])
	d = d[idx+1:]

	o.PurchaseDate = GetStr(d[:8])
	if o.Encoding == '\x00' {
		o.Seller = GetStr(d[8:])
	} else if o.Encoding == '\x01' {
		o.Utf16 = true

		o.Seller = GetUnicodeStr(d[8:])
	}

	return o
}
