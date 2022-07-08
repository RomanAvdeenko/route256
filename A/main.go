package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func getLinesNumber(scanner *bufio.Scanner) int {
	scanner.Scan()
	num := scanner.Text()
	res, _ := strconv.Atoi(num)
	return res
}

func getSumOfLine(str string) int {
	words := strings.Split(str, " ")

	a, _ := strconv.Atoi(words[0])
	b, _ := strconv.Atoi(words[1])

	return a + b
}

func main() {
	var (
		scanner = bufio.NewScanner(os.Stdin)

		num, sum int
	)

	num = getLinesNumber(scanner)

	for i := 0; i < num; i++ {
		scanner.Scan()
		sum = getSumOfLine(scanner.Text())

		io.WriteString(os.Stdout, strconv.Itoa(sum)+"\n")
	}
}
