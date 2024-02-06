package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	value1 := scanner.Text()
	scanner.Scan()
	value2 := scanner.Text()
	fmt.Println("Hello from test program1:-", value1)
	fmt.Println("Hello from test program2:-", value2)
}
