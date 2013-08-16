package indicsoundex

import "testing"

func TestCalculate(t *testing.T) {
	inArray := []string{"vasudeva", "kamath", "ವಾಸುದೇವ", "वासुदॆव"}
	outArray := []string{"v231", "k53", "ವASCKDR0", "वASCKDR0"}

	for index, value := range inArray {
		if x, output := Calculate(value, 8), outArray[index]; x != output {
			t.Errorf("Calculate(%s) = %s was expecting %s", value,
				x, output)
		}
	}
}

func TestCompare(t *testing.T) {
	inArray := []string{"vasudev", "ವಾಸುದೇವ", "वासुदॆव", "vasudev",
		"vasudev"}
	outArray := []string{"kamath", "വാസുദേവ", "వాసుదేవ", "vaasudev",
		"vasudev"}
	resultArray := []int{SOUNDEX_STRINGS_SOUNDS_DIFFERENT,
		SOUNDEX_STRINGS_SOUNDS_ALIKE,
		SOUNDEX_STRINGS_SOUNDS_ALIKE,
		SOUNDEX_STRINGS_SOUNDSS_ALIKE,
		SOUNDEX_SAME_STRING}

	for index, value := range inArray {
		if x := Compare(value, outArray[index]); x != resultArray[index] {
			t.Errorf("Compare(%s, %s) = %d was expecting %d", value,
				outArray[index], x, resultArray[index])
		}
	}
}
