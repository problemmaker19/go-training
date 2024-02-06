package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func gameCheck(line []string) string {
	var sum, count int
	for _, v := range line {
		num, _ := strconv.Atoi(v)
		sum += num
		if num == 2 {
			count++
		}
	}

	if sum != 20 || count != 3 {
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
		result := gameCheck(lineSlice)
		fmt.Println(result)
	}
}
