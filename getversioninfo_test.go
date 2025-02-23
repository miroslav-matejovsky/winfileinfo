package winfiledetails

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetFileVersionInfo(t *testing.T) {
	_, err := GetFileVersionInfo(`C:\Windows\System32\notepad.exe`)
	require.NoError(t, err)
}
