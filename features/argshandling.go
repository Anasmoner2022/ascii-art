package asciiart

func storeInputAndBanner(Args []string) (string, string) {
	var stringToArt string
	var bannerStyle = "standard"
	switch len(Args) {
	case 3:
		stringToArt = Args[0]
		bannerStyle = Args[1]
	case 2:
		stringToArt = Args[0]
	default:
		stringToArt = ""
	}
	return stringToArt, bannerStyle
}
