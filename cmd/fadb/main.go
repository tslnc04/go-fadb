package main

import (
	"fmt"
	"os"

	"github.com/fadb/go-fadb"
)

const usageText = `usage:
  fadb                         - interactive mode
  fadb make  [PROTOTYPE] NAME  - clones a new PROTOTYPE to NAME
  fadb make  NAME [PROTOTYPE]  - clones a new PROTOTYPE to NAME
  fadb check [FADB]            - validates FADB
  fadb flat  [FADB]            - outputs flattened key = value pairs

If ommitted FADB is assumed to be the current working directory.
`

func usage() { fmt.Print(usageText) }

func interactive() { fmt.Println("TODO interactive mode"); os.Exit(0) }

func main() {

	if len(os.Args) == 1 {
		interactive()
	}
	switch os.Args[1] {
	case "check", "c":
		_, err := fadb.Load(os.Args[2])
		if err != nil {
			panic(err)
		}
	case "make", "m":
		if len(os.Args) == 3 || len(os.Args) == 4 {
			err := fadb.Make(os.Args[2:])
			if err != nil {
				panic(err)
			}
		} else {
			usage()
			os.Exit(1)
		}
	case "flat", "f":
		flat, err := fadb.Flat(os.Args[2])
		if err != nil {
			panic(err)
		}
		println(flat)
	default:
		usage()
	}

}
