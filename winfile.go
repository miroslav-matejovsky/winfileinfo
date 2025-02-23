package winfileinfo

import (
	"fmt"
	"os"
)

type WinFile struct {
	path string
}

func NewWinFile(path string) (*WinFile, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("file does not exist: %s", path)
	}
	if err != nil {
		return nil, fmt.Errorf("error checking file: %s", err)
	}
	return &WinFile{path: path}, nil
}
