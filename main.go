package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	flags := flag.Args()
	fmt.Print(parse(flags))
}
