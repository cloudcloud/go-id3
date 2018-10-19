package frames

var (
	// Version23Frames defines the version 2.3 frame mapping
	// http://id3.org/id3v2.3.0
	Version23Frames = map[string]func() IFrame{
		"AENC": Gen("AENC", "Audio encryption", Version3),
		"APIC": Gen("APIC", "Attached picture", Version3),
		"COMM": Gen("COMM", "User comment", Version3),
		"COMR": Gen("COMR", "Commercial frame", Version3),
		"ENCR": Gen("ENCR", "Encryption method registration", Version3),
		"EQUA": Gen("EQUA", "Equalization", Version3),
		"ETCO": Gen("ETCO", "Event timing codes", Version3),
		"GEOB": Gen("GEOB", "General encapsulated object", Version3),
		"GRID": Gen("GRID", "Group identification registration", Version3),
		"IPLS": Gen("IPLS", "Involved people list", Version3),
		"LINK": Gen("LINK", "Linked information", Version3),
		"MCDI": Gen("MCDI", "Music CD identifier", Version3),
		"MLLT": Gen("MLLT", "MPEG location lookup table", Version3),
		"OWNE": Gen("OWNE", "Ownership frame", Version3),
		"PRIV": Gen("PRIV", "Private frame", Version3),
		"PCNT": Gen("PCNT", "Play counter", Version3),
		"POPM": Gen("POPM", "Popularimeter", Version3),
		"POSS": Gen("POSS", "Position synchronisation frame", Version3),
		"RBUF": Gen("RBUF", "Recommended buffer size", Version3),
		"RVAD": Gen("RVAD", "Relative volume adjustment", Version3),
		"RVRB": Gen("RVRB", "Reverb", Version3),
		"SYLT": Gen("SYLT", "Synchronized lyrics/text", Version3),
		"SYTC": Gen("SYTC", "Synchronized tempo codes", Version3),
		"TALB": Gen("TALB", "Album/Show/Movie title", Version3),
		"TBPM": Gen("TBPM", "BPM (beats per minute)", Version3),
		"TCOM": Gen("TCOM", "Composer", Version3),
		"TCON": Gen("TCON", "Content type", Version3),
		"TCOP": Gen("TCOP", "Copyright message", Version3),
		"TDAT": Gen("TDAT", "Date", Version3),
		"TDLY": Gen("TDLY", "Playlist delay", Version3),
		"TENC": Gen("TENC", "Encoded by", Version3),
		"TEXT": Gen("TEXT", "Lyricist/Text writer", Version3),
		"TFLT": Gen("TFLT", "File type", Version3),
		"TIME": Gen("TIME", "Time", Version3),
		"TIT1": Gen("TIT1", "Content group description", Version3),
		"TIT2": Gen("TIT2", "Title/songname/content description", Version3),
		"TIT3": Gen("TIT3", "Subtitle/Description refinement", Version3),
		"TKEY": Gen("TKEY", "Initial key", Version3),
		"TLAN": Gen("TLAN", "Language(s)", Version3),
		"TLEN": Gen("TLEN", "Length", Version3),
		"TMED": Gen("TMED", "Media type", Version3),
		"TOAL": Gen("TOAL", "Original album/movie/show title", Version3),
		"TOFN": Gen("TOFN", "Original filename", Version3),
		"TOLY": Gen("TOLY", "Original lyricist(s)/text writer(s)", Version3),
		"TOPE": Gen("TOPE", "Original artist(s)/performer(s)", Version3),
		"TORY": Gen("TORY", "Original release year", Version3),
		"TOWN": Gen("TOWN", "File owner/licensee", Version3),
		"TPE1": Gen("TPE1", "Lead performer(s)/Soloist(s)", Version3),
		"TPE2": Gen("TPE2", "Band/orchestra/accompaniment", Version3),
		"TPE3": Gen("TPE3", "Conductor/performer refinement", Version3),
		"TPE4": Gen("TPE4", "Interpreted, remixed, or otherwise modified by", Version3),
		"TPOS": Gen("TPOS", "Part of a set", Version3),
		"TPUB": Gen("TPUB", "Publisher", Version3),
		"TRCK": Gen("TRCK", "Track number/Position in set", Version3),
		"TRDA": Gen("TRDA", "Recording dates", Version3),
		"TRSN": Gen("TRSN", "Internet radio station name", Version3),
		"TRSO": Gen("TRSO", "Internet radio station owner", Version3),
		"TSIZ": Gen("TSIZ", "Size", Version3),
		"TSRC": Gen("TSRC", "ISRC (international standard recording code)", Version3),
		"TSSE": Gen("TSSE", "Software/Hardware and settings used for encoding", Version3),
		"TYER": Gen("TYER", "Year", Version3),
		"TXXX": Gen("TXXX", "User defined text information frame", Version3),
		"UFID": Gen("UFID", "Unique File Identifier", Version3),
		"USER": Gen("USER", "Terms of use", Version3),
		"USLT": Gen("USLT", "Unsynchronised lyrics/text transcription", Version3),
		"WCOM": Gen("WCOM", "Commercial information webpage", Version3),
		"WCOP": Gen("WCOP", "Copyright/legal information webpage", Version3),
		"WOAF": Gen("WOAF", "Official audio file webpage", Version3),
		"WOAR": Gen("WOAR", "Official artist/performer webpage", Version3),
		"WOAS": Gen("WOAS", "Official audio source webpage", Version3),
		"WORS": Gen("WORS", "Official internet radio station homepage", Version3),
		"WPAY": Gen("WPAY", "Payment webpage", Version3),
		"WPUB": Gen("WPUB", "Publishers official webpage", Version3),
		"WXXX": Gen("WXXX", "User defined webpage", Version3),
	}
)
