package v2

// FText provides the structure for processing the TEXT frame
// type of Id3 V2. This struct is generic for most TEXT style
// frames.
type FText struct {
	Frame
}

// NewTEXT will provision a new instance of the FText struct
// for processing.
func NewTEXT(n string) *FText {
	c := new(FText)

	c.Name = n

	c.TagPreserve = false
	c.FilePreserve = false
	c.ReadOnly = false
	c.Compression = false
	c.Encryption = false
	c.Grouping = false

	return c
}

// GetExplain provides a description of the specific TEXT frame.
func (t *FText) GetExplain() string {
	a := "("

	switch t.Name {
	case "TALB":
		a += "Album/Movie/Show title"
	case "TBPM":
		a += "BPM (beats per minute)"
	case "TCOM":
		a += "Composer"
	case "TCON":
		a += "Content type"
	case "TCOP":
		a += "Copyright message"
	case "TDAT":
		a += "Date"
	case "TDLY":
		a += "Playlist delay"
	case "TENC":
		a += "Encoded by"
	case "TEXT":
		a += "Lyricist/Text writer"
	case "TFLT":
		a += "File type"
	case "TIME":
		a += "Time"
	case "TIT1":
		a += "Content group description"
	case "TIT2":
		a += "Title/songname/content description"
	case "TIT3":
		a += "Subtitle/Description refinement"
	case "TKEY":
		a += "Initial key"
	case "TLAN":
		a += "Language(s)"
	case "TLEN":
		a += "Length"
	case "TMED":
		a += "Media type"
	case "TOAL":
		a += "Original album/movie/show title"
	case "TOFN":
		a += "Original filename"
	case "TOLY":
		a += "Original lyricist(s)/text writer(s)"
	case "TOPE":
		a += "Original artist(s)/performer(s)"
	case "TORY":
		a += "Original release year"
	case "TOWN":
		a += "File owner/licensee"
	case "TPE1":
		a += "Lead performer(s)/Soloist(s)"
	case "TPE2":
		a += "Band/orchestra/accompaniment"
	case "TPE3":
		a += "Conductor/performer refinement"
	case "TPE4":
		a += "Interpreted, remixed, or otherwise modified by"
	case "TPOS":
		a += "Part of a set"
	case "TPUB":
		a += "Publisher"
	case "TRCK":
		a += "Track number/Position in set"
	case "TRDA":
		a += "Recording dates"
	case "TRSN":
		a += "Internet radio station name"
	case "TRSO":
		a += "Internet radio station owner"
	case "TSIZ":
		a += "Size"
	case "TSRC":
		a += "ISRC (international standard recording code)"
	case "TSSE":
		a += "Software/Hardware and settings used for encoding"
	case "TYER":
		a += "Year"
	}

	return a + ")"
}
