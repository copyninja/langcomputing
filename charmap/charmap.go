package charmap

import "strings"

// Private types for rune and string slices
type runeSequence []rune
type stringSequence []string

type sequenceIndex interface {
	index(char rune) int
}

// Private map to hold all sequences
type charMap map[string]sequenceIndex

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

// Soundex values for English alphabet series
var soundexEnglish = runeSequence{'0', '1', '2', '3', '0', '1', '2', '0', '0', '2', '2', '4', '5', '5', '0', '1', '2', '6', '2', '3', '0', '1', '0', '2', '0', '2'}

// Soundex values for Indian language unicode series.
var soundexIndic = runeSequence{'0', 'N', '0', '0', 'A', 'A', 'B', 'B', 'C', 'C', 'P', 'Q', '0', 'D', 'D', 'D', 'E', 'E', 'E', 'E', 'F', 'F', 'F', 'F', 'G', 'H', 'H', 'H', 'H', 'G', 'I', 'I', 'I', 'I', 'J', 'K', 'K', 'K', 'K', 'L', 'L', 'M', 'M', 'M', 'M', 'N', 'O', 'P', 'P', 'Q', 'Q', 'Q', 'R', 'S', 'S', 'S', 'T', '0', '0', '0', '0', 'A', 'B', 'B', 'C', 'C', 'P', 'P', 'E', 'D', 'D', 'D', 'D', 'E', 'E', 'E', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', 'E', '0', '0', '0', '0', '0', '0', '0', '0', 'P', 'Q', 'Q', 'Q', '0', '0', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', 'J', 'J', 'Q', 'P', 'P', 'F'}

// ISO15919 series specific to Indian languages
var iso15919IndicSeries = stringSequence{`m̐`, `ṁ`, `ḥ`, ``, `a`, `ā`, `i`, `ī`, `u`, `ū`, `ṛ`, `ḷ`, `ê`, `e`, `ē`, `ai`, `ô`, `o`, `ō`, `au`, `ka`, `kha`, `ga`, `gha`, `ṅa`, `ca`, `cha`, `ja`, `jha`, `ña`, `ṭa`, `ṭha`, `ḍa`, `ḍha`, `ṇa`, `ta`, `tha`, `da`, `dha`, `na`, `ṉa`, `pa`, `pha`, `ba`, `bha`, `ma`, `ya`, `ra`, `ṟa`, `la`, `ḷa`, `ḻa`, `va`, `śa`, `ṣa`, `sa`, `ha`, ``, ``, ``, `'`, `ā`, `i`, `ī`, `u`, `ū`, `ṛ`, `ṝ`, `ê`, `e`, `ē`, `ai`, `ô`, `o`, `ō`, `au`, ``, ``, ``, `oṃ`, ``, ``, ``, ``, ``, ``, ``, `qa`, `ḵẖa`, `ġ`, `za`, `ṛa`, `ṛha`, `fa`, `ẏa`, `ṝ`, `ḹ`, `ḷ`, `ḹ`, `.`, `..`, `0`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `…`, ``, ``, ``, ``, ``, ``, ``}

// IPA series specific for Indian languages
var ipaIndicSeries = stringSequence{`m`, `m`, ``, ``, `ə`, `aː`, `i`, `iː`, `u`, `uː`, `r̩`, `l̩`, `æ`, `e`, `eː`, `ɛː`, `ɔ`, `o`, `oː`, `ow`, `kə`, `kʰə`, `gə`, `gʱə`, `ŋə`, `ʧə`, `ʧʰə`, `ʤə`, `ʤʱə`, `ɲə`, `ʈə`, `ʈʰə`, `ɖə`, `ɖʱə`, `ɳə`, `t̪ə`, `t̪ʰə`, `d̪ə`, `d̪ʱə`, `n̪ə`, `nə`, `pə`, `pʰə`, `bə`, `bʱə`, `mə`, `jə`, `ɾə`, `rə`, `lə`, `ɭə`, `ɻə`, `ʋə`, `ɕə`, `ʂə`, `sə`, `ɦə`, ``, ``, ``, `ഽ`, `aː`, `i`, `iː`, `u`, `uː`, `r̩`, `l̩`, `e`, `eː`, `ɛː`, `ɔ`, `o`, `oː`, `ow`, ``, ``, ``, ``, ``, ``, ``, ``, ``, ``, `ow`, ``, ``, ``, ``, ``, ``, ``, ``, `r̩ː`, `l̩ː`, ``, ``, ``, ``, `0`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `൰`, ``, ``, ``, ``, ``, ``, ``, ``, ``, ``, ``, ``, ``, ``, ``}

// Map to hold rune sequence of each languages
var langMap = charMap{
	"hi_IN":      devaAlphabets,
	"bn_IN":      bengAlphabets,
	"pa_IN":      guruAlphabets,
	"gu_IN":      gujrAlphabets,
	"or_IN":      oryaAlphabets,
	"ta_IN":      tamlAlphabets,
	"te_IN":      teluAlphabets,
	"kn_IN":      kndaAlphabets,
	"ml_IN":      mlymAlphabets,
	"soundex_en": soundexEnglish,
	"soundex_in": soundexIndic,
	"IPA":        ipaIndicSeries,
	"ISO15919":   iso15919IndicSeries,
}

// Map to hold ISO and IPA string sequences
// var iso15919IPACharmap = stringMap{
// 	"ISO15919": iso15919IndicSeries,
// 	"IPA":      ipaIndicSeries,
// }

func initializeUnicodeRange(slice runeSequence, begin, length int) {
	for i := 0; i < length; i++ {
		slice[i] = rune(begin + i)
	}
}

func init() {
	for key, value := range langMap {
		if _, ok := value.(runeSequence); ok {
			initializeUnicodeRange(value.(runeSequence), langBases[key], len(value.(runeSequence)))
		}
	}
}

func (r runeSequence) index(char rune) int {
	for i, value := range r {
		if value == char {
			return i
		}
	}

	return -1
}

func (s stringSequence) index(char rune) int {
	for i, value := range s {
		if strings.ContainsRune(value, char) {
			return i
		}
	}

	return -1
}

func LanguageOf(char rune) string {
	for lang, langRange := range langMap {
		if langRange.index(char) != -1 {
			return lang
		}
	}
	// Still not found then something wrong
	return "unknown"
}

func CharCompare(char1, char2 rune) bool {

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
