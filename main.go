package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	flags := flag.Args()
	// todo: remove recursion inside switch statements
	fmt.Print(parse(flags))
}
