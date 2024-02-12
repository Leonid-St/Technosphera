package main

import (
	"bufio"
	"fmt"
	"os"
)

func isProcessCorrectUniversal(process string) bool {
	stack := []rune{}

	for _, action := range process {
		switch action {
		case 'M':
			stack = append(stack, 'M')
		case 'R':
			if len(stack) == 0 || stack[len(stack)-1] != 'M' {
				return false
			}
		case 'C':
			if len(stack) == 0 || (stack[len(stack)-1] != 'M' && stack[len(stack)-1] != 'D') {
				return false
			}

			stack = stack[:len(stack)-1]
		case 'D':
			if len(stack) == 0 || stack[len(stack)-1] != 'D' {
				return false
			}
		}
	}

	return len(stack) == 0 || stack[len(stack)-1] == 'D'
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var a, i int32

	fmt.Fscan(in, &a)
	var answer []string = make([]string, a)
	for i = 0; i < a; i++ {
		var str string
		fmt.Fscan(in, &str)
		if isProcessCorrectUniversal(str) {
			answer[i] = "YES"
		} else {
			answer[i] = "NO"
		}
	}
	for i = 0; i < a; i++ {
		fmt.Fprint(out, answer[i]+"\n")
	}
}
