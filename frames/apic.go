package frames

import "bytes"

// APIC will house any specific APIC frame data
type APIC struct {
	Frame

	MimeType    string `json:"mime_type"`
	PictureType int    `json:"picture_type"`
	Image       []byte `json:"image"`
	Title       string `json:"title"`
}

// Init will provide the initial values
func (a *APIC) Init(n, d string, v int) {
	a.Name = n
	a.Description = d
	a.Version = v
}

// DisplayContent will comprehensively display known information for APIC
func (a *APIC) DisplayContent() string {
	return ""
}

// GetExplain will provide output formatting briefly for APIC
func (a *APIC) GetExplain() string {
	return ""
}

// GetLength will provide the length of the APIC frame
func (a *APIC) GetLength() string {
	return ""
}

// GetName will provide the Name of APIC
func (a *APIC) GetName() string {
	return "APIC"
}

// ProcessData grabs the meta and binary detail for the image
func (a *APIC) ProcessData(s int, d []byte) IFrame {
	a.Size = s
	a.Data = d

	// encoding is first, 1 is unicode
	enc := d[0]
	d = d[1:]

	// mime type next, null term
	idx := bytes.IndexByte(d, '\x00')
	a.MimeType = GetStr(d[:idx])

	// picture type
	a.PictureType = GetDirectInt(d[idx+1])
	d = d[idx+2:]

	// image description, null term
	if enc == '\x00' {
		idx = bytes.IndexByte(d, '\x00')
		a.Title = GetStr(d[:idx])

		// image is next now
		a.Image = d[idx+1:]
	} else if enc == '\x01' {
		a.Utf16 = true

		idx = bytes.Index(d, []byte{'\x00', '\x00'})
		a.Title = GetUnicodeStr(d[:idx])

		a.Image = d[idx+2:]
	}
	a.Size = len(a.Image)

	return a
}
