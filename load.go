package fadb

func Load(db string) (interface{}, error) {
	println("will load it")
	blah := []pair{
		{"string", "thing"},
		{"integer", 1},
		{"float", 2.1},
		{"bool", true},
	}
	return blah, nil
}
