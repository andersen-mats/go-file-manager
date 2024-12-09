package main

import (
	"flag"
)

func main() {
	flag.Parse()
	flags := flag.Args()
	// todo: remove recursion inside switch statements
	parse(flags)
}
