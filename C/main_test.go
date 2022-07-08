package matrix

import (
	"bufio"
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwapRows(t *testing.T) {
	testCases := []struct {
		name  string
		have  [][]int
		r1    int
		r2    int
		want  [][]int
		valid bool
	}{
		{
			"4x3 valid",
			[][]int{
				[]int{3, 4, 1},
				[]int{2, 2, 5},
				[]int{2, 4, 2},
				[]int{2, 2, 1},
			},
			0,
			3,
			[][]int{
				[]int{2, 2, 1},
				[]int{2, 2, 5},
				[]int{2, 4, 2},
				[]int{3, 4, 1},
			},
			true,
		},
		{
			"5x1 valid",
			[][]int{
				[]int{3},
				[]int{2},
				[]int{2},
				[]int{9},
			},
			1,
			2,
			[][]int{
				[]int{3},
				[]int{2},
				[]int{2},
				[]int{9},
			},
			true,
		},
		{
			"4x3 invalid",
			[][]int{
				[]int{3, 4, 1},
				[]int{2, 2, 5},
				[]int{2, 4, 2},
				[]int{2, 2, 1},
			},
			0,
			3,
			[][]int{
				[]int{3, 4, 1},
				[]int{2, 2, 5},
				[]int{2, 4, 2},
				[]int{2, 2, 1},
			},
			false,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			SwapRows(test.have, test.r1, test.r2)
			if test.valid {
				assert.Equal(t, test.want, test.have)
			} else {
				assert.NotEqual(t, test.want, test.have)
			}
		})
	}
}

func TestPrint(t *testing.T) {
	buffer := new(bytes.Buffer)
	writer := bufio.NewWriter(buffer)

	testCases := []struct {
		name string
		have [][]int
		want string
	}{
		{
			"Print 4x3",
			[][]int{
				[]int{3, 4, 1},
				[]int{2, 2, 5},
				[]int{2, 4, 2},
				[]int{2, 2, 1},
			},
			"3 4 1\n2 2 5\n2 4 2\n2 2 1\n\n",
		},
		{
			"Print 3x1",
			[][]int{
				[]int{100},
				[]int{9},
				[]int{10},
			},
			"100\n9\n10\n\n",
		},
		{
			"Print 7x1",
			[][]int{
				[]int{13, 48, 1, 54, 2, -23, -3},
			},
			"13 48 1 54 2 -23 -3\n\n",
		},
		{
			"Print blank",
			[][]int{[]int{}},
			"\n\n",
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			Print(writer, test.have)
			assert.Equal(t, test.want, buffer.String())
			buffer.Reset()
		})
	}
}
