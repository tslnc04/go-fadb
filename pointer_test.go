package pointer_test

import (
	"fmt"
)

func ExamplePathFrom() {
	fmt.Println(fadb.PathFrom("sample.pointer.right-here"))

	// Output:
	// sample/pointer.toml
}
