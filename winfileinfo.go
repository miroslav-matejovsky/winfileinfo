package winfileinfo

type WinFileInfo struct {
	FileVersion    WinFileVersion
	ProductVersion WinFileVersion
}

type WinFileVersion struct {
	Major uint16
	Minor uint16
	Patch uint16
	Build uint16
}
