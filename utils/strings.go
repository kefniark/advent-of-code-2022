package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadLines(filename string) []string {
	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}

	content := []string{}

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		content = append(content, strings.TrimSpace(sc.Text()))
	}

	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
	}

	return content
}

func SplitByLine(content []string, splitter string) [][]string {
	res := [][]string{}

	current := []string{}
	for _, line := range content {
		if line == splitter {
			res = append(res, current)
			current = []string{}
		} else {
			current = append(current, line)
		}
	}

	if len(current) > 0 {
		res = append(res, current)
	}

	return res
}
