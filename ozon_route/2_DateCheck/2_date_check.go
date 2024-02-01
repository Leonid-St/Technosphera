package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func DateCheck() {
	var numOfLine int
	scanner := bufio.NewScanner(os.Stdin)
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
		dateSlice := strings.Fields(date)
		fmt.Println(dateSlice)
		d := dateSlice[0]
		m := dateSlice[1]
		y := dateSlice[2]
		stringDate := ""
		// parse string date to golang time
		_, err := time.Parse(d+"/"+m+"/"+y, stringDate)
		if err != nil {
			answers[i] = "NO"
			fmt.Println(err.Error())

		} else {
			answers[i] = "YES"
		}

	}
	for _, j := range answers {
		fmt.Println(j)
	}
}
