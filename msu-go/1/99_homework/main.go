package main

import (
	"reflect"
	"sort"
	"strconv"
)

func main() {

}

func ReturnInt() int {
	return 1
}

func ReturnFloat() float32 {
	return 1.1
}

func ReturnIntArray() [3]int {
	return [3]int{1, 3, 4}
}

func ReturnIntSlice() []int {
	return []int{1, 2, 3}
}

func IntSliceToString(slInt []int) string {
	var b = ""
	for _, a := range slInt {
		b = b + strconv.Itoa(a)
	}
	return b
}

func MergeSlices(arg ...any) any {
	var c []int
	for _, v := range arg {
		for _, q := range v.([]any) {
			switch v := reflect.ValueOf(q); v.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				c = append(c, q.(int))
			case reflect.Float32:
				c = append(c, int(q.(float32)))
			case reflect.Float64:
				c = append(c, int(q.(float64)))
			}
		}
	}
	return c
}

func GetMapValuesSortedByKey(m map[int]string) []string {
	res := make([]string, len(m))
	a := make([]int, len(m))
	i := 0
	for k := range m {
		a[i] = k
		i++
	}
	i = 0
	sort.Ints(a)
	for _, v := range a {
		
		res[i] = m[v]
		i++
	}
	return res
}
