package main

import (
	"fmt"
	"slices"

	"github.com/kefniark/advent-of-code-2022/utils"
	"github.com/samber/lo"
)

func main() {
	lines := utils.ReadLines("day01/input.txt")
	elves := utils.SplitByLine(lines, "")

	calories := lo.Map(elves, func(e []string, index int) int {
		return lo.Sum(utils.ParseListInt(e))
	})

	part1(calories)
	part2(calories)
}

func part1(calories []int) {
	fmt.Println("Part 1 =", lo.Max(calories))
}

func part2(calories []int) {
	slices.SortFunc(calories, func(a, b int) int {
		return a - b
	})
	top3 := calories[len(calories)-3:]
	fmt.Println("Part 2 =", lo.Sum(top3))
}
