package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var a, b int64
	fmt.Fscan(in, &a, &b)
	fmt.Fprint(out, fmt.Sprintf("%d\n", a-b))
}
