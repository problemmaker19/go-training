package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func lineReadInt(s *bufio.Scanner) int {
	s.Scan()
	v, _ := strconv.Atoi(s.Text())

	return v
}

func compress(s *bufio.Scanner) {
	s.Scan()
	s.Scan()
	data := s.Text()
	fmt.Println(data)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tasks := lineReadInt(scanner)

	for i := 0; i < tasks; i++ {
		compress(scanner)
	}
}
