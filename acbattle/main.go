package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func lineReadInt(s *bufio.Scanner) (int, error) {
	s.Scan()
	v, err := strconv.Atoi(s.Text())

	if err != nil {
		return 0, err
	}

	return v, nil
}

func fight(s *bufio.Scanner, n int) {

	min, max := 15, 30
	limiter := make(map[byte]int)
	limiter[62] = min
	limiter[60] = max

	for i := 0; i < n; i++ {

		s.Scan()
		adj := s.Text()

		num, _ := strconv.Atoi(adj[3:])
		operator := adj[0]

		if operator == '>' && num > limiter[operator] {
			limiter[operator] = num
		}

		if operator == '<' && num < limiter[operator] {
			limiter[operator] = num
		}

		if limiter[62] > limiter[60] {
			fmt.Println(-1, limiter)
		} else {
			fmt.Println(limiter[62], limiter)
		}
	}

	fmt.Println("")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tasks, _ := lineReadInt(scanner)

	for i := 0; i < tasks; i++ {
		tasknum, _ := lineReadInt(scanner)
		fight(scanner, tasknum)
	}
	fmt.Println("")
}
