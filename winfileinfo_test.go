package winfileinfo_test

import (
	"testing"

	wfi "github.com/miroslav-matejovsky/winfileinfo"
	"github.com/stretchr/testify/require"
)

func TestNonExistentFile(t *testing.T) {
	_, err := wfi.NewWinFileInfo("C:\\nonexistent.txt")
	require.ErrorContains(t, err, "file does not exist")
}
