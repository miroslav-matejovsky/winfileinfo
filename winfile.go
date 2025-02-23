// Package winfileinfo provides utilities for retrieving file information on Windows systems.
package winfileinfo

import (
	"fmt"
	"os"

	"golang.org/x/sys/windows"
)

// WinFile represents a file on the Windows filesystem.
type WinFile struct {
	path string
}

// NewWinFile creates a new WinFile for the given path.
// It returns an error if the file does not exist or if there is an error checking the file.
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

// GetFileTime retrieves the file time information for the file.
// It returns a WinFileTime struct containing the file time information.
func (wf *WinFile) GetFileTime() (*WinFileTime, error) {
	return wf.getFileTime()
}

// GetFileInfo retrieves the file version information for the file.
// It returns a WinFileInfo struct containing the file version information.
func (wf *WinFile) GetFileInfo() (*WinFileInfo, error) {
	ffi, err := wf.GetFixedFileInfo()
	if err != nil {
		return nil, err
	}
	return newWinFileInfo(ffi), nil
}

// GetFixedFileInfo retrieves the fixed file information for the file.
// It returns a windows.VS_FIXEDFILEINFO struct containing the fixed file information.
func (wf *WinFile) GetFixedFileInfo() (*windows.VS_FIXEDFILEINFO, error) {
	winver, err := initWinVer(wf.path)
	if err != nil {
		return nil, err
	}
	return winver.queryFixedFileInfo()
}
