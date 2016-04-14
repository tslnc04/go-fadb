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
		println("would check")
	case "flat", "f":
		println("would flatten")
	default:
		usage()
	}
}
