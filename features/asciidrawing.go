package asciiart

import "log"

func DrawingInput(input []string, bannerSlice []string) string {
	var convertedString string
	for i, str := range input {
		if str != "" && i != 0 {
			convertedString += "\n"
		}
		for h := range 9 {
			for _, w := range str {

				if w < 32 || w > 126 {
					log.Fatalf("You Write Non-Printable Char. ")
				}
				if w == '\n' {
					convertedString += "\n"
					continue
				}
				selectChar := int((w-32))*9 + h
				convertedString += bannerSlice[selectChar]
			}
			convertedString += "\n"
		}
	}

	return convertedString
}
