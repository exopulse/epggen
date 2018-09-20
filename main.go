package main

//go:generate go-bindata tmpl/

import (
	"github.com/exopulse/epggen/cli"

	"github.com/exopulse/epggen/generator"
)

func main() {
	cli.Execute(generator.NewGenerator(AssetLoader))
}
