package main

import (
	"fmt"
	"github.com/pelletier/go-toml"
)

func main() {
	tree, err := toml.LoadFile("./config/api.toml")
	if err != nil {
		panic(err)
	}
	fmt.Println(tree.Get("path").(string))
}
