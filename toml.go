package fadb

import (
	"github.com/pelletier/go-toml"
)

func loadToml(path string) map[string]interface{} {
	tree, err := toml.LoadFile(path)
	if err == nil {
		return tree.ToMap()
	}
	return nil
}
