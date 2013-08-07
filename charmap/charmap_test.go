package charmap

import "testing"

func TestLanguageOf(t *testing.T) {
	const in, out = 'ಅ', "kn_IN"
	if x := LanguageOf(in); x != out {
		t.Errorf("LanguageOf(%v) = %s we need %s", in, x, out)
	}
}

func TestCharCompare(t *testing.T) {
	const in1, in2 = 'ಅ', 'अ'
	if x := CharCompare(in1, in2); !x {
		t.Errorf("CharCompare(%v, %v) = %v we need %v", in1, in2, x, true)
	}
}
