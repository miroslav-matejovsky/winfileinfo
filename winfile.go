package winfileinfo

import (
	"fmt"
	"os"

	"golang.org/x/sys/windows"
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

func (wf *WinFile) GetFileInfo() (*WinFileInfo, error) {
	ffi, err := wf.GetFixedFileInfo()
	if err != nil {
		return nil, err
	}
	return newWinFileInfo(ffi), nil
}

func (wf *WinFile) GetFixedFileInfo() (*windows.VS_FIXEDFILEINFO, error) {
	winver, err := initWinVer(wf.path)
	if err != nil {
		return nil, err
	}
	return winver.queryFixedFileInfo()
}
