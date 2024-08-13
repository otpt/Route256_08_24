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
        fmt.Fscan(in, &n)
        values := make([]uint64, n)

        for j := 0; j < n; j++ {
            fmt.Fscan(in, &values[j])
        }

        pairCounter := map[int][]int{}
        
        for j := 1; j < n; j++ {
            diff := values[j] - values[j - 1]
            if _, ok := pairCounter[int(diff)]; ok {
                pairCounter[int(diff)] = append(pairCounter[int(diff)], j)
            } else {
                pairCounter[int(diff)] = []int{j}
            }
        }

        var result uint64 = 0
        for _, v := range pairCounter { 
            near := 0
            for k := 0; k < len(v) - 1; k++ {
                if v[k + 1] - v[k] == 1 {
                    near++
                }
            }
            number := uint64(len(v))
            result += number * (number - 1) / 2 - uint64(near)
            
            // fmt.Printf("key[%s] value[%s]\n", k, v)
        }
        fmt.Fprintln(out, result) 
    }
}
