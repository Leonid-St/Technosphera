package main

import "fmt"

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
			stack[len(stack)-1] = 'M'
		case 'D':
			if len(stack) == 0 || stack[len(stack)-1] != 'M' {
				return false
			}
			stack[len(stack)-1] = 'D'
		}
	}

	return len(stack) == 1 && stack[0] == 'D'
}

func main() {
	processes := []string{"MRCMD", "MDD", "M", "MDMRCMD", "MMDD"}

	for _, process := range processes {
		if isProcessCorrectUniversal(process) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
