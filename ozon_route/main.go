package main

import "fmt"

func main() {
	// for i := 1; i <= 15; i++ {
	// 	Reader("./exam/difference_in_number/main.go",
	// 		fmt.Sprintf("./1_SeaBattle/tests/%d", i),
	// 		fmt.Sprintf("./1_SeaBattle/tests/%d.a", i))
	// }
	// for i := 1; i <= 25; i++ {
	// 	Reader("./exam/difference_in_number/main.go",
	// 		fmt.Sprintf("./exam/1_difference_in_number/tests/%d", i),
	// 		fmt.Sprintf("./exam/1_difference_in_number/tests/%d.a", i))
	// }

	for i := 1; i <= 25; i++ {
		Reader("./exam/2_rounding_error/main.go",
			fmt.Sprintf("./exam/2_rounding_error/tests/%d", i),
			fmt.Sprintf("./exam/2_rounding_error/tests/%d.a", i))
	}
	// for i := 1; i <= 25; i++ {
	// 	Reader("./exam/3_hight_load_system/main",
	// 		fmt.Sprintf("./exam/3_hight_load_system/tests/%d", i),
	// 		fmt.Sprintf("./exam/3_hight_load_system/tests/%d.a", i))
	// }
}
