package utils

import (
	"regexp"
	"strconv"

	"github.com/samber/lo"
)

func ParseInt(content string) []int {
	re := regexp.MustCompile(`\d+`)
	numbers := []int{}
	for _, val := range re.FindAllStringSubmatch(content, -1) {
		num, _ := strconv.Atoi(val[0])
		numbers = append(numbers, num)
	}

	return numbers
}

func ParseListInts(list []string) [][]int {
	return lo.Map(list, func(x string, index int) []int {
		return ParseInt(x)
	})
}

func ParseListInt(list []string) []int {
	return lo.Map(list, func(x string, index int) int {
		return ParseInt(x)[0]
	})
}
