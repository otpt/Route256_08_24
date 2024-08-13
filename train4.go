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

        user := make([]int, n)

        for j := 0; j < n; j++ {
            fmt.Fscan(in, &user[j])
        }

        k := 0
        current := map[int]int{}
        ans := 0
        for j := 0; j < n; j++ {            
            for k < n && len(current) < 3 {
                if _, ok := current[user[k]]; ok {
                    current[user[k]]++
                } else {
                    current[user[k]] = 1
                }
                
                if (len(current) < 3) {
                    ans = max(ans, k - j + 1)
                    //fmt.Fprintf(out, "max: %d   %d\n", j, k)
                    //fmt.Fprintf(out, "max: %d\n", ans)
                    //fmt.Fprintln(out, current)
                    
                }
                
                k++
            }
            current[user[j]]--
            if (current[user[j]] == 0) {
                delete(current, user[j])
            }
        }
        
        fmt.Fprintf(out, "%d\n", ans)
    }
}

func max (a, b int) int {
    if (a > b) {
        return a
    } 
    return b
}