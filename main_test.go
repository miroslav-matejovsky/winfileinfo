package winfileinfo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNotExistingFile(t *testing.T) {
	_, err := getInfo("C:\\Windows\\System32\\notepad2.exe")
	require.ErrorContains(t, err, "The system cannot find the file specified.")
}

func TestFile(t *testing.T) {
	d, err := getInfo("C:\\Windows\\System32\\notepad.exe")
	require.NoError(t, err)

	tenYearsAgo := time.Now().AddDate(-10, 0, 0)
	now := time.Now()

	t.Logf("Creation Time: %s", d.CreationTime.Format(time.RFC1123))
	require.True(t, d.CreationTime.After(tenYearsAgo))
	require.True(t, d.CreationTime.Before(now))

	t.Logf("Last Access Time: %s", d.LastAccessTime.Format(time.RFC1123))
	require.True(t, d.LastAccessTime.After(tenYearsAgo))
	require.True(t, d.LastAccessTime.Before(now))

	t.Logf("Last Write Time: %s", d.LastWriteTime.Format(time.RFC1123))
	require.True(t, d.LastWriteTime.After(tenYearsAgo))
	require.True(t, d.LastWriteTime.Before(now))

	t.Logf("File Version: %d.%d.%d.%d", d.FileVersion.Major, d.FileVersion.Minor, d.FileVersion.Patch, d.FileVersion.Build)
	// require.Greater(t, d.FileVersion.Major, uint16(0))
	// require.Greater(t, d.FileVersion.Minor, uint16(0))
	// require.Greater(t, d.FileVersion.Patch, uint16(0))
	// require.Greater(t, d.FileVersion.Build, 0)
}
