package main

import (
	"bufio"
	"matrix"
	"matrix/utils"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(in)
	//Set the split function for the scanning words operation.
	scanner.Split(bufio.ScanWords)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	testLen, _ := utils.ScanInt(scanner)
	scanner.Text()

	for test := 0; test < testLen; test++ {
		//fmt.Fprintln(out, "#test:", test)
		sl := matrix.Init(scanner, out)
		doTest(scanner, out, sl)
		matrix.Print(out, sl)
	}
}

func doTest(scanner *bufio.Scanner, writer *bufio.Writer, sl [][]int) {
	var tmp int

	rowCount := len(sl)
	colCount := len(sl[0])
	jobCount, _ := utils.ScanInt(scanner)
	jobs := make([]int, jobCount)

	// Get jobs
	for i := 0; i < jobCount; i++ {
		tmp, _ = utils.ScanInt(scanner)
		jobs[i] = tmp - 1
	}
	scanner.Text()

	// Make jobs
	sortedSl := make([][]int, rowCount)
	for r := 0; r < rowCount; r++ {
		sortedSl[r] = make([]int, colCount)
	}
	for _, job := range jobs {
		col := make([]matrix.NV, rowCount)
		for r := 0; r < rowCount; r++ {
			col[r] = matrix.NV{N: r, V: sl[r][job]}
		}
		sort.Stable(matrix.ByValue(col))
		for r := 0; r < rowCount; r++ {
			copy(sortedSl[r], sl[col[r].N])
		}
		for r := 0; r < rowCount; r++ {
			copy(sl[r], sortedSl[r])
		}
	}
	writer.Flush()
}
