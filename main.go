package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

func main() {
	flag.Parse()

	flags := flag.Args()
	switch flags[0] {
	case "ls":
		mlt := false

		if len(flags) >= 3 {
			mlt = true
		} else if len(flags) == 1 {
			flags = append(flags, ".")
		}

		for i, arg := range flags[1:] {
			files, err := os.ReadDir(arg)
			if err != nil {
				log.Fatal(err)
			}
			if mlt && i > 0 {
				fmt.Println()
			}
			if mlt {
				fmt.Println(flags[i + 1] + ":")
			}
			for _, file := range files {
				n := file.Name()
				s := " "
				if file == files[len(files) - 1] {
					s = ""
				}
				switch string(n[0]) {
				case ".":
					continue
				default:
					switch file.IsDir() {
					case true:
						d := color.New(color.FgBlue, color.Bold)
						d.Print(n + s)
					default:
						fmt.Print(n + s)
					}
				}
			}
			fmt.Println()
		}

	case "rm":
		if len(flags) == 1 {
			log.Fatal("No argument given")
		}
		for _, flag := range flags[1:] {
			path := flag
			err := os.Remove(path)
			if err != nil {
				log.Fatal(err)
			}
		}

	case "pwd":
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(wd)
	}
}
