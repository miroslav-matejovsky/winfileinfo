package winfiledetails

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

	hundredYearsAgo := time.Now().AddDate(-100, 0, 0)
	hundredYearsInFuture := time.Now().AddDate(100, 0, 0)

	t.Logf("Creation Time: %s", d.CreationTime.Format(time.RFC1123))
	require.True(t, d.CreationTime.After(hundredYearsAgo))
	require.True(t, d.CreationTime.Before(hundredYearsInFuture))

	t.Logf("Last Access Time: %s", d.LastAccessTime.Format(time.RFC1123))
	require.True(t, d.LastAccessTime.After(hundredYearsAgo))
	require.True(t, d.LastAccessTime.Before(hundredYearsInFuture))

	t.Logf("Last Write Time: %s", d.LastWriteTime.Format(time.RFC1123))
	require.True(t, d.LastWriteTime.After(hundredYearsAgo))
	require.True(t, d.LastWriteTime.Before(hundredYearsInFuture))
}
