package fadb

import "regexp"

// validNameChars define what characters can be in a valid data
// object and associated base file names
const validNameChars = `[a-zA-Z_][a-zA-Z0-9_-]*`

// validNameX is a regular expression matching ValidChars
var validNameX = regexp.MustCompile(`^` + validNameChars + `$`)

// protoNameChars define what looks like a Prototype name (initial cap)
const protoNameChars = `[A-Z][a-zA-Z0-9_-]*`

// protoNameX is a regular expression matching ProtoNameChars
var protoNameX = regexp.MustCompile(`^` + protoNameChars + `$`)

// defaultFADBs is a list of path strings to use by default. Thesee are
// added to the end of the logical FADBPATH (if found) and will always
// be searched.
