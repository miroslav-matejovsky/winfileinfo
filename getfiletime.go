package winfileinfo

import (
	"fmt"
	"time"

	"golang.org/x/sys/windows"
)

type FileTime struct {
	CreationTime   time.Time
	LastAccessTime time.Time
	LastWriteTime  time.Time
}

func GetFileTime(filePath string) (*FileTime, error) {
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
	creationTime, lastAccessTime, lastWriteTime, err := getFileTime(handle)
	if err != nil {
		return nil, fmt.Errorf("failed to get timestamps: %w", err)
	}
	return &FileTime{
		CreationTime:   creationTime,
		LastAccessTime: lastAccessTime,
		LastWriteTime:  lastWriteTime,
	}, nil
}

func getFileTime(handle windows.Handle) (creationTime, lastAccessTime, lastWriteTime time.Time, err error) {
	var ctime, atime, wtime windows.Filetime
	err = windows.GetFileTime(handle, &ctime, &atime, &wtime)
	if err != nil {
		return time.Time{}, time.Time{}, time.Time{}, fmt.Errorf("failed to get file time: %w", err)
	}

	return time.Unix(0, ctime.Nanoseconds()), time.Unix(0, atime.Nanoseconds()), time.Unix(0, wtime.Nanoseconds()), nil
}
