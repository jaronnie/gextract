package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jaronnie/extract"
	"github.com/spf13/pflag"
)

var (
	OutputDir string
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Printf("empty source file\n")
		return
	}
	bindFlag()

	err := extract.Extract(args[1], extract.WithOutputPath(OutputDir))
	if err != nil {
		fmt.Println(err)
	}
}

func bindFlag() {
	pflag.StringVarP(&OutputDir,
		"output",
		"p",
		"",
		"set output dir")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
}
