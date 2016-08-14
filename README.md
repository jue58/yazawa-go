# YAZAWA-GO

『YAZAWA-GO』 is a imitation of [YAZAWA](https://github.com/tobynet/yazawa).

# Requirements

YAZAWA-GO requires following:
- MeCab
- Go

# Installation
Configuration for using MeCab(See also: [go-mecab](https://github.com/shogo82148/go-mecab)):
```bash
export CGO_LDFLAGS="`mecab-config --libs`"
export CGO_CFLAGS="-I`mecab-config --inc-dir`"
```

Install with `go get`:
```bash
$ go get github.com/jue58/yazawa-go
```

# Usage

Example:
```go
package main

import (
	"fmt"

	"github.com/jue58/yazawa-go/yazawa"
)

func main() {
	fmt.Println(yazawa.Convert("俺達の熱意で世界が変わる", false))
}
```

## Command

```bash
$ yazawa-go '俺達の熱意で世界が変わる'
俺たちの『NETSUI』で世界が変わる

$ yazawa-go -r '俺たちの熱意で世界が変わる'
俺たちの熱意で世界が『KAWARU』
```

# License
YAZAWA-GO is licensed under the [MIT](https://github.com/jue58/yazawa-go/blob/master/LICENSE) license.

