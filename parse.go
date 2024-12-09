package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

func parse(flags []string) string {
	var ret string = ""
	switch flags[0] {
	case "ls":
		mlt := false

		if len(flags) >= 3 {
			mlt = true
		} else if len(flags) == 1 {
			flags = append(flags, ".")
		}

		for i, arg := range flags[1:] {
			if arg == "&&" {
				ret += parse(flags[i + 1:])
				return ret
			}
			files, err := os.ReadDir(arg)
			if err != nil {
				log.Fatal(err)
			}
			if mlt && i > 0 {
				ret += "\n"
			}
			if mlt {
				ret += fmt.Sprintln(flags[i + 1] + ":")
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
						ret += d.Sprint(n + s)
					default:
						ret += fmt.Sprint(n + s)
					}
				}
			}
			ret += "\n"
		}

	case "rm":
		if len(flags) == 1 {
			log.Fatal("No argument given")
		}
		for i, path := range flags[1:] {
			if path == "&&" {
				ret += parse(flags[i + 1:])
				return ret
			}
			info, err := os.Stat(path)
			if err != nil {
				log.Fatal(err)
			}

			if info.IsDir() {
				fmt.Printf("Cannot remove '%s': Is a directory\n", path)
				continue
			}

			err = os.Remove(path)
			if err != nil {
				log.Fatal(err)
			}
		}

	case "pwd":
		if len(flags) > 1 && flags[1] == "&&" {
			ret += parse(flags[2:])
		}

		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		ret = wd + ret + "\n"
	}
	return ret
}
