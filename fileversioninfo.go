package winfileinfo

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

type fileVersionInfo struct {
	dataPointer unsafe.Pointer
	data        []byte
}

func initFileVersionInfo(filePath string) (*fileVersionInfo, error) {

	var zHandle windows.Handle
	// https://learn.microsoft.com/en-us/windows/win32/api/winver/nf-winver-getfileversioninfosizea
	size, err := windows.GetFileVersionInfoSize(filePath, &zHandle)
	if err != nil {
		return nil, fmt.Errorf("failed to get file version info size: %w", err)
	}

	// https://learn.microsoft.com/en-us/windows/win32/api/winver/nf-winver-getfileversioninfoa
	var ignoredHandle uint32 // described as Ignored in the documentation
	buffer := make([]byte, size)
	var lpData unsafe.Pointer = unsafe.Pointer(&buffer[0])
	err = windows.GetFileVersionInfo(filePath, ignoredHandle, size, lpData)
	if err != nil {
		return nil, fmt.Errorf("failed to get file version info: %w", err)
	}
	return &fileVersionInfo{
		dataPointer: lpData,
		data:        buffer,
	}, nil
}

func (f *fileVersionInfo) queryFixedFileInfo() (*windows.VS_FIXEDFILEINFO, error) {
	// https://learn.microsoft.com/en-us/windows/win32/api/verrsrc/ns-verrsrc-vs_fixedfileinfo
	// var fixedFileInfo windows.VS_FIXEDFILEINFO
	// sizePtr := unsafe.Sizeof(fixedFileInfo)
	// size := uint32(sizePtr)
	// varBuffer := make([]byte, size)
	var offset uintptr = 0
	var offsetPointer unsafe.Pointer = unsafe.Pointer(&offset)
	var length uint32
	// https://learn.microsoft.com/en-us/windows/win32/api/winver/nf-winver-verqueryvaluea
	err := windows.VerQueryValue(f.dataPointer, `\`, offsetPointer, &length)
	if err != nil {
		return nil, fmt.Errorf("failed to query file version info: %w", err)
	}
	if length == 0 {
		return nil, fmt.Errorf("no version info found")
	}
	start := offset - uintptr(f.dataPointer)
	end := start + uintptr(length)
	// my data: start 40 end: 92, data: []uint8 len: 52, cap: 1716, [189,4,239,254,0,0,1,0,2,0,6,0,88,14,93,88,0,0,10,0,88,14,93,88,63,0,0,0,0,0,0,0,4,0,4,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
	// bizon data: start: 40 end: 92, data: []uint8 len: 52, cap: 1716, [189,4,239,254,0,0,1,0,2,0,6,0,88,14,93,88,0,0,10,0,88,14,93,88,63,0,0,0,0,0,0,0,4,0,4,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
	data := f.data[start:end]
	// r := bytes.Reader(varBuffer)
	// err = binary.Read(varBuffer, binary.LittleEndian, &fixedFileInfo)
	fixedFileInfo := *(*windows.VS_FIXEDFILEINFO)(unsafe.Pointer(&data[0]))
	return &fixedFileInfo, nil
}
