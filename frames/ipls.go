package frames

import (
	"bytes"
	"fmt"
	"sort"
)

// IPLS provides the involved people list frame
type IPLS struct {
	Frame

	People map[string]string `json:"people"`
}

// DisplayContent will comprehensively display known information
func (i *IPLS) DisplayContent() string {
	out := fmt.Sprintf("Involved People:\n")
	for k, v := range i.People {
		out = fmt.Sprintf("%s\t%s: %s\n", out, k, v)
	}

	return out
}

// GetName includes deprecation notice for v2.4.*
func (i *IPLS) GetName() string {
	if i.Version == Version4 {
		return fmt.Sprintf("%s (deprecated)", i.Name)
	}

	return i.Name
}

// ProcessData will handle the acquisition of all data
func (i *IPLS) ProcessData(s int, d []byte) IFrame {
	i.Size = s
	i.Data = d

	i.People = map[string]string{}
	k := []string{}
	t := map[string]string{}

	if d[0] == '\x01' {
		i.Utf16 = true
	}
	d = d[1:]

	// loop through lines, should be even numbered
	for len(d) > 2 {
		if !i.Utf16 {
			idx := bytes.IndexByte(d, '\x00')
			name := GetStr(d[:idx])
			d = d[idx+LengthStandard:]

			idx = bytes.IndexByte(d, '\x00')
			if idx == -1 {
				value := GetStr(d)
				k = append(k, name)
				t[name] = value

				break
			}
			value := GetStr(d[:idx])
			k = append(k, name)
			t[name] = value

			d = d[idx+LengthStandard:]
		} else {
			idx := bytes.Index(d, []byte{'\x00', '\x00'})
			name := GetUnicodeStr(d[:idx])
			d = d[idx+LengthUnicode:]

			idx = bytes.Index(d, []byte{'\x00', '\x00'})
			if idx == -1 {
				value := GetUnicodeStr(d)
				k = append(k, name)
				t[name] = value

				break
			}
			value := GetUnicodeStr(d[:idx])
			k = append(k, name)
			t[name] = value

			d = d[idx+LengthUnicode:]
		}
	}

	sort.Strings(k)
	for _, x := range k {
		if len(x) > 0 {
			i.People[x] = t[x]
		}
	}

	return i
}
