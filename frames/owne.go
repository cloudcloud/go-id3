package frames

import (
	"bytes"
	"fmt"
)

// OWNE is the ownership frame
type OWNE struct {
	Frame

	Currency     string `json:"currency"`
	Payed        string `json:"payed"`
	PurchaseDate string `json:"purchase_date"`
	Seller       string `json:"seller"`
}

// DisplayContent will comprehensively display known information
func (o *OWNE) DisplayContent() string {
	return fmt.Sprintf("Ownership\n\tCurrency: %s\n\tPayed: %s\n\tDate: %s\n\tSeller: %s\n",
		o.Currency,
		o.Payed,
		o.PurchaseDate,
		o.Seller)
}

// ProcessData will parse bytes for details
func (o *OWNE) ProcessData(s int, d []byte) IFrame {
	o.Size = s
	o.Data = d

	o.Utf16 = GetBoolBit(d[0], 0)
	o.Currency = GetStr(d[1:4])

	d = d[4:]
	idx := bytes.IndexByte(d, '\x00')
	o.Payed = GetStr(d[:idx])
	d = d[idx+1:]

	o.PurchaseDate = GetStr(d[:8])
	if !o.Utf16 {
		o.Seller = GetStr(d[8:])
	} else {
		o.Seller = GetUnicodeStr(d[8:])
	}

	return o
}
