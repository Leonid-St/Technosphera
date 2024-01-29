package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var numOfLine int
	scanner := bufio.NewScanner(os.Stdin)

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
		scanner.Scan()
		date := scanner.Text()
		d := strings.Fields(date)[0]
		m := strings.Fields(date)[0]
		y := strings.Fields(date)[0]
		stringDate := ""
		// 10 9 2022
		// 21 9 2022
		// 29 2 2022
		// 31 2 2022
		// 29 2 2000
		// 29 2 2100
		// 31 11 1999
		// 31 12 1999
		// 29 2 2024
		// 29 2 2023

		// parse string date to golang time
		_, err := time.Parse(d+"/"+m+"/"+y, stringDate)
		if err != nil {
			answers[i] = "NO"

		} else {
			answers[i] = "YES"
		}

	}
	for _, j := range answers {
		fmt.Println(j)
	}
}
