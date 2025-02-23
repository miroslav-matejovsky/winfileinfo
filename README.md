# winfileinfo

`winfileinfo` is a Go module for retrieving file version and file time information on Windows.

## Installation

To install the module, use `go get`:

```sh
go get github.com/miroslav-matejovsky/winfileinfo
```

## Usage

### Retrieving File Version Information

You can retrieve the file version information using the `WinFileInfo` struct.

```go
package main

import (
    "fmt"
    "log"

    "github.com/miroslav-matejovsky/winfileinfo"
)

func main() {
    file := `C:\Windows\System32\notepad.exe`
    wf, err := winfileinfo.NewWinFile(file)
    if err != nil {
        log.Fatalf("Error creating WinFile: %v", err)
    }

    fi, err := wf.GetFileInfo()
    if err != nil {
        log.Fatalf("Error getting file info: %v", err)
    }

    fmt.Printf("File Version: %s\n", fi.FileVersion)
    fmt.Printf("Product Version: %s\n", fi.ProductVersion)
}
```

### Retrieving File Time Information

You can retrieve the file time information using the `WinFileTime` struct.

```go
package main

import (
    "fmt"
    "log"

    "github.com/miroslav-matejovsky/winfileinfo"
)

func main() {
    file := `C:\Windows\System32\notepad.exe`
    wf, err := winfileinfo.NewWinFile(file)
    if err != nil {
        log.Fatalf("Error creating WinFile: %v", err)
    }

    ft, err := wf.GetFileTime()
    if err != nil {
        log.Fatalf("Error getting file time: %v", err)
    }

    fmt.Printf("Creation Time: %s\n", ft.CreationTime)
    fmt.Printf("Last Access Time: %s\n", ft.LastAccessTime)
    fmt.Printf("Last Write Time: %s\n", ft.LastWriteTime)
}
```

## Testing

To run the tests, use the `go test` command:

```sh
go test ./...
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
