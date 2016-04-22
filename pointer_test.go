package pointer_test

import (
	"fmt"

	"github.com/fadb/go-fadb"
)

func ExamplePathFrom() {
	fmt.Println(fadb.PathFrom("sample.pointer.right-here"))

	// Output:
	// sample/pointer.toml
}
