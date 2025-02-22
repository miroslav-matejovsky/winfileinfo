package winfiledetails

import (
	"fmt"
	"time"

	"golang.org/x/sys/windows"
)

type FileDetails struct {
	CreationTime   time.Time
	LastAccessTime time.Time
	LastWriteTime  time.Time
	FileVersion    string
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

func getFileVersion(filePath string) (string, error) {
	var zHandle windows.Handle
	// https://learn.microsoft.com/en-us/windows/win32/api/winver/nf-winver-getfileversioninfosizea
	size, err := windows.GetFileVersionInfoSize(filePath, &zHandle)
	if err != nil {
		return "", fmt.Errorf("failed to get file version info size: %w", err)
	}

	// https://learn.microsoft.com/en-us/windows/win32/api/winver/nf-winver-getfileversioninfoa
	err = windows.GetFileVersionInfo(filePath, 0, size, nil)
	if err != nil {
		return "", fmt.Errorf("failed to get file version info: %w", err)
	}
	fmt.Printf("Size: %v\n", size)
	fmt.Printf("Handle: %v\n", zHandle)
	return "", nil
}

func getFileInformation(handle windows.Handle) error {
	var data windows.ByHandleFileInformation
	err := windows.GetFileInformationByHandle(handle, &data)
	if err != nil {
		return fmt.Errorf("failed to get file information by handle: %w", err)
	}

	// Print file attributes
	fmt.Printf("File Attributes: %v\n", data.FileAttributes)
	fmt.Printf("Volume Serial Number: %v\n", data.VolumeSerialNumber)

	fmt.Printf("In handle: %v\n", handle)

	// windows.GetFileVersionInfo("C:\\Windows\\System32\\notepad.exe", uint32(handle), 0, nil)

	return nil
}
