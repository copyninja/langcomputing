package approxsearch

import (
	"strings"
)

type Strings []string

type Index interface {
	index(char string) int
}

func (s Strings) index(char string) int {
	for index, value := range s {
		if value == char {
			return index
		}
	}

	return -1
}

func createBigram(stringSlice []string) (bigramSlice []string) {
	bigramSlice = make([]string, 1)
	for i := 1; i < len(stringSlice); i++ {
		bigramSlice = append(bigramSlice, strings.Join(stringSlice[i-1:i+1], ""))
	}

	return
}

func countCommon(shortBigram, longBigram []string, average float64) (common float64) {
	common = float64(0)
	for indexS, bigram := range shortBigram {
		if indexL := Strings(longBigram).index(bigram); indexL != -1 {
			if indexL == indexS {
				common += float64(1)
			} else {
				dislocation := float64(indexL-
					indexS) / average
				if dislocation < 0 {
					dislocation *= -1
				}

				common += 1.0 - dislocation
			}

			longBigram[indexL] = ""
		}
	}

	return
}

func BigramAverage(str1, str2 string) float64 {
	if str1 == str2 {
		return float64(1)
	}

	bigramOne := createBigram(strings.Split(str1, ""))
	bigramTwo := createBigram(strings.Split(str2, ""))

	average := float64(len(bigramOne)+len(bigramTwo)) / 2.0

	common := float64(0)

	if len(bigramOne) < len(bigramTwo) {
		common = countCommon(bigramOne, bigramTwo, average)
	} else if len(bigramOne) > len(bigramTwo) {
		common = countCommon(bigramTwo, bigramOne, average)
	}

	return common / average
}
