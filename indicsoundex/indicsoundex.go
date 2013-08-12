package indicsoundex

import (
	"langcomputing/charmap"
	"strings"
	// "strconv"
	"fmt"
)

func Calculate(word string, padding int) string {
	unicodeWord := strings.Split(word, "")
	wordLength, fcLength := len(unicodeWord), len(unicodeWord[0])-1

	sndx := make([]rune, wordLength)

	var isFc = true

	i := 0

	for _, value := range word {
		if isFc {
			isFc = false
			sndx[i] = value
			i++
			continue
		}

		d, err := charmap.SoundexCode(value)
		if err == nil {
			if len(sndx) != 0 || d != sndx[len(sndx)-1] {
				sndx[i] = d
				i++
			}
		}
	}

	result := string(sndx) + strings.Repeat(`0`, padding)
	fmt.Println(result)
	return string(result[0 : padding+fcLength])
}
