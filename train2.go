package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
		var in *bufio.Reader
		var out *bufio.Writer
		in = bufio.NewReader(os.Stdin)
		out = bufio.NewWriter(os.Stdout)
		defer out.Flush()

		var t int
		fmt.Fscan(in, &t)

		for i := 0; i < t; i++ {
				var n int 
				var p, a, ans int64
				ans = 0

				fmt.Fscan(in, &n, &p)

				for j := 0; j < n; j++ {
						fmt.Fscan(in, &a)

						ans += (p * a) % 100
				}

				fmt.Fprintf(out, "%.2f\n", float32(ans) / 100)
		}

}