package winfiledetails

import (
	"fmt"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

type FileDetails struct {
	CreationTime   time.Time
	LastAccessTime time.Time
	LastWriteTime  time.Time
	FileVersion    FileVersion
}

type FileVersion struct {
	Major uint16
	Minor uint16
	Patch uint16
	Build uint16
}

func getInfo(filePath string) (*FileDetails, error) {
	// Convert path to UTF-16
	utf16Path, err := windows.UTF16PtrFromString(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to convert path to UTF-16: %w", err)
	}

	// Open file with required access flags
	handle, err := windows.CreateFile(
		utf16Path,
		windows.FILE_READ_EA,
		windows.FILE_SHARE_READ,
		nil,
		windows.OPEN_EXISTING,
		windows.FILE_FLAG_BACKUP_SEMANTICS,
		0,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer windows.Close(handle)

	// Get timestamps
	creationTime, lastAccessTime, lastWriteTime, err := getTimestamps(handle)
	if err != nil {
		return nil, fmt.Errorf("failed to get file time: %w", err)
	}

	fileVersion, err := getFileVersion(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to get file version: %w", err)
	}

	var details FileDetails

	details.CreationTime = creationTime
	details.LastAccessTime = lastAccessTime
	details.LastWriteTime = lastWriteTime
	details.FileVersion = fileVersion

	return &details, nil
}

func getTimestamps(handle windows.Handle) (creationTime, lastAccessTime, lastWriteTime time.Time, err error) {
	var ctime, atime, wtime windows.Filetime
	err = windows.GetFileTime(handle, &ctime, &atime, &wtime)
	if err != nil {
		return time.Time{}, time.Time{}, time.Time{}, fmt.Errorf("failed to get file time: %w", err)
	}

	return time.Unix(0, ctime.Nanoseconds()), time.Unix(0, atime.Nanoseconds()), time.Unix(0, wtime.Nanoseconds()), nil
}

func getFileVersion(filePath string) (FileVersion, error) {
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
