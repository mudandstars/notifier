package utils

import (
	"fmt"
	"regexp"
	"strconv"
)

func Path(url string, pattern string) (uint, error) {
	fullPattern := pattern + `/(\d+)`
	re := regexp.MustCompile(fullPattern)

	submatches := re.FindStringSubmatch(url)

	if len(submatches) < 2 {
		return 0, fmt.Errorf("No match found in URL")
	}

	number := submatches[1]

	result, err := strconv.Atoi(number)
	if err != nil {
		return 0, err
	}

	return uint(result), nil
}
