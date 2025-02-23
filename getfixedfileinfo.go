package winfileinfo

func GetFixedFileInfo(filePath string) (*FixedFileInfo, error) {
	fvi, err := initFileVersionInfo(filePath)
	if err != nil {
		return nil, err
	}
	ffi, err := fvi.queryFixedFileInfo()
	if err != nil {
		return nil, err
	}
	return newFixedFileInfo(ffi), nil
}
