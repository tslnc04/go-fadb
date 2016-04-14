package main

import (
	"fmt"
	"os"

	"github.com/fadb/go-fadb"
)

const usageText = `usage:
  fadb check FADB      - validates FADB
  fadb flat  FADB      - outputs flattened key = value pairs
`

func usage() { fmt.Print(usageText) }

func main() {
	switch os.Args[1] {
	case "check", "c":
		_, err := fadb.Load(os.Args[2])
		if err != nil {
			panic(err)
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
