package fadb

func Flat(db string) ([]pair, error) {
	println("will flatten it")
	blah := []pair{
		{"string", "thing"},
		{"integer", 1},
		{"float", 2.1},
		{"bool", true},
	}
	return blah, nil
}
