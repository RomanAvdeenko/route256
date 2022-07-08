package matrix

import (
	"bufio"
	"fmt"
	"matrix/utils"
)

func Init(scanner *bufio.Scanner, writer *bufio.Writer) [][]int {
	var tmp int

	rowCount, _ := utils.ScanInt(scanner)
	colCount, _ := utils.ScanInt(scanner)

	//
	//fmt.Fprintln(writer, "init:", rowCount, "x", colCount)
	//
	//// Init slice
	sl := make([][]int, rowCount)
	for r := 0; r < rowCount; r++ {
		sl[r] = make([]int, colCount)
		for c := 0; c < colCount; c++ {
			tmp, _ = utils.ScanInt(scanner)
			sl[r][c] = tmp
		}
	}
	writer.Flush()
	return sl
}

func SwapRows(sl [][]int, r1, r2 int) {
	if r1 == r2 {
		return
	}

	for i := 0; i < len(sl[0]); i++ {
		sl[r1][i], sl[r2][i] = sl[r2][i], sl[r1][i]
	}
}

func Print(writer *bufio.Writer, sl [][]int) {
	defer writer.Flush()

	colLen := len(sl[0])
	//rowLen := len(sl)

	////fmt.Fprintln(writer, "print:", rowLen, "x", colLen)

	for _, r := range sl {
		for colNum, val := range r {
			fmt.Fprint(writer, val)
			if colNum != colLen-1 {
				fmt.Fprint(writer, " ")
			}
		}
		fmt.Fprintln(writer)
	}
	fmt.Fprintln(writer)
}
