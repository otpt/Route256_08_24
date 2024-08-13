package main

import (
  "fmt"
  "bufio"
  "os"
  "math"
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
        fmt.Fscan(in, &n)
        min := make([]uint64, n)
        max := make([]uint64, n)

        for j := 0; j < n; j++ {
            fmt.Fscan(in, &min[j])
        }

        for j := 0; j < n; j++ {
            fmt.Fscan(in, &max[j])
        }


        var result uint64 = 1
        for j := 0; j < n; j++ {
            var low, top uint64
            if min[j] < uint64(j + 1) {
                low = 1
            } else {
                if min[j] % uint64(j + 1) == 0 {
                    low = min[j] / uint64(j + 1)
                } else {
                    low = min[j] / uint64(j + 1) + 1
                }
            }

            top = max[j] / uint64(j + 1)

            if (top < low) {
                // fmt.Fprintf(out, "0\n") 
                result = 0
                break
            }

            result = (result * (top - low + 1)) % (1000000007)
        }

        fmt.Fprintf(out, "%d\n", result) 

    }
}

func round(x float64) int{
    return int(math.Floor(x + 0.5))
}
