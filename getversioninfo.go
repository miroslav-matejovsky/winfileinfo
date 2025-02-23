package winfiledetails

import (
	"fmt"
	"log"
	"unsafe"

	"github.com/Microsoft/go-winio"
	"golang.org/x/sys/windows"
)

func GetFileVersionInfo(filePath string) ([]byte, error) {
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

	// https://learn.microsoft.com/en-us/windows/win32/api/winver/nf-winver-verqueryvaluea

	attrs, err := winio.DecodeExtendedAttributes(buffer)
	if err != nil {
		return nil, fmt.Errorf("failed to decode extended attributes: %w", err)
	}
	for _, attr := range attrs {
		log.Printf("Attribute: %s", attr)
	}

	return buffer, nil
}
