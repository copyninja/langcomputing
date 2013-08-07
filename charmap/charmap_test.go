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
