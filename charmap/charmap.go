package charmap

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

var devaAlphabets = make([]rune, 80)
var bengAlphabets = make([]rune, 80)
var guruAlphabets = make([]rune, 80)
var gujrAlphabets = make([]rune, 80)
var oryaAlphabets = make([]rune, 80)
var tamlAlphabets = make([]rune, 80)
var teluAlphabets = make([]rune, 80)
var kndaAlphabets = make([]rune, 80)
var mlymAlphabets = make([]rune, 80)

var charmap = map[string][]rune{
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

func initializeUnicodeRange(slice []rune, begin, length int) {
	for i := 0; i < length; i++ {
		slice[i] = rune(begin + i)
	}
}

func init() {
	for key, value := range charmap {
		initializeUnicodeRange(value, langBases[key], 80)
	}
}

func LanguageOf(char rune) string {
	for lang, langRange := range charmap {
		for _, value := range langRange {
			if value == char {
				return lang
			}
		}
	}

	return "unknown"
}

func CharCompare(char1, char2 rune) bool {
	char1Index, char2Index := -1, -1

	if char1 == char2 {
		return true
	}

	for _, slice := range charmap {
		for index, value := range slice {
			if value == char1 {
				char1Index = index
			} else if value == char2 {
				char2Index = index
			}
		}
	}

	if char1Index == char2Index {
		return true
	}

	return false
}
