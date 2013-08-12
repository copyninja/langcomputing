package charmap

import "testing"

func TestLanguageOf(t *testing.T) {

	// inArray := []rune{'ಅ', `ḥ`, `uː`, 'ഊ'}
	inArray := []rune{'ಅ', 'ഊ'}
	// outArray := []string{"kn_IN", "ISO15919", "IPA", "ml_IN"}
	outArray := []string{"kn_IN", "ml_IN"}

	for index, value := range inArray {
		if x, output := LanguageOf(value), outArray[index]; x != output {
			t.Errorf("LanguageOf(%v) = %s we need %s", value, x, output)
		}
	}

}

func TestCharCompare(t *testing.T) {
	// inArray := []interface{}{'ँ', 'అ', `aː`, 'ಆ'}
	inArray := []rune{'ँ', 'అ', 'ಆ'}
	// outArray := []interface{}{'ಁ', 'അ', 'ആ', `ā`}
	outArray := []rune{'ಁ', 'അ', 'ആ'}

	for index, value := range inArray {
		if x := CharCompare(value, outArray[index]); !x {
			t.Errorf("CharCompare(%v, %v) = %v we need %v", value, outArray[index], x, true)
		}
	}

}

func TestSoundexCode(t *testing.T) {
	inArray := []interface{}{'b', 't', 'ಇ', 'B', 'A'}
	outArray := []interface{}{'1', '3', 'B', '1', '0'}

	for index, value := range inArray {
		if x, err := SoundexCode(value); err != nil && x != outArray[index] {
			t.Errorf("SoundexCode(%v) = %v was expecting %v", value, x, outArray[index])
		}
	}
}
