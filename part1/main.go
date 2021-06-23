package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(CapitalizeEveryThirdAlphanumericChar("Aspiration"))
}

//CapitalizeEverythirdAlphanumericChar capitalizes only every third alphanumeric character of a string,
//Characters other than each third should be downcased,
//For the purposes of counting and capitalizing every three characters, consider only alphanumeric
//characters, ie [a-z][A-Z][0-9].
//Example Input: Aspiration
//Example Output: asPirAtiOns
func CapitalizeEveryThirdAlphanumericChar(s string) string {

	if len(s) < 1 {
		return s
	}

	var count int32
	runeSlice := []rune(s)

	for index, char := range runeSlice {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			count++
			if count == 3 {
				runeSlice[index] = unicode.ToUpper(char)
				count = 0
			} else {
				runeSlice[index] = unicode.ToLower(char)
			}
		}
	}

	return string(runeSlice)
}
