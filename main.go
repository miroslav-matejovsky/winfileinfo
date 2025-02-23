package winfiledetails

import (
	"fmt"
	"time"
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

// String returns a string representation of the version.
func (f FileVersion) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", f.Major, f.Minor, f.Patch, f.Build)
}

func getInfo(filePath string) (*FileDetails, error) {

	ft, err := GetFileTime(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to get timestamps: %w", err)
	}

	fixedFileInfo, err := GetFixedFileInfo(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to get file version: %w", err)
	}

	var details FileDetails

	details.CreationTime = ft.CreationTime
	details.LastAccessTime = ft.LastAccessTime
	details.LastWriteTime = ft.LastWriteTime
	details.FileVersion = FileVersion{
		Major: uint16(fixedFileInfo.FileVersionMS >> 16),
		Minor: uint16(fixedFileInfo.FileVersionMS & 0xFFFF),
		Patch: uint16(fixedFileInfo.FileVersionLS >> 16),
		Build: uint16(fixedFileInfo.FileVersionLS & 0xFFFF),
	}

	return &details, nil
}
