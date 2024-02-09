package main

import "fmt"

func main() {
	for i := 1; i <= 15; i++ {
		Reader("./1_SeaBattle/1_sea_battle.go",
			fmt.Sprintf("./1_SeaBattle/tests/%d", i),
			fmt.Sprintf("./1_SeaBattle/tests/%d.a", i))
	}
}
