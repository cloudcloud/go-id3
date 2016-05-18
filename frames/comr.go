package frames

import "bytes"

// COMR contains commerical details
type COMR struct {
	Frame

	Price          string `json:"price"`
	ValidUntil     string `json:"valid_until"`
	ContactURL     string `json:"contact_url"`
	ReceivedAs     byte   `json:"received_as"`
	SellerName     string `json:"seller_name"`
	CommercialName string `json:"commercial_name"`
	PictureMime    string `json:"picture_mime"`
	Logo           []byte `json:"logo"`
}

// Init will provide the initial values
func (c *COMR) Init(n, d string, v int) {
	c.Name = n
	c.Description = d
	c.Version = v
}

// DisplayContent will comprehensively display known information
func (c *COMR) DisplayContent() string {
	return ""
}

// GetExplain will provide output formatting briefly
func (c *COMR) GetExplain() string {
	return ""
}

// GetLength will provide the length of frame
func (c *COMR) GetLength() string {
	return ""
}

// GetName will provide the Name
func (c *COMR) GetName() string {
	return c.Name
}

// ProcessData will parse the frame bytes
func (c *COMR) ProcessData(s int, d []byte) IFrame {
	c.Size = s
	c.Data = d

	// text encoding is a single byte, 0 for latin, 1 for unicode
	if len(d) > 2 {
		// for the unicode bit
		enc := d[0]

		// pricing up first, null term
		idx := bytes.IndexByte(d[1:], '\x00')
		c.Price = GetStr(d[:idx])
		d = d[idx+1:]

		// valid until date is 8 bytes
		c.ValidUntil = GetStr(d[:8]) // date: YYYYMMDD
		d = d[8:]

		// contact url next, null term
		idx = bytes.IndexByte(d, '\x00')
		c.ContactURL = GetStr(d[:idx])

		// received as is method of song reception, single byte
		c.ReceivedAs = d[idx+1]
		d = d[idx+2:]

		// seller name, null term
		if enc == '\x00' {
			idx = bytes.IndexByte(d, '\x00')
			c.SellerName = GetStr(d[:idx])
			d = d[idx+1:]

			idx = bytes.IndexByte(d, '\x00')
			c.CommercialName = GetStr(d[:idx])
			d = d[idx+1:]
		} else if enc == '\x01' {
			c.Utf16 = true

			idx = bytes.Index(d, []byte{'\x00', '\x00'})
			c.SellerName = GetUnicodeStr(d[:idx])
			d = d[idx+2:]

			idx = bytes.Index(d, []byte{'\x00', '\x00'})
			c.CommercialName = GetUnicodeStr(d[:idx])
			d = d[idx+2:]
		}

		// media mime, null term
		idx = bytes.IndexByte(d, '\x00')
		c.PictureMime = GetStr(d[:idx])

		// media binary data
		c.Logo = d[idx+1:]
	}

	return c
}
