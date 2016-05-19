package frames

import (
	"bytes"
	"fmt"
)

// CRM provides legacy meta encryption store
type CRM struct {
	Frame

	Owner       string `json:"owner"`
	Explanation string `json:"explanation"`
	Block       []byte `json:"block"`
}

// DisplayContent shows a simple overview for encryption
func (c *CRM) DisplayContent() string {
	return fmt.Sprintf("Encryption Meta\n\tOwner: %s\n\tExplanation: %s\n", c.Owner, c.Explanation)
}

// ProcessData will take some bytes and parse it
func (c *CRM) ProcessData(s int, d []byte) IFrame {
	c.Size = s
	c.Data = d

	c.Utf16 = false
	if d[0] == '\x01' {
		c.Utf16 = true
	}
	d = d[1:]

	term := []byte{'\x00'}
	if c.Utf16 {
		term = append(term, '\x00')
	}

	idx := bytes.Index(d, term)
	if c.Utf16 {
		c.Owner = GetUnicodeStr(d[:idx])
	} else {
		c.Owner = GetStr(d[:idx])
	}
	d = d[idx+len(term):]

	idx = bytes.Index(d, term)
	if c.Utf16 {
		c.Explanation = GetUnicodeStr(d[:idx])
	} else {
		c.Explanation = GetStr(d[:idx])
	}
	d = d[idx+len(term):]

	c.Block = d

	return c
}
