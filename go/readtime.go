package readtime

import (
	"errors"
	"math"
	"regexp"
)

var wordRE = regexp.MustCompile(`[\p{L}][\p{L}\p{N}'’-]*`)

func CountWords(text string) int {
	return len(wordRE.FindAllString(text, -1))
}

func ReadingTime(text string, PPM int) (float64, error) {
	if PPM < 80 {
		return 0, errors.New("PPM must be >= 80")
	}
	min := float64(CountWords(text)) / float64(PPM)
	return math.Round(min*100) / 100, nil
}