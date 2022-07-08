package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type result bool

var loginRE *regexp.Regexp

func (r result) String() string {
	if r == true {
		return "YES"
	}
	return "NO"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	//loginRE, _ = regexp.Compile(`^[^-][a-zA-Z0-9_-]{1,23}$`)
	loginRE, _ = regexp.Compile(`^[a-zA-Z0-9_][a-zA-Z0-9_-]{1,23}$`)

	testNum, _ := ScanInt(scanner)

	for test := 0; test < testNum; test++ {
		doTest(scanner, out)
		fmt.Fprintln(out)
	}
}

func doTest(scanner *bufio.Scanner, out *bufio.Writer) {
	goodLogins := make(map[string]struct{})
	loginNum, _ := ScanInt(scanner)

	for c := 0; c < loginNum; c++ {
		scanner.Scan()
		text := scanner.Text()
		// Validate login
		validated := validate(text)
		if validated {
			// Check if exist
			// Store logins in LowerCase
			_, ok := goodLogins[strings.ToLower(text)]
			if ok {
				validated = false
			} else {
				goodLogins[strings.ToLower(text)] = struct{}{}
			}
		}
		fmt.Fprintln(out, validated)
		//fmt.Fprintf(out, "%v %v(%v)\n", validated, text, len(text))
	}
}

func validate(val string) result {
	return result(loginRE.MatchString(val))
}

func ScanInt(scanner *bufio.Scanner) (int, error) {
	scanner.Scan()
	res, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return 0, err
	}
	return res, nil
}
