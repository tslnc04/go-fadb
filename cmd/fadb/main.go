package main

import (
	"flag"
	"fmt"

	//"github.com/fadb/go-fadb"
)

const usageText = `usage:
  fadb check FADB      - validates FADB
  fadb flat  FADB      - outputs flattened key = value pairs
`

func usage() { fmt.Print(usageText) }

func main() {
	switch flag.Arg(0) {
	case "check", "c":
		_, err = fadb.Load(flag.Arg(1))
		if err != nil {
			panic(err)
		}
	case "flat", "f":
		flat, err = fadb.Flat(flag.Arg(1))
		if err != nil {
			panic(err)
		}
		println(flat)
	default:
		usage()
	}
}
