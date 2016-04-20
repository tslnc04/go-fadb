package fadb

import (
	"github.com/pelletier/go-toml"
)

// LoadToml loads a TomlTree and strips back to simple
// map[string]interface{} for more generic usage
func LoadToml(path string) map[string]interface{} {
	tree, err := toml.LoadFile(path)
	if err != nil {
		panic(err)
	}
	return tree.ToMap()
}
