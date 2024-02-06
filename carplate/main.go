package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func parsePlate(line string) string {

	p1 := `\b(([A-Z][0-9][0-9][A-Z][A-Z])|([A-Z][0-9][A-Z][A-Z]))+\b`
	p2 := `([A-Z][0-9][0-9][A-Z][A-Z])|([A-Z][0-9][A-Z][A-Z])`

	r1 := regexp.MustCompile(p1)

	if !r1.MatchString(line) {
		return "-"
	}

	r2 := regexp.MustCompile(p2)

	matches := r2.FindAllString(line, -1)
	result := strings.Join(matches, " ")

	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	for scanner.Scan() {
		lineStr := scanner.Text()
		result := parsePlate(lineStr)
		fmt.Println(result)
	}
}
