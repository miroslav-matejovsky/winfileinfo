package winfiledetails

import "time"

// FileEA represents an extended attribute of a Windows file
type FileEA struct {
	Name  string
	Value []byte
	Flags uint8
}

// FileInformation provides methods for retrieving Windows file information
type FileInformation struct {
	retryAttempts int
	retryDelay    time.Duration
}

// NewFileInformation returns a new FileInformation instance with default settings
func NewFileInformation() *FileInformation {
	return &FileInformation{
		retryAttempts: 3,
		retryDelay:    500 * time.Millisecond,
	}
}

// GetExtendedAttributes retrieves extended attributes for a file
func (fi *FileInformation) GetExtendedAttributes(path string) ([]FileEA, error) {
	// Implementation with retry logic and error handling
	return nil, nil
}

// SetExtendedAttributes sets extended attributes for a file
func (fi *FileInformation) SetExtendedAttributes(path string, eas []FileEA) error {
	// Implementation with retry logic and error handling
	return nil
}
