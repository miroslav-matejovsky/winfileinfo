package winfiledetails

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

func GetFileVersion(filePath string) (FileVersion, error) {
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

	// https://learn.microsoft.com/en-us/windows/win32/api/winver/nf-winver-verqueryvaluea
	varBuffer := make([]byte, size)
	var lplpBuffer unsafe.Pointer = unsafe.Pointer(&varBuffer[0])
	var puLen uint32
	locale := DefaultLocales[2]
	localeStr := fmt.Sprintf("%04x%04x", locale.LangID, locale.CharsetID)
	subBlock := fmt.Sprintf("\\StringFileInfo\\%s\\FileVersion", localeStr)
	err = windows.VerQueryValue(lpData, subBlock, lplpBuffer, &puLen)
	if err != nil {
		return fileVersion, fmt.Errorf("failed to query file version info: %w", err)
	}
	if puLen == 0 {
		return fileVersion, fmt.Errorf("no version info found")
	}

	fixedFileInfo := (*windows.VS_FIXEDFILEINFO)(lplpBuffer)
	fileVersion.Major = uint16(fixedFileInfo.FileVersionMS >> 16)
	fileVersion.Minor = uint16(fixedFileInfo.FileVersionMS & 0xFFFF)
	fileVersion.Patch = uint16(fixedFileInfo.FileVersionLS >> 16)
	fileVersion.Build = uint16(fixedFileInfo.FileVersionLS & 0xFFFF)

	return fileVersion, nil
}
