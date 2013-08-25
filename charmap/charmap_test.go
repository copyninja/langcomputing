package charmap

import "testing"

func testRuneInput(input runeSequence, output []string, t *testing.T) {

	for index, value := range input {
		if x, op := LanguageOf(value), output[index]; x != op {
			t.Errorf("LanguageOf(%v) = %s we need %s", value, x, op)
		}
	}
}

func testStringInput(input unicodeSequence, output []string, t *testing.T) {
	for index, value := range input {
		if x, op := LanguageOf(value), output[index]; x != op {
			t.Errorf("LanguageOf(%v) = %s we need %s", value, x, op)
		}
	}
}

func TestLanguageOf(t *testing.T) {

	inArray := []rune{'ಅ', 'ഊ'}
	outArray := []string{"kn_IN", "ml_IN"}

	inSplArray := []string{`ḥ`, `uː`}
	outSplArray := []string{"ISO15919", "IPA"}

	testRuneInput(inArray, outArray, t)
	testStringInput(inSplArray, outSplArray, t)

}

func TestCharCompare(t *testing.T) {
	inArray := []rune{'ँ', 'అ', 'ಆ'}
	outArray := []rune{'ಁ', 'അ', 'ആ'}

	for index, value := range inArray {
		if x := CharCompare(value, outArray[index]); !x {
			t.Errorf("CharCompare(%v, %v) = %v we need %v",
				value, outArray[index], x, true)
		}
	}

}

func TestSoundexCode(t *testing.T) {
	inArray := []interface{}{'b', 't', 'ಇ', 'B', 'A'}
	outArray := []interface{}{'1', '3', 'B', '1', '0'}

	for index, value := range inArray {
		if x, err := SoundexCode(value); err != nil && x !=
			outArray[index] {
			t.Errorf("SoundexCode(%v) = %v was expecting %v",
				value, x, outArray[index])
		}
	}
}
