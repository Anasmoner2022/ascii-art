package main

import (
	ascii "asciiart/features"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	stringToArt, banner := ascii.StoreInputAndBanner(args)
	fmt.Println("String:", stringToArt)
	bannerFile := ascii.ReadBanner(banner)

}
