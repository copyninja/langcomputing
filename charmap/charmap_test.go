package charmap

import "testing"

func TestLanguageOf(t *testing.T) {

	inArray := []string{`ಅ`, `ḥ`, `uː`, `ഊ`}
	outArray := []string{"kn_IN", "ISO15919", "IPA", "ml_IN"}

	for index, value := range inArray {
		if x, output := LanguageOf(value), outArray[index]; x != output {
			t.Errorf("LanguageOf(%v) = %s we need %s", value, x, output)
		}
	}

}

func TestCharCompare(t *testing.T) {
	inArray := []string{`ँ`, `అ`, `aː`, `ಆ`}
	outArray := []string{`ಁ`, `അ`, `ആ`, `ā`}

	for index, value := range inArray {
		if x := CharCompare(value, outArray[index]); !x {
			t.Errorf("CharCompare(%v, %v) = %v we need %v", value, outArray[index], x, true)
		}
	}

}

func TestSoundexCode(t *testing.T) {
	inArray := []string{`b`, `t`, "ಇ", "B", "A"}
	outArray := []string{"1", "3", "B", "1", "0"}

	for index, value := range inArray {
		if x, err := SoundexCode(value); err != nil && x != outArray[index] {
			t.Errorf("SoundexCode(%v) = %v was expecting %v", value, x, outArray[index])
		}
	}
}
