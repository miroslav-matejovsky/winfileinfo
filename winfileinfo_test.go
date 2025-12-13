package winfileinfo_test

import (
	"os"
	"testing"

	wfi "github.com/miroslav-matejovsky/winfileinfo"
	"github.com/stretchr/testify/require"
)

func TestNonExistentFile(t *testing.T) {
	_, err := wfi.NewWinFileInfo("C:\\nonexistent.txt")
	require.ErrorContains(t, err, "file does not exist")
}

func TestNotClosedFileReturnError(t *testing.T) {
	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "testfile_*.txt")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name()) // Clean up

	// Write something to it to ensure it's a valid file
	_, err = tmpFile.WriteString("test content")
	require.NoError(t, err)

	// Do not close the file yet
	// Attempt to create WinFileInfo from the opened file
	_, err = wfi.NewWinFileInfo(tmpFile.Name())
	require.Error(t, err) // Should return an error because the file is open

	// Now close the file
	tmpFile.Close()
}
