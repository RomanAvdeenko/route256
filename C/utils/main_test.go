package utils

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestScanInt(t *testing.T) {
	testCases := []struct {
		name  string
		have  string
		want  []int
		valid bool
	}{
		{
			"Strings of int",
			"1 34 57 3 -234 2\n-3 5 -33 23 5304\n",
			[]int{1, 34, 57, 3, -234, 2, -3, 5, -33, 23, 5304},
			true,
		},
		{
			"Invalid test",
			"1 34 4325",
			[]int{1, 34, 4325, 12},
			false,
		},
		{
			"Blank input test",
			"\n \n",
			[]int{},
			true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := make([]int, 0)

			scanner := bufio.NewScanner(strings.NewReader(test.have))
			//Set the split function for the scanning words operation.
			scanner.Split(bufio.ScanWords)
			for {
				res, err := ScanInt(scanner)
				if err != nil {
					break
				}
				result = append(result, res)
			}
			if test.valid {
				assert.NoError(t, scanner.Err())
				assert.Equal(t, test.want, result)
			} else {
				assert.NoError(t, scanner.Err())
				assert.NotEqual(t, test.want, result)
			}
		})
	}

}
