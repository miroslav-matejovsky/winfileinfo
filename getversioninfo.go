package winfiledetails

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

func GetFileVersionInfo(filePath string) (FileVersion, error) {
	fileVersion := FileVersion{}
	var zHandle windows.Handle
	// https://learn.microsoft.com/en-us/windows/win32/api/winver/nf-winver-getfileversioninfosizea
	size, err := windows.GetFileVersionInfoSize(filePath, &zHandle)
	if err != nil {
		return fileVersion, fmt.Errorf("failed to get file version info size: %w", err)
	}

	// https://learn.microsoft.com/en-us/windows/win32/api/winver/nf-winver-getfileversioninfoa
	var ignoredHandle uint32 // described as Ignored in the documentation
	buffer := make([]byte, size)
	var lpData unsafe.Pointer = unsafe.Pointer(&buffer[0])
	err = windows.GetFileVersionInfo(filePath, ignoredHandle, size, lpData)
	if err != nil {
		return fileVersion, fmt.Errorf("failed to get file version info: %w", err)
	}



	return fileVersion, nil
}
