package fileversion

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVersionInfo(t *testing.T) {
	info, err := New(`C:\Windows\System32\notepad.exe`)
	require.NoError(t, err)
	version, err := info.GetProperty("FileVersion")
	require.NoError(t, err)
	assert.Equal(t, "10.0.22621.1 (WinBuild.160101.0800)", version)
}
