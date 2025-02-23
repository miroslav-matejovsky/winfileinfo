package winfileinfo

type WinFile struct {
	path string
}

func NewWinFile(path string) *WinFile {
	return &WinFile{path: path}
}
