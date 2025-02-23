package winfileinfo

import (
	"fmt"
	"time"
)

type FileDetails struct {
	CreationTime   time.Time
	LastAccessTime time.Time
	LastWriteTime  time.Time
	FileVersion    FileVersion
	ProductVersion FileVersion
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
	details.FileVersion = fixedFileInfo.FileVersion

	return &details, nil
}
