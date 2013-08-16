// Soundex algorithm implementation for Indian languages.
package indicsoundex

import (
	"langcomputing/charmap"
	"strings"
)

const (
	SOUNDEX_NO_ENGLISH_COMPARE = -1
	SOUNDEX_SAME_STRING        = 0
	SOUNDEX_STRINGS_MATCH      = 1
	SOUNDEX_STRING_NOMATCH     = 2
)

func soundex(word string, length int) (string, string) {
	sndx := make([]rune, 1)
	var lang string

	// Is this the first char
	var isFc = true

	i := 0

	// Note that range splits string on Unicode code point
	for _, value := range word {
		if isFc {
			// First letter of calculated soundex should
			// be replaced with first letter of the word.
			//
			// We don't need to calculate Soundex code for
			// first letter of the word.
			isFc = false
			sndx = append(sndx[:i], value)
			lang = charmap.LanguageOf(value)
			i++
			continue
		}

		d, err := charmap.SoundexCode(value)

		// FIXME: do we need to do error handling?
		if err == nil {
			if d == '0' {
				continue
			}

			// Ignore consecutive characters
			if len(sndx) != 0 || d != sndx[len(sndx)-1] {
				sndx = append(sndx[:i], d)
				i++
			}
		}
	}

	return string(sndx), lang
}

func Calculate(word string, padding int) string {
	unicodeWord := strings.Split(word, "")

	// We need Unicode length of the word not length of UTF-8
	// encoded Unicode word.
	// .
	// Additionally unlike expected word[0] is not a Unicode
	// letter  instead first byte of UTF-8 encoded Unicode letter
	// (utf-8 encoded Unicode letter for Indian language is
	// normally  3 bytes in length). We need to reduce length by 1
	// to get last index of first Unicode character as strings are
	// 0 indexed.

	wordLength, firstCharLastIndex := len(unicodeWord), len(unicodeWord[0])-1
	result, lang := soundex(word, wordLength)

	if lang == "en_US" {
		return result
	}

	// Convert sndx a rune slice into single string and padd it
	// with `padding' number of 0
	result += strings.Repeat(`0`, padding)

	// Return the string slice 0 to padding+firstCharLastIndex
	return result[0 : padding+firstCharLastIndex]
}

func Compare(word1, word2 string) int {
	// If words are same no need of soundex calculation return
	// matching
	if word1 == word2 {
		return SOUNDEX_SAME_STRING
	}

	lang1 := charmap.LanguageOf(rune(word1[0]))
	lang2 := charmap.LanguageOf(rune(word2[0]))

	if lang1 == "en_US" && lang2 != "en_US" ||
		lang1 != "en_US" && lang2 == "en_US" {
		return SOUNDEX_NO_ENGLISH_COMPARE
	}

	soundex1 := strings.Split(Calculate(word1, 8), "")
	soundex2 := strings.Split(Calculate(word2, 8), "")

	if strings.Join(soundex1[1:], "") == strings.Join(soundex2[1:], "") {
		return SOUNDEX_STRINGS_MATCH
	}

	return SOUNDEX_STRING_NOMATCH

}
