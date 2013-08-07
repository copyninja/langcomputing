package charmap

import "testing"

func TestLanguageOf(t *testing.T) {

	inArray := []rune{'ಅ', 'ḥ', 'ɭ', '\u0d0a'}
	outArray := []string{"kn_IN", "ISO15919", "IPA", "ml_IN"}

	for index, value := range inArray {
		if x, output := LanguageOf(value), outArray[index]; x != output {
			t.Errorf("LanguageOf(%v) = %s we need %s", value, x, output)
		}
	}

}

func TestCharCompare(t *testing.T) {
	inArray := []rune{'ँ', 'అ'}
	outArray := []rune{'ಁ', 'അ'}

	// multiByteIn := []string {`aː`, `ಆ` }
	// multiByteOut := []string {`ആ`, `ā`}

	for index, value := range inArray {
		if x := CharCompare(value, outArray[index]); !x {
			t.Errorf("CharCompare(%v, %v) = %v we need %v", value, outArray[index], x, true)
		}
	}

	// for index, value := range multiByteIn {
	// 	if x := CharCompare(value, multiByteOut[index]); !x {
	// 		t.Errorf("CharCompare(%v, %v) = %v we need %v", value, multiByteOut[index], x, true)
	// 	}
	// }
}
