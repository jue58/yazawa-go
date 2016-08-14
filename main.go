package main

import (
	"flag"
	"fmt"

	"github.com/jue58/yazawa-go/yazawa"
)

func main() {
	rp := flag.Bool("r", false, "Random mode. Default: false")
	flag.Parse()
	text := flag.Arg(0)
	fmt.Println(yazawa.Convert(text, *rp))
}
