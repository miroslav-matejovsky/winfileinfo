package winfiledetails

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

func GetFixedFileInfo(filePath string) (*windows.VS_FIXEDFILEINFO, error) {
	var zHandle windows.Handle
	// https://learn.microsoft.com/en-us/windows/win32/api/winver/nf-winver-getfileversioninfosizea
	size, err := windows.GetFileVersionInfoSize(filePath, &zHandle)
	if err != nil {
		return nil, fmt.Errorf("failed to get file version info size: %w", err)
	}

	// https://learn.microsoft.com/en-us/windows/win32/api/winver/nf-winver-getfileversioninfoa
	var ignoredHandle uint32 // described as Ignored in the documentation
	buffer := make([]byte, size)
	var lpData unsafe.Pointer = unsafe.Pointer(&buffer[0])
	err = windows.GetFileVersionInfo(filePath, ignoredHandle, size, lpData)
	if err != nil {
		return nil, fmt.Errorf("failed to get file version info: %w", err)
	}
	return queryFixedFileInfo(lpData)
}

// https://learn.microsoft.com/en-us/windows/win32/api/verrsrc/ns-verrsrc-vs_fixedfileinfo
func queryFixedFileInfo(lpData unsafe.Pointer) (*windows.VS_FIXEDFILEINFO, error) {
	var fixedFileInfo windows.VS_FIXEDFILEINFO
	size := unsafe.Sizeof(fixedFileInfo)
	varBuffer := make([]byte, size)
	var lpBuffer unsafe.Pointer = unsafe.Pointer(&varBuffer[0])
	var puLen uint32
	err := windows.VerQueryValue(lpData, "\\", lpBuffer, &puLen)
	if err != nil {
		return nil, fmt.Errorf("failed to query file version info: %w", err)
	}
	if puLen == 0 {
		return nil, fmt.Errorf("no version info found")
	}
	fixedFileInfo = *(*windows.VS_FIXEDFILEINFO)(lpBuffer)
	return &fixedFileInfo, nil
}
