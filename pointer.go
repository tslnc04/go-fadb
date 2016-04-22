package fadb

import (
	fp "path/filepath"
	"strings"
)

// PathFrom resolves an FADB pointer to its actual file system path.
func PathFrom(pointer string) (path string) {
	p := strings.Split(pointer, ".")
	//TODO first look localling in the current directory
	//TODO then look in each of the fadbpath locations
	path = fp.Join(p...)
	return
}
