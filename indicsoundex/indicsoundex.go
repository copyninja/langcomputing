/*
 Soundex algorithm implementation for English and Indian languages

 Soundex algorithm was originally developed only to work with English
 languages but this package implements soundex algorithm for Indian
 language which is not same as English.

 Original algorithm was designed by Santhosh Thottingal and more
 information on this algorithm can be found at his website.
 http://thottingal.in/blog/2009/07/26/indicsoundex/

*/package indicsoundex

import (
	"langcomputing/charmap"
	"strings"
)

/*
  Return values defined here
*/
const (
	SOUNDEX_NO_ENGLISH_COMPARE       = -1 // We can't implement English to Indic comparision
	SOUNDEX_SAME_STRING              = 0  // Input strings are same
	SOUNDEX_STRINGS_SOUNDS_ALIKE     = 1  // Input strings sounds alike but of different languages
	SOUNDEX_STRINGS_SOUNDS_DIFFERENT = 2  // Strings doesn't sound alike
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

/*
  Calculate returns the calculated soundex for given string if the
  language of given string is English padding is ignored and only
  calculated soundex is returned. If string is Indian language first
  the soundex is calculated and the string is padded with padding
  number of 0's and 0-padding of resulting string is returned
*/
func Calculate(word string, padding int) string {
	// Lets first split word to get unicode sequences in string
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

	// We don't need to padd English words soundex value
	if lang == "en_US" {
		return result
	}

	// Convert sndx a rune slice into single string and padd it
	// with `padding' number of 0
	result += strings.Repeat(`0`, padding)

	// Return the string slice 0 to padding+firstCharLastIndex
	return result[0 : padding+firstCharLastIndex]
}

/*
  Compare compares word1 and word2 to check if they sounds alike. If
  they sounds alike SOUNDEX_STRINGS_SOUNDS_ALIKE is returned. If same string
  is given as input, function simply returns SOUNDEX_SAME_STRING
  without calculating soundex for strings.

  If one of the string is of English language, function returns
  SOUNDEX_NO_ENGLISH_COMPARE as there is no way to compare an English
  and Indian language string.

  If strings does not match SOUNDEX_STRINGS_SOUNDS_DIFFERENT is returned.
*/
func Compare(word1, word2 string) int {
	// If words are same no need of soundex calculation return
	// matching
	if word1 == word2 {
		return SOUNDEX_SAME_STRING
	}

	// Check language of first char to find out if this string is
	// Englihs language. Indexing first char works only for English language
	// as UTF-8 representation for English is 1byte where as Indian language
	// it will be 3 bytes
	lang1 := charmap.LanguageOf(rune(word1[0]))
	lang2 := charmap.LanguageOf(rune(word2[0]))

	if lang1 == "en_US" && lang2 != "en_US" ||
		lang1 != "en_US" && lang2 == "en_US" {
		return SOUNDEX_NO_ENGLISH_COMPARE
	}

	// Purposefully splitting the string just to make sure we get
	// unicode sequences properly. In case of Indian language single first
	// letter will take 3 bytes.
	soundex1 := strings.Split(Calculate(word1, 8), "")
	soundex2 := strings.Split(Calculate(word2, 8), "")

	// Ignoring first letter if rest of string matches then word1 and word2
	// sounds alike
	if strings.Join(soundex1[1:], "") == strings.Join(soundex2[1:], "") {
		return SOUNDEX_STRINGS_SOUNDS_ALIKE
	}

	return SOUNDEX_STRINGS_SOUNDS_DIFFERENT

}
