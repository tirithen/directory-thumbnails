package main

import (
	"flag"
	"fmt"

	"./thumbnails"
)

func main() {
	flag.Parse()
	directoryPath := flag.Arg(0)
	imagePaths := thumbnails.GetForDirectory(directoryPath, 4)
	fmt.Print("imagePaths")
	fmt.Println(imagePaths)
}
