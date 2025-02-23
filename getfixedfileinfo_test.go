package winfiledetails

import (
	"testing"

	"github.com/bi-zone/go-fileversion"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBiZoneFileVersion(t *testing.T) {
	file := `C:\Windows\System32\notepad.exe`
	gfv, err := fileversion.New(file)
	require.NoError(t, err)
	fixedInfo := gfv.FixedInfo()
	fv := fixedInfo.FileVersion
	// the actual version should be 6.2.22621.3672, there is a bug opened for this issue
	// https://github.com/bi-zone/go-fileversion/issues/3
	assert.Equal(t, "6.2.3672.22621", fv.String())
	assert.Equal(t, uint16(6), fv.Major)
	assert.Equal(t, uint16(2), fv.Minor)
	assert.Equal(t, uint16(3672), fv.Patch)
	assert.Equal(t, uint16(22621), fv.Build)
}

func TestGetFileVersionInfo(t *testing.T) {
	file := `C:\Windows\System32\notepad.exe`
	gfv, err := fileversion.New(file)
	require.NoError(t, err)
	expected := gfv.FixedInfo()

	fi, err := GetFixedFileInfo(file)
	require.NoError(t, err)
	assert.NotEqual(t, expected.FileType, fi.Signature)
	assert.NotEqual(t, uint32(0x00000000), fi.StrucVersion)
	assert.NotEqual(t, uint32(0x00000000), fi.FileVersionMS)
	assert.NotEqual(t, uint32(0x00000000), fi.FileVersionLS)
}
