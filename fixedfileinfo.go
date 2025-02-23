package winfiledetails

import (
	"fmt"

	"golang.org/x/sys/windows"
)

type FixedFileInfo struct {
	FileVersion    FileVersion
	ProductVersion FileVersion
}

type FileVersion struct {
	Major uint16
	Minor uint16
	Patch uint16
	Build uint16
}

func (f FileVersion) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", f.Major, f.Minor, f.Patch, f.Build)
}

func newFixedFileInfo(vsFixedInfo *windows.VS_FIXEDFILEINFO) *FixedFileInfo {
	return &FixedFileInfo{
		FileVersion: FileVersion{
			Major: uint16(vsFixedInfo.FileVersionMS >> 16),
			Minor: uint16(vsFixedInfo.FileVersionMS & 0xffff),
			Patch: uint16(vsFixedInfo.FileVersionLS >> 16),
			Build: uint16(vsFixedInfo.FileVersionLS & 0xffff),
		},
		ProductVersion: FileVersion{
			Major: uint16(vsFixedInfo.ProductVersionMS >> 16),
			Minor: uint16(vsFixedInfo.ProductVersionMS & 0xffff),
			Patch: uint16(vsFixedInfo.ProductVersionLS >> 16),
			Build: uint16(vsFixedInfo.ProductVersionLS & 0xffff),
		},
	}
}
