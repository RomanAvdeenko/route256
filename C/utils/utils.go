package utils

import (
	"bufio"
	"strconv"
)

func ScanInt(scanner *bufio.Scanner) (int, error) {
	scanner.Scan()
	res, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return 0, err
	}
	return res, nil
}
