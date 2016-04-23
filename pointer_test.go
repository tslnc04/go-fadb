package fadb_test

import (
	"fmt"
	"github.com/fadb/go-fadb"
)

//------------------------------- Local -------------------------------

func ExampleLocalFor() {
	files := fadb.LocalFor("sample.summer.persons.robmuh")
	for _, p := range files {
		fmt.Println(p)
	}
	// Output:
	// sample/summer/persons/robmuh/_.toml
	// sample/summer/persons/robmuh/_.json
	// sample/summer/persons/robmuh.toml
	// sample/summer/persons/robmuh.json
	// sample/summer/persons/robmuh.list
	// sample/summer/persons/robmuh.array
	// sample/summer/persons/robmuh.csv
	// sample/summer/persons/robmuh.tsv
	// sample/summer/persons/robmuh.dsv
	// sample/summer/persons/robmuh.true
	// sample/summer/persons/robmuh.on
	// sample/summer/persons/robmuh.md
	// sample/summer/persons/robmuh.mdown
	// sample/summer/persons/robmuh.mdkn
	// sample/summer/persons/robmuh.mkd
	// sample/summer/persons/robmuh.mdwn
	// sample/summer/persons/robmuh.mdtxt
	// sample/summer/persons/robmuh.mdtext
	// sample/summer/persons/robmuh.markdown
	// sample/summer/persons/robmuh.mark
	// sample/summer/persons/robmuh.text
	// sample/summer/persons/robmuh.txt
	// sample/summer/persons/robmuh.rst
	// sample/summer/persons/robmuh.html
	// sample/summer/persons/robmuh.asc
	// sample/summer/persons/_.toml
	// sample/summer/persons/_.json
	// sample/summer/persons.toml
	// sample/summer/persons.json
	// sample/summer/persons.csv
	// sample/summer/persons.tsv
	// sample/summer/persons.dsv
	// sample/summer/persons.true
	// sample/summer/persons.on
	// sample/summer/_.toml
	// sample/summer/_.json
	// sample/summer.toml
	// sample/summer.json
	// sample/_.toml
	// sample/_.json
	// sample.toml
	// sample.json
	// _.toml
	// _.json
}

func ExampleLocalFoundFor() {
	files := fadb.LocalFoundFor("sample.summer.persons.betropper")
	for _, p := range files {
		fmt.Println(p)
	}
	// Output:
	// sample/summer/persons/betropper/_.toml
	// sample/summer/persons/betropper.toml
}

func ExampleLocalFileFor() {
	file := fadb.LocalFileFor("sample.summer.persons.betropper")
	fmt.Println(file)
	// Output:
	// sample/summer/persons/betropper/_.toml
}

//------------------------------- Local -------------------------------

func ExampleAllFor() {
	files := fadb.AllFor("sample.summer.persons.robmuh")
	for _, p := range files {
		fmt.Println(p)
	}
	// no output because it varies per system
}

func ExampleAllFoundFor() {
	files := fadb.AllFoundFor("sample.summer.persons.betropper")
	for _, p := range files {
		fmt.Println(p)
	}
	// Output:
	// sample/summer/persons/betropper/_.toml
	// sample/summer/persons/betropper.toml
}

func ExampleFileFor() {
	file := fadb.FileFor("sample.summer.persons.betropper")
	fmt.Println(file)
	// Output:
	// sample/summer/persons/betropper/_.toml
}
