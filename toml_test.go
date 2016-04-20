package fadb

import (
	fp "path/filepath"
	"testing"
)

var example = fp.Join("samples", "example.toml")

func TestLoadToml(t *testing.T) {
	m := LoadToml(example)
	println(m["title"].(string))
}
