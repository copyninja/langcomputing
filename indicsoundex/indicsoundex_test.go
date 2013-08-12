package indicsoundex

import "testing"

func TestCalculate(t *testing.T) {
	// inArray := []string {"vasudeva", "kamath", "ವಾಸುದೇವ", "वासुदॆव"}
	inArray := []string{`ವಾಸುದೇವ`, `वासुदॆव`}
	// outArray := []string {"vA2C3D1A", "kA5A3000", "ವASCKDR0", "वASCKDR0" }
	outArray := []string{`ವASCKDR0`, `वASCKDR0`}

	for index, value := range inArray {
		if x, output := Calculate(value, 8), outArray[index]; x != output {
			t.Errorf("Calculate(%s) = %s was expecting %s", value, x, output)
		}
	}
}
