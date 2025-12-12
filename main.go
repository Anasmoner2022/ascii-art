package main

import (
	ascii "asciiart/features"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	stringToArt, banner := ascii.StoreInputAndBanner(args)
	bannerFile := ascii.ReadBanner(banner)
	bannerSlice := strings.Split(bannerFile, "\n")
	for h := range 9 {
		for _, w := range stringToArt {
			selectChar := int((w-32)*9) + h
			fmt.Print(bannerSlice[selectChar])
		}
		fmt.Println()
	}

}
