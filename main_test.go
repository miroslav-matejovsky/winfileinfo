package winfiledetails

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"golang.org/x/sys/windows"
)

func TestNotExistingFile(t *testing.T) {
	_, err := getInfo("C:\\Windows\\System32\\notepad2.exe")
	require.ErrorContains(t, err, "The system cannot find the file specified.")

}

func TestFile(t *testing.T) {
	d, err := getInfo("C:\\Windows\\System32\\notepad.exe")
	require.NoError(t, err)

	hundredYearsAgo := time.Now().AddDate(-100, 0, 0)
	hundredYearsInFuture := time.Now().AddDate(100, 0, 0)

	t.Logf("Creation Time: %s", d.creationTime.Format(time.RFC1123))
	require.True(t, d.creationTime.After(hundredYearsAgo))
	require.True(t, d.creationTime.Before(hundredYearsInFuture))
	
	t.Logf("Last Access Time: %s", d.lastAccessTime.Format(time.RFC1123))
	require.True(t, d.lastAccessTime.After(hundredYearsAgo))
	require.True(t, d.lastAccessTime.Before(hundredYearsInFuture))

	t.Logf("Last Write Time: %s", d.lastWriteTime.Format(time.RFC1123))
	require.True(t, d.lastWriteTime.After(hundredYearsAgo))
	require.True(t, d.lastWriteTime.Before(hundredYearsInFuture))
}

type FileDetails struct {
	creationTime, lastAccessTime, lastWriteTime time.Time
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
	var creationTime, lastAccessTime, lastWriteTime windows.Filetime
	err = windows.GetFileTime(handle, &creationTime, &lastAccessTime, &lastWriteTime)
	if err != nil {
		return nil, fmt.Errorf("failed to get file time: %w", err)
	}
	var details FileDetails

	details.creationTime = time.Unix(0, creationTime.Nanoseconds())
	details.lastAccessTime = time.Unix(0, lastAccessTime.Nanoseconds())
	details.lastWriteTime = time.Unix(0, lastWriteTime.Nanoseconds())

	return &details, nil
}
