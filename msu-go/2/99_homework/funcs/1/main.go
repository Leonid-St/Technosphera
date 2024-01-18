package main

import (
	"fmt"
)

//  Golang program for
//  Conversion from Decimal to roman number

// Display roman value of n
func result(n int) string {
	switch n {
	case 1:
		return "I"
	case 4:
		return "IV"
	case 5:
		return "V"
	case 9:
		return "IX"
	case 10:
		return "X"
	case 40:
		return "XL"
	case 50:
		return "L"
	case 90:
		return "XC"
	case 100:
		return "C"
	case 400:
		return "CD"
	case 500:
		return "D"
	case 900:
		return "DM"
	case 1000:
		return "M"
	}
	return ""
}
func selecting(number int, collection []int, size int) (int, string) {
	var n int = 1
	var i int = 0
	for i = 0; i < size; i++ {
		if number >= collection[i] {
			n = collection[i]
		} else {
			break
		}
	}
	res := result(n)
	return number - n, res
}
func romanNo(number int) (res string) {
	if number <= 0 {
		// When is not a natural number
		return
	}
	// Base case collection
	var collection = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	// Get the size of collection
	var size int = int(len(collection))
	fmt.Printf(" %d : ", number)
	result := ""
	for number > 0 {
		number, result = selecting(number, collection, size)
		res += result
	}
	return res
}

type memoizeFunction func(int, ...int) interface{}

// TODO реализовать
var fibonacci memoizeFunction
var romanForDecimal memoizeFunction

//TODO Write memoization function

func memoize(function memoizeFunction) memoizeFunction {
	m := make(map[any]any)
	return func(arg int, arg1 ...int) (res interface{}) {
		v, ok := m[arg]
		if ok {
			res = v
		} else {
			res = function(arg)
			m[arg] = res.(any)
		}
		return res
	}
}

// TODO обернуть функции fibonacci и roman в memoize
func init() {
	fibonacci = memoize(
		func(arg2 int, arg3 ...int) (res any) {
			first := 0
			second := 1
			i := 0
			next := 0
			for {
				if i > arg2-2 {
					res = second
					break
				}
				next = first + second
				first = second
				second = next
				i++
			}
			return res
		})
	romanForDecimal = memoize(
		func(arg2 int, arg3 ...int) (res any) {
			return romanNo(arg2)
		})

	// fibonacci = (func(arg int, arg1 ...int) interface{} {
	// 	m := make(map[int]int)
	// 	m[0] = 0
	// 	m[1] = 1
	// 	first := m[0]
	// 	second := m[1]
	// 	fib := func(arg2 int) (res int) {
	// 		v, ok := m[arg2]
	// 		if ok {
	// 			res = v
	// 		} else {
	// 			i := 0
	// 			next := 0
	// 			for {
	// 				if i > arg2-2 {
	// 					res = second
	// 					break
	// 				}
	// 				next = first + second
	// 				first = second
	// 				second = next
	// 				i++
	// 			}
	// 		}
	// 		return res
	// 	}
	// 	if arg != 0 {
	// 		m[arg] = fib(arg)
	// 	}
	// 	if arg1 != nil {
	// 		argI := 2
	// 		for {
	// 			m[argI] = fib(argI)
	// 		}
	// 	}
	// 	return fib
	// })(45).(memoizeFunction)

}

func main() {

	fmt.Println("Fibonacci(45) =", fibonacci(45).(int))
	for _, x := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
		14, 15, 16, 17, 18, 19, 20, 25, 30, 40, 50, 60, 69, 70, 80,
		90, 99, 100, 200, 300, 400, 500, 600, 666, 700, 800, 900,
		1000, 1009, 1444, 1666, 1945, 1997, 1999, 2000, 2008, 2010,
		2012, 2500, 3000, 3999} {
		fmt.Printf("%4d = %s\n", x, romanForDecimal(x))
	}
}
