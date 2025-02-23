package winfileinfo

import "time"

type WinFileTime struct {
	CreationTime   time.Time
	LastAccessTime time.Time
	LastWriteTime  time.Time
}
