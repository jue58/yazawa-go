# YAZAWA-GO

『YAZAWA-GO』 is a imitation of [YAZAWA](https://github.com/tobynet/yazawa).

# Requirements

YAZAWA-GO requires following:
- MeCab
- Go

# Installation

# Usage

example:
```
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

```
$ yazawa-go '俺達の熱意で世界が変わる'
俺たちの『NETSUI』で世界が変わる

$ yazawa-go -r '俺たちの熱意で世界が変わる'
俺たちの熱意で世界が『KAWARU』
```

# License
YAZAWA-GO is licensed under the [MIT](https://github.com/jue58/yazawa-go/blob/master/LICENSE) license.

