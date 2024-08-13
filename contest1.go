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
        var str string
        fmt.Fscan(in, &str)
        
        if len(str) == 1 {
            fmt.Fprintf(out, "0\n")
            continue
        }

        changed := false
        var index = -1
        for i := 1; i < len(str) ; i++ {
            if str[i] > str[i - 1] {
                if (i == len(str) - 1) {
                    fmt.Fprintf(out, "%s\n", str[0 : i - 1] + str[i : ])
                } else {
                    fmt.Fprintf(out, "%s\n", str[0 : i - 1 ] + str[i : len(str)])
                }
                changed = true
                break
            }
        }

        if changed {continue}

        if (index == -1) {
            index = len(str) - 1
        }

        if index < len(str) - 1 {
            fmt.Fprintf(out, "%s\n", str[0 : index] + str[index + 1 : len(str)]) 
        } else {
            fmt.Fprintf(out, "%s\n", str[0 : index]) 
        }
   }
}
