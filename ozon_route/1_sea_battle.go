package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type NumOfEachShips struct {
	numOf4 int
	numOf3 int
	numOf2 int
	numOf1 int
}

func main() {
	var numOfLine int
	scanner := bufio.NewScanner(os.Stdin)
	max4 := 1
	max3 := 2
	max2 := 3
	max1 := 4
	// mm := map[int]string{}

	// mm[1] = "2 1 3 1 2 3 1 1 4 2"
	// mm[2] = "1 1 1 2 2 2 3 3 3 4"
	// mm[3] = "1 1 1 1 2 2 2 3 3 4"
	// mm[4] = "4 3 3 2 2 2 1 1 1 1"
	// mm[5] = "4 4 4 4 4 4 4 4 4 4"

	//fmt.Scanf("%d\n", &numOfLine)
	scanner.Scan()
	numOfLineStr := scanner.Text()
	numOfLine, err := strconv.Atoi(numOfLineStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	i := 0
	answers := make([]string, numOfLine)
	for ; i < numOfLine; i++ {
		var lineShips string //= mm[i+1]
		scanner.Scan()
		lineShips = scanner.Text()
		var numOfEachShips NumOfEachShips
		for _, j := range strings.Fields(lineShips) {
			switch j {
			case "1":
				numOfEachShips.numOf1++
				if numOfEachShips.numOf1 > max1 {
					answers[i] = "NO"
					break
				}
			case "2":
				numOfEachShips.numOf2++
				if numOfEachShips.numOf2 > max2 {
					answers[i] = "NO"
					break
				}
			case "3":
				numOfEachShips.numOf3++
				if numOfEachShips.numOf3 > max3 {
					answers[i] = "NO"
					break
				}
			case "4":
				numOfEachShips.numOf4++
				if numOfEachShips.numOf4 > max4 {
					answers[i] = "NO"
					break
				}
			}
		}
		if answers[i] != "NO" {
			answers[i] = "YES"
		}
	}
	for _, j := range answers {
		fmt.Println(j)
	}
}
