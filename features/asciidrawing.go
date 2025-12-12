package asciiart

import "log"

func DrawingInput(input []string, bannerSlice []string) string {
	var convertedString string
	for i, str := range input {
		if str != "" && i != 0 {
			convertedString += "\n"
		}
		for h := 1; h < 9; h++ {
			for _, w := range str {

				if w < 32 || w > 126 {
					log.Fatalf("You Write Non-Printable Char. ")
				}
				selectChar := int((w-32))*9 + h
				convertedString += bannerSlice[selectChar]
			}
			convertedString += "\n"
		}
	}

	return convertedString
}
