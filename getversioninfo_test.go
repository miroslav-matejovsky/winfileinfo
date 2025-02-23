package winfiledetails

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetFileVersionInfo(t *testing.T) {
	v, err := GetFileVersionInfo(`C:\Windows\System32\notepad.exe`)
	require.NoError(t, err)
	require.Equal(t, "10.0.22621.3572", v.String())
}
