package main

import (
	"fmt"
	"strings"

	"github.com/kefniark/advent-of-code-2022/utils"
	"github.com/samber/lo"
)

var handPoints = map[string]int{"A": 1, "X": 1, "B": 2, "Y": 2, "C": 3, "Z": 3}

func main() {
	lines := utils.ReadLines("day02/input.txt")
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	points := lo.Map(lines, func(line string, index int) int {
		values := handValues(line)
		result := calculateGamePoints(values)
		return result + values[1]
	})

	fmt.Println("Part 1 =", lo.Sum(points))
}

func part2(lines []string) {
	points := lo.Map(lines, func(line string, index int) int {
		values := handValuesPart2(line)
		result := calculateGamePoints(values)
		return result + values[1]
	})

	fmt.Println("Part 2 =", lo.Sum(points))
}

func calculateGamePoints(res []int) int {
	if res[0] == res[1] {
		return 3
	}
	if (res[0] == 1 && res[1] == 2) || (res[0] == 2 && res[1] == 3) || (res[0] == 3 && res[1] == 1) {
		return 6
	}
	return 0
}

func handValues(s string) []int {
	return lo.Map(strings.Split(s, " "), func(s string, index int) int {
		return handPoints[s]
	})
}

func handValuesPart2(s string) []int {
	val := strings.Split(s, " ")
	opponent := handPoints[val[0]]
	me := 0
	if val[1] == "Y" {
		me = opponent
	} else if val[1] == "X" {
		me = (opponent+1)%3 + 1
	} else {
		me = opponent%3 + 1
	}

	return []int{opponent, me}
}
