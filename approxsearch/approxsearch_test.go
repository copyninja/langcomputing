package approxsearch

import "testing"

type bigramAverageResult struct {
	input1, input2 string
	bigramAverage  float64
}

func runBigramAverageTests(in1, in2 []string, c chan bigramAverageResult) {
	for index, value := range in1 {
		bg := BigramAverage(value, in2[index])
		c <- bigramAverageResult{value, in2[index], bg}
	}

	close(c)
}

func TestBigramAverage(t *testing.T) {
	input1 := []string{"Mangalore", "Cauliflower"}
	input2 := []string{"Bangalore", "Sunflower"}
	output := []float64{0.875, 0.0}

	ch := make(chan bigramAverageResult)

	go runBigramAverageTests(input1, input2, ch)

	for result := range ch {
		if index := Strings(input1).index(result.input1); index != -1 {
			if result.bigramAverage != output[index] {
				t.Errorf("BigramAverage(%s, %s) = %f was expecting %f",
					result.input1,
					result.input2, result.bigramAverage,
					output[index])
			}
		} else {
			t.Fatalf("Unexpected input? %s, %s failing now!",
				result.input1, result.input2)
		}
	}

}
