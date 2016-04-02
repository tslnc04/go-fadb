package main

import (
	"log"
	"os"
	"strconv"

	"github.com/fadb/go-fadb"
)

func main() {
	var err error
	days := 365
	start := ""
	if len(os.Args) > 1 {
		start, err = os.Args[2]
		if err != nil {
			log.Fatal(err)
		}
	}
	if len(os.Args) > 2 {
		days, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
	}
	fadb.MakeDays(days, start)
}
