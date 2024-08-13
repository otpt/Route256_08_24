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
        var m int

        fmt.Fscan(in, &m)
        var maxNumber int = -1

        var parent [101]int = [101]int{}
        
        for m > 0 {
            var number, subsNumber, current int
            fmt.Fscan(in, &number)

            maxNumber = max(number, maxNumber)
            m--

            fmt.Fscan(in, &subsNumber)
            m--

            for j := 0; j < subsNumber; j++ {
                fmt.Fscan(in, &current)
                parent[current] = number
                m--
            }
        }
        for j := 1; j <= maxNumber; j++ {
            if parent[j] == 0 {
                fmt.Fprintf(out, "%d\n", j)
            }
        }
    }
}

func max (a, b int) int {
    if (a > b) {
        return a
    } 
    return b
}