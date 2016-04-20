package fadb

type reference struct {
	value *interface{} // pointer to the Map value
	file  string       // full, os-aware, path to containing file
}

// TODO document the Index as being one-dimensional map to values
// with pointers for keys
var Index = make(map[string]reference)
