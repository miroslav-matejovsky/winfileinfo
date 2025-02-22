package winfiledetails

// LangID is a Windows language identifier. Could be one of the codes listed in
// `langID` section of
// https://docs.microsoft.com/en-us/windows/win32/menurc/versioninfo-resource
type LangID uint16

// CharsetID is character-set identifier. Could be one of the codes listed in
// `charsetID` section of
// https://docs.microsoft.com/en-us/windows/win32/menurc/versioninfo-resource
type CharsetID uint16

// Locale defines a pair of a language ID and a charsetID. It can be either any
// combination of predefined LangID and CharsetID or crafted manually suing
// values from https://docs.microsoft.com/en-us/windows/win32/menurc/versioninfo-resource
type Locale struct {
	LangID    LangID
	CharsetID CharsetID
}

// The package defines a list of most commonly used LangID and CharsetID
// constant. More combinations you can find in windows docs or at
// https://godoc.org/github.com/josephspurrier/goversioninfo#pkg-constants
const (
	LangEnglish = LangID(0x049)

	CSAscii   = CharsetID(0x04e4)
	CSUnicode = CharsetID(0x04B0)
	CSUnknown = CharsetID(0x0000)
)

// DefaultLocales is a list of default Locale values. It's used as a fallback
// in a calls with automatic locales detection.
//
//nolint:gochecknoglobals
var DefaultLocales = []Locale{
	{
		LangID:    LangEnglish,
		CharsetID: CSAscii,
	},
	{
		LangID:    LangEnglish,
		CharsetID: CSUnicode,
	},
	{
		LangID:    LangEnglish,
		CharsetID: CSUnknown,
	},
}
