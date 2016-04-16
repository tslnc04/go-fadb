package fadb

func Make(args []string) error {
	name := ""
	proto := ""
	if len(args) == 1 {
		name = args[0]
	}
	if len(args) == 2 {
		if protoNameX.MatchString(args[0]) {
			proto = args[0]
			name = args[1]
		} else {
			name = args[0]
			proto = args[1]
		}
	}
	print("Let's make ", name)
	if proto == "" {
		println("... after we find the prototype")
	} else {
		var path string
		path = Find(proto)
		println(" a", proto)
		println(path)
	}

	return nil
}
