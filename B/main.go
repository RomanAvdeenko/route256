package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	text := scanner.Text()
	testCount, _ := strconv.Atoi(text)

	for i := 0; i < testCount; i++ {
		sum := calcSum(scanner)
		fmt.Fprintln(out, sum)
	}
}

func calcSum(scanner *bufio.Scanner) int {
	scanner.Scan()
	text := scanner.Text()
	buyCount, _ := strconv.Atoi(text)

	buys := make([]int, buyCount)
	for i := 0; i < buyCount; i++ {
		scanner.Scan()
		text = scanner.Text()
		buy, _ := strconv.Atoi(text)
		buys[i] = buy
	}

	order := make(map[int]int)

	for _, buy := range buys {
		order[buy] += 1
	}

	var sum int
	for key, val := range order {
		discount := val / 3

		sum += (val - discount) * key
	}
	return sum
}
