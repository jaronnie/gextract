package main

import (
	"github.com/jaronnie/extract"
)

func main() {
	extract.Extract("./xx.zip", extract.WithOutputPath("./"))
}
