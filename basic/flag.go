package basic

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", "$", "seperator")

func FlagMain() {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))
	fmt.Println("n:",!*n)
	if !*n {
		fmt.Println("end")
	}
}
