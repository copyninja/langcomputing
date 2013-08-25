/*
  Charmap package implements character map and functions

  This package implements character range slices for Indian languages
  and English. It also implements soundex codes for Indian languages
  and English. Soundex codes are accessed using function SoundexCode.

  It also  implements language detection and char compare function.

*/
package charmap

import (
	"strconv"
	"strings"
)

// Private types for rune and string slices
type unicodeSequence []string
type runeSequence []rune

type sequenceIndex interface {
	index(char interface{}) int
}

// Private map to hold all sequences
type charMap map[string]sequenceIndex

/*
  Custom Error structure for missing or unknown character
*/
type UnknownCharError struct {
	char    interface{}
	lang    string
	message string
}

func (e *UnknownCharError) Error() string {
	var returnString string

	switch e.char.(type) {
	case rune:
		returnString = strconv.QuoteRune(e.char.(rune))
	case string:
		returnString = e.char.(string)
	}

	if len(e.lang) == 0 {
		return returnString + " : " + e.message
	}

	return returnString + " " + e.message + " " + e.lang
}

// Languagewise unicode ranges
var langBases = map[string]int{
	"en_US": 0,
	"en_IN": 0,
	"hi_IN": '\u0901',
	"bn_IN": '\u0981',
	"pa_IN": '\u0a01',
	"gu_IN": '\u0a81',
	"or_IN": '\u0b01',
	"ta_IN": '\u0b81',
	"te_IN": '\u0c01',
	"kn_IN": '\u0c81',
	"ml_IN": '\u0D01',
}

// Slices to hold unicode range for each languagges
var devaAlphabets = make(runeSequence, 80)
var bengAlphabets = make(runeSequence, 80)
var guruAlphabets = make(runeSequence, 80)
var gujrAlphabets = make(runeSequence, 80)
var oryaAlphabets = make(runeSequence, 80)
var tamlAlphabets = make(runeSequence, 80)
var teluAlphabets = make(runeSequence, 80)
var kndaAlphabets = make(runeSequence, 80)
var mlymAlphabets = make(runeSequence, 80)

var enUsAlphabets = runeSequence{'a', 'b', 'c', 'd', 'e', 'f', 'g',
	'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u',
	'v', 'w', 'x', 'y', 'z'}

// Soundex values for English alphabet series
var soundexEnglish = runeSequence{'0', '1', '2', '3', '0', '1', '2',
	'0', '0', '2', '2', '4', '5', '5', '0', '1', '2', '6', '2', '3', '0',
	'1', '0', '2', '0', '2'}

// Soundex values for Indian language unicode series.
var soundexIndic = runeSequence{'0', 'N', '0', '0', 'A', 'A', 'B',
	'B', 'C', 'C', 'P', 'Q', '0', 'D', 'D', 'D', 'E', 'E', 'E', 'E', 'F',
	'F', 'F', 'F', 'G', 'H', 'H', 'H', 'H', 'G', 'I', 'I', 'I', 'I', 'J',
	'K', 'K', 'K', 'K', 'L', 'L', 'M', 'M', 'M', 'M', 'N', 'O', 'P', 'P',
	'Q', 'Q', 'Q', 'R', 'S', 'S', 'S', 'T', '0', '0', '0', '0', 'A', 'B',
	'B', 'C', 'C', 'P', 'P', 'E', 'D', 'D', 'D', 'D', 'E', 'E', 'E', '0',
	'0', '0', '0', '0', '0', '0', '0', '0', '0', 'E', '0', '0', '0', '0',
	'0', '0', '0', '0', 'P', 'Q', 'Q', 'Q', '0', '0', '0', '1', '2', '3',
	'4', '5', '6', '7', '8', '9', '0', '0', '0', '0', '0', '0', '0', '0',
	'0', '0', 'J', 'J', 'Q', 'P', 'P', 'F'}

// ISO15919 series specific to Indian languages
var iso15919IndicSeries = unicodeSequence{`m̐`, `ṁ`, `ḥ`, ``, `a`, `ā`,
	`i`, `ī`, `u`, `ū`, `ṛ`, `ḷ`, `ê`, `e`, `ē`, `ai`, `ô`, `o`, `ō`,
	`au`, `ka`, `kha`, `ga`, `gha`, `ṅa`, `ca`, `cha`, `ja`, `jha`, `ña`,
	`ṭa`, `ṭha`, `ḍa`, `ḍha`, `ṇa`, `ta`, `tha`, `da`, `dha`, `na`, `ṉa`,
	`pa`, `pha`, `ba`, `bha`, `ma`, `ya`, `ra`, `ṟa`, `la`, `ḷa`, `ḻa`,
	`va`, `śa`, `ṣa`, `sa`, `ha`, ``, ``, ``, `'`, `ā`, `i`, `ī`, `u`,
	`ū`, `ṛ`, `ṝ`, `ê`, `e`, `ē`, `ai`, `ô`, `o`, `ō`, `au`, ``, ``, ``,
	`oṃ`, ``, ``, ``, ``, ``, ``, ``, `qa`, `ḵẖa`, `ġ`, `za`, `ṛa`, `ṛha`,
	`fa`, `ẏa`, `ṝ`, `ḹ`, `ḷ`, `ḹ`, `.`, `..`, `0`, `1`, `2`, `3`, `4`,
	`5`, `6`, `7`, `8`, `9`, `…`, ``, ``, ``, ``, ``, ``, ``}

// IPA series specific for Indian languages
var ipaIndicSeries = unicodeSequence{`m`, `m`, ``, ``, `ə`, `aː`, `i`,
	`iː`, `u`, `uː`, `r̩`, `l̩`, `æ`, `e`, `eː`, `ɛː`, `ɔ`, `o`, `oː`, `ow`,
	`kə`, `kʰə`, `gə`, `gʱə`, `ŋə`, `ʧə`, `ʧʰə`, `ʤə`, `ʤʱə`, `ɲə`, `ʈə`,
	`ʈʰə`, `ɖə`, `ɖʱə`, `ɳə`, `t̪ə`, `t̪ʰə`, `d̪ə`, `d̪ʱə`, `n̪ə`, `nə`, `pə`,
	`pʰə`, `bə`, `bʱə`, `mə`, `jə`, `ɾə`, `rə`, `lə`, `ɭə`, `ɻə`, `ʋə`,
	`ɕə`, `ʂə`, `sə`, `ɦə`, ``, ``, ``, `ഽ`, `aː`, `i`, `iː`, `u`, `uː`,
	`r̩`, `l̩`, `e`, `eː`, `ɛː`, `ɔ`, `o`, `oː`, `ow`, ``, ``, ``, ``, ``,
	``, ``, ``, ``, ``, `ow`, ``, ``, ``, ``, ``, ``, ``, ``, `r̩ː`, `l̩ː`,
	``, ``, ``, ``, `0`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `൰`,
	``, ``, ``, ``, ``, ``, ``, ``, ``, ``, ``, ``, ``, ``, ``}

// Map to hold rune sequence of each languages
var langMap = charMap{
	"hi_IN": devaAlphabets,
	"bn_IN": bengAlphabets,
	"pa_IN": guruAlphabets,
	"gu_IN": gujrAlphabets,
	"or_IN": oryaAlphabets,
	"ta_IN": tamlAlphabets,
	"te_IN": teluAlphabets,
	"kn_IN": kndaAlphabets,
	"ml_IN": mlymAlphabets,
}

func initializeUnicodeRange(slice runeSequence, begin int) {
	for i := 0; i < len(slice); i++ {
		slice[i] = rune(begin + i)
	}
}

func init() {
	for key, value := range langMap {
		initializeUnicodeRange(value.(runeSequence), langBases[key])
	}

	langMap["soundex_en"] = soundexEnglish
	langMap["soundex_in"] = soundexIndic
	langMap["ISO15919"] = iso15919IndicSeries
	langMap["IPA"] = ipaIndicSeries
	langMap["en_US"] = enUsAlphabets
}

func (r unicodeSequence) index(char interface{}) int {
	for i, value := range r {
		if value == char {
			return i
		}
	}

	return -1
}

func (r runeSequence) index(char interface{}) int {

	for i, value := range r {
		if value == char.(rune) {
			return i
		}
	}

	return -1
}

/*
  Language of returns the language of given character. Function
  accepts string or rune as argument. If language can not be detected
  function returns string "unknown"
*/
func LanguageOf(char interface{}) string {
	for lang, langRange := range langMap {
		if langRange.index(char) != -1 {
			return lang
		}
	}
	// Still not found then something wrong
	return "unknown"
}

/*
  CharCompare compares the given character or string literal to see if
  they are similar from different languages.

  Function accepts string or rune this is because IPA and ISO15919
  series are multibyte sequence and can not be represented as single
  rune.

  Another characteristic of Indian language is all languages have
  similar character set this is the basis of this function if both
  character are in 2 Indian languages but at similar Unicode code
  point function returns true otherwise returns false.
*/
func CharCompare(char1, char2 interface{}) bool {

	if char1 == char2 {
		return true
	}

	char1Index := langMap[LanguageOf(char1)].index(char1)
	char2Index := langMap[LanguageOf(char2)].index(char2)

	if char1Index == char2Index {
		return true
	}

	return false
}

/*
  SoundexCode returns the soundex code for given character function
  returns rune code and UnknownCharError.

  For the sake of allowing IPA and ISO15919 series this function also
  accepts either rune or string.

  If code for char is not found "0" is returned which will be ignored
  by Soundex calculation algorithm and error initiated using composite
  literal which can be used to print error message.
*/
func SoundexCode(char interface{}) (rune, error) {
	var lang string

	switch char.(type) {
	case string:
		char = strings.ToLower(char.(string))
	}

	if lang = LanguageOf(char); lang != "unknown" {
		if charIndex := langMap[lang].index(char); charIndex != -1 {
			var sequence runeSequence

			switch lang {
			case "en_US":
				sequence = langMap["soundex_en"].(runeSequence)
			default:
				sequence = langMap["soundex_in"].(runeSequence)
			}
			return sequence[charIndex], nil
		}
		return '0', &UnknownCharError{char, lang, "not found"}
	}
	return '0', &UnknownCharError{char, lang, "unknown language"}
}
