package fadb_test

import (
	"fmt"
	"github.com/fadb/go-fadb"
)

func ExampleFileFrom() {
	fmt.Println(fadb.FileFrom("sample.persons.robmuh"))

	// Output:
	// sample/persons/robmuh/_.toml
}
