# plist - A pure Go property list transcoder [![coverage report](https://gitlab.howett.net/go/plist/badges/main/coverage.svg)](https://gitlab.howett.net/go/plist/commits/main)
## INSTALL
```
$ go get github.com/lilili87222/plist
```

## FEATURES
* Supports encoding/decoding property lists (Apple XML, Apple Binary, OpenStep and GNUStep) from/to arbitrary Go types

## USE
```go
package main
import (
	"github.com/lilili87222/plist"
	"os"
)
func main() {
	encoder := plist.NewEncoder(os.Stdout)
	encoder.Encode(map[string]string{"hello": "world"})
}
```


this code is adapted from https://github.com/DHowett/go-plist

and added some features 
1.support xml version 1.1