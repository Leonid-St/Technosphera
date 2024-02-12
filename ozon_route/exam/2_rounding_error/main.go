package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func gerRightPercentage(price int64, percentage int) float64 {
	return float64(price) * (float64(percentage) / float64(100))
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var numberOfDate int
	fmt.Fscan(in, &numberOfDate)
	var n, p int
	var answer []float64 = make([]float64, numberOfDate)
	for i := 0; i < numberOfDate; i++ {
		fmt.Fscan(in, &n, &p)
		for k := 0; k < n; k++ {
			var ai int64
			fmt.Fscan(in, &ai)
			rp := gerRightPercentage(ai, p)
			_, fraction := math.Modf(rp)
			answer[i] = answer[i] + fraction
		}
	}
	for i := 0; i < numberOfDate; i++ {
		if answer[i] == 0 {
			fmt.Fprintln(out, "0.00")
		} else {
			fmt.Fprint(out, fmt.Sprintf("%.2f\n", answer[i]))
		}
	}
}
