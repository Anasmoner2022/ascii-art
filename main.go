package main

import (
	ascii "asciiart/features"
	"os"
	"strings"
)

func main() {
	args := os.Args
	stringToArt, banner := ascii.StoreInputAndBanner(args)
	bannerFile := ascii.ReadBanner(banner)
	bannerSlice := strings.Split(bannerFile, "\n")
	splitInput := strings.Split(stringToArt, "\\n")
	outPut := ascii.DrawingInput(splitInput, bannerSlice)
	println(outPut)
}
