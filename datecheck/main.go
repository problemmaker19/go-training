package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isLeapYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return true
	}
	return false
}

func dateCheck(date []string) string {
	day, _ := strconv.Atoi(date[0])
	month, _ := strconv.Atoi(date[1])
	year, _ := strconv.Atoi(date[2])

	shouldBe := 30 + (month+month/8)%2

	if month == 2 {
		if isLeapYear(year) {
			shouldBe = 29
		} else {
			shouldBe = 28
		}
	}

	if shouldBe < day {
		return "NO"
	}
	return "YES"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		lineSlice := strings.Fields(line)
		result := dateCheck(lineSlice)
		fmt.Println(result)
	}
}
