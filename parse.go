package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

func parse(flags []string) string {
	var ret string = ""
	if len(flags) == 0 {
		return ret
	}
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
				ret += "\n"
			}
			if mlt {
				ret += fmt.Sprintln(flags[i + 1] + ":")
			}
			for j, file := range files {
				n := file.Name()
				var s string
				if j != len(files) - 1 {
					s = "  "
				} else {
					s = ""
				}
				if string(n[0]) == "." {
					continue
				}
				switch file.IsDir() {
				case true:
					d := color.New(color.FgBlue, color.Bold)
					ret += d.Sprint(n) + s
				default:
					ret += n + s
				}
			}
			ret += "\n"
		}

	case "rm":
		if len(flags) == 1 {
			log.Fatal("No argument given")
		}
		for _, path := range flags[1:] {
			info, err := os.Stat(path)
			if err != nil {
				log.Fatal(err)
			}

			if info.IsDir() {
				fmt.Printf("Cannot remove '%s': Is a directory\n", path)
				continue
			} else {
				err = os.Remove(path)
				if err != nil {
					log.Fatal(err)
				}
			}
		}

	case "pwd":
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		ret += wd + "\n"
	case "rmdir":
		for _, flag := range flags[1:] {
			info, err := os.Stat(flag)
			if err != nil {
				log.Fatal(err)
			}

			if info.IsDir() {
				dir, err := os.ReadDir(flag)
				if err != nil {
					log.Fatal(err)
				}
				if len(dir) > 0 {
					fmt.Printf("Cannot remove '%s': Not empty\n", flag)
				} else {
					os.RemoveAll(flag)
				}
			}
		}
	}
	return ret
}
