package main

import (
	"flag"

	bimg "gopkg.in/h2non/bimg.v1"

	"github.com/tirithen/directory-thumbnails/thumbnails"
)

func main() {
	flag.Parse()
	directoryPath := flag.Arg(0)
	if directoryPath == "" {
		directoryPath = "."
	}

	imagePaths := thumbnails.GetForDirectory(directoryPath, 4)

	//buffer, err := bimg.Read("/usr/share/icons/Adwaita/512x512/status/folder-drag-accept.png")

	bottomBuffer, err := bimg.Read("folder-bottom.png")
	if err != nil {
		panic(err)
	}

	topBuffer, err := bimg.Read("folder-top.png")
	if err != nil {
		panic(err)
	}

	icon, err := thumbnails.CreateFromImages(bottomBuffer, topBuffer, imagePaths)
	if err != nil {
		panic(err)
	}

	bimg.Write("icon.png", icon)
}
