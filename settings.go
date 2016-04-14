package fadb

import "regexp"

// ValidNameChars defined what characters can be in a valid data
// object and associated base file names
const ValidNameChars = `[a-zA-Z_][a-zA-Z0-9_-]`

// ValidNamePattern is a regular expression matching ValidNameChars
var ValidNamePattern = regexp.MustCompile(`^` + ValidNameChars + `$`)
