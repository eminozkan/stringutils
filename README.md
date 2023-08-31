![Version](https://img.shields.io/badge/version-0.0.1-orange.svg)
![Test](https://github.com/eminozkan/stringutils/actions/workflows/go.yml/badge.svg)
![Go](https://img.shields.io/github/go-mod/go-version/eminozkan/stringutils)
[![Go Report Card](https://goreportcard.com/badge/github.com/eminozkan/stringutils)](https://goreportcard.com/report/github.com/eminozkan/stringutils)

# stringutils

A basic stringutils package. It contains only Reverse() function.

```
func Reverse(s string) (string, error)
```

____

## Installation

Go to root folder which contains go.mod file.

```
go get github.com/eminozkan/stringutils
```

______

## Usage

```
package main

import (
	"fmt"

	"github.com/eminozkan/stringutils"
)

func main(){
	reversed, err := stringutils.Reverse("ozkan")
	if err != nil {
		log.Fatal(err)
	}    
	fmt.Println(reversed) // nakzo
}
```


____

## Contributor(s)

* [Muhammet Emin Özkan](https://github.com/eminozkan) - Creator, maintainer


## Contribute

All PR’s are welcome!

1. `fork` (https://github.com/vigo/stringutils-demo/fork)
1. Create your `branch` (`git checkout -b my-feature`)
1. `commit` yours (`git commit -am 'add some functionality'`)
1. `push` your `branch` (`git push origin my-feature`)
1. Than create a new **Pull Request**!

---

## License

This project is licensed under MIT

---

This project is intended to be a safe, welcoming space for collaboration, and
contributors are expected to adhere to the [code of conduct][coc].

[coc]: https://github.com/eminozkan/stringutils/blob/main/CODE_OF_CONDUCT.md