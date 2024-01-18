package main

import (
	"fmt"
	"strconv"
	"time"
)

type Calendar struct {
	parsedTime time.Time
}

func (c Calendar) CurrentQuarter() (res int) {
	i, _ := strconv.Atoi(c.parsedTime.String()[5:7])

	if i%3 != 0 {
		res = (i / 3) + 1
	} else {
		res = (i / 3)
	}
	return res
}

func NewCalendar(parsedTime time.Time) (cal Calendar) {
	cal.parsedTime = parsedTime
	return cal
}

func main() {
	cases := []struct {
		month   string
		quarter int
	}{
		{month: "01", quarter: 1},
		{month: "02", quarter: 1},
		{month: "03", quarter: 1},
		{month: "04", quarter: 2},
		{month: "05", quarter: 2},
		{month: "06", quarter: 2},
		{month: "07", quarter: 3},
		{month: "08", quarter: 3},
		{month: "09", quarter: 3},
		{month: "10", quarter: 4},
		{month: "11", quarter: 4},
		{month: "12", quarter: 4},
	}
	for _, test := range cases {
		parsed, _ := time.Parse("2006-01-02", fmt.Sprintf("2015-%s-15", test.month))
		calendar := NewCalendar(parsed)
		actual := calendar.CurrentQuarter()
		fmt.Println(actual)
	}
}
