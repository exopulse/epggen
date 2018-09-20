package main

import (
	"github.com/exopulse/epggen/generator"
)

type assetLoader struct{}

// AssetLoader loads assets (templates) from bindata.go.
var AssetLoader generator.TemplateLoader = new(assetLoader)

func (*assetLoader) LoadTemplate(name string) (string, error) {
	bytes, err := Asset(name)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
