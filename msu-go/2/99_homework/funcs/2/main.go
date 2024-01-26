package main

import (
	"fmt"
)

func showMeTheType(i interface{}) string {
	return fmt.Sprintf("%T", i)
}

func main() {

}
