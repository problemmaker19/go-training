package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func compress(s *bufio.Scanner) {
	s.Scan()
	text := s.Text()

	arr := strings.Split(text, " ")
	prev, _ := strconv.Atoi(arr[0])
	counter := 0

	for i, v := range arr {
		if i == 0 {
			continue
		}

		cur, _ := strconv.Atoi(v)

		if prev-cur == 1 {
			prev = cur
			counter--
			continue
		}
		// if cur-prev == 1 {
		// 	prev = cur
		// 	continue
		// }

		counter = 0
		fmt.Print(-counter-cur, counter, "; ")
		prev = cur
	}
	fmt.Println()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	tasks, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < tasks; i++ {
		scanner.Scan()
		compress(scanner)
	}
}
