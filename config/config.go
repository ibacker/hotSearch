package config

import (
	"fmt"
	"github.com/pelletier/go-toml"
)

var (
	Conf = New()
)

func New() (config *toml.Tree) {
	config, err := toml.LoadFile("./config/config.toml")

	if err != nil {
		fmt.Println("TomlError", err.Error())
	}
	return
}
